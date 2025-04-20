package handlers

import (
	"math"
	"sort"
	"xbooklab/models"

	"github.com/gin-gonic/gin"
)

// Recommend 基于用户协同过滤的推荐接口，输入用户ID，返回推荐内容
// 新用户采用热门内容+随机话题的混合推荐策略
// 优化方向：可缓存用户相似度矩阵，更新周期24小时
func Recommend(c *gin.Context, repo models.UserBehaviorRepository) {
	currentUserID := c.GetUint("user_id")
	if currentUserID == 0 {
		c.JSON(400, gin.H{"error": "无效用户ID"})
		return
	}

	// 获取混合推荐结果
	rec := models.GenerateRecommendations(currentUserID)

	// 获取完整的推荐内容详情
	var books []models.Book
	models.GetDB().Where("id IN ?", rec.BookIDs).Find(&books)

	var topics []models.Topic
	models.GetDB().Where("id IN ?", rec.TopicIDs).Find(&topics)

	// 获取用户行为数据
	behaviors, _ := repo.GetAll()

	// 检查是否是冷启动用户(没有行为记录)
	isNewUser := true
	for _, b := range behaviors {
		if b.UserID == currentUserID {
			isNewUser = false
			break
		}
	}

	if isNewUser {
		// 冷启动推荐: 热门内容+随机话题
		popularBooks := models.GetPopularBooks(5)
		randomTopics := models.GetRandomTopics(3)

		c.JSON(200, gin.H{
			"recommendations": gin.H{
				"books":  popularBooks,
				"topics": randomTopics,
			},
		})
		return
	}

	// 2. 构建用户-物品评分矩阵
	userItemMatrix := make(map[uint]map[uint]float64)
	for _, b := range behaviors {
		if _, ok := userItemMatrix[b.UserID]; !ok {
			userItemMatrix[b.UserID] = make(map[uint]float64)
		}
		userItemMatrix[b.UserID][b.ItemID] = b.Rating
	}

	// 3. 计算当前用户与其他用户的相似度
	similarities := make(map[uint]float64)
	for userID, ratings := range userItemMatrix {
		if userID == currentUserID {
			continue
		}
		similarities[userID] = cosineSimilarity(userItemMatrix[currentUserID], ratings)
	}

	// 4. 获取最相似的K个用户
	k := 5
	topSimilarUsers := getTopKSimilarUsers(similarities, k)

	// 5. 基于相似用户的喜好推荐物品
	recommendations := getRecommendations(userItemMatrix, currentUserID, topSimilarUsers)

	// 6. 返回推荐结果
	c.JSON(200, gin.H{
		"recommendations": recommendations,
	})
}

// cosineSimilarity 计算余弦相似度
// 余弦相似度公式：cosθ = (A·B)/(||A||*||B||)
// 优化方向：对于稀疏矩阵可采用向量压缩存储优化
func cosineSimilarity(a, b map[uint]float64) float64 {
	var dotProduct, normA, normB float64

	// 计算共同物品
	commonItems := make(map[uint]bool)
	for itemID := range a {
		commonItems[itemID] = true
	}
	for itemID := range b {
		commonItems[itemID] = true
	}

	for itemID := range commonItems {
		dotProduct += a[itemID] * b[itemID]
		normA += a[itemID] * a[itemID]
		normB += b[itemID] * b[itemID]
	}

	if normA == 0 || normB == 0 {
		return 0
	}

	return dotProduct / (math.Sqrt(normA) * math.Sqrt(normB))
}

// getTopKSimilarUsers 获取最相似的K个用户
// 基于堆排序选择TopK相似用户，时间复杂度O(n logk)
// 优化方向：可引入相似度阈值过滤低相关性用户
func getTopKSimilarUsers(similarities map[uint]float64, k int) []uint {
	type userSimilarity struct {
		userID uint
		score  float64
	}

	var users []userSimilarity
	for userID, score := range similarities {
		users = append(users, userSimilarity{userID, score})
	}

	sort.Slice(users, func(i, j int) bool {
		return users[i].score > users[j].score
	})

	var result []uint
	for i := 0; i < k && i < len(users); i++ {
		result = append(result, users[i].userID)
	}

	return result
}

// getRecommendations 获取推荐物品
// 排除用户已交互物品，按加权评分降序排列
// 优化方向：建议采用批量数据库查询优化item_type判断
func getRecommendations(userItemMatrix map[uint]map[uint]float64, currentUserID uint, similarUsers []uint) map[string][]uint {
	recommendations := make(map[string][]uint)

	// 统计相似用户喜欢的物品
	itemScores := make(map[uint]float64)
	for _, userID := range similarUsers {
		for itemID, rating := range userItemMatrix[userID] {
			// 排除当前用户已经交互过的物品
			if _, ok := userItemMatrix[currentUserID][itemID]; !ok {
				itemScores[itemID] += rating
			}
		}
	}

	// 按评分排序
	type itemScore struct {
		itemID uint
		score  float64
	}

	var items []itemScore
	for itemID, score := range itemScores {
		items = append(items, itemScore{itemID, score})
	}

	sort.Slice(items, func(i, j int) bool {
		return items[i].score > items[j].score
	})

	// 获取前N个推荐
	n := 10
	var bookIDs, topicIDs []uint
	for i := 0; i < n && i < len(items); i++ {
		// 根据ItemType字段区分书籍和话题
		var behavior models.UserInteraction
		if err := models.GetDB().Where("item_id = ?", items[i].itemID).First(&behavior).Error; err == nil {
			if behavior.ItemType == "book" {
				bookIDs = append(bookIDs, items[i].itemID)
			} else if behavior.ItemType == "topic" {
				topicIDs = append(topicIDs, items[i].itemID)
			}
		}
	}

	recommendations["books"] = bookIDs
	recommendations["topics"] = topicIDs

	return recommendations
}
