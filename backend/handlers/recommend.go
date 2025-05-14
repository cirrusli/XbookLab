package handlers

import (
	"log"
	"math"
	"net/http"
	"sort"
	"xbooklab/models"

	"github.com/gin-gonic/gin"
)

type GetRecommendedBooksRequest struct {
	UserID    uint `json:"user_id"`
	Limit     uint `json:"limit"`
	Offset    uint `json:"offset"`
	TagFilter uint `json:"tag_filter"`
}

type GetRecommendedBooksResponse struct {
	Books                []models.Book `json:"books"`
	Algorithm            string        `json:"algorithm"`
	RecommendationReason string        `json:"recommendation_reason"`
}

// GetRecommendedBooks 实现书籍推荐逻辑
func GetRecommendedBooks(c *gin.Context) {
	// 解析URL查询参数
	var req GetRecommendedBooksRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	log.Printf("Received query params: %+v", req)
	var recommendedBooks []models.Book

	var err error
	userID := c.GetUint("userID")
	log.Println(userID)
	if userID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "登录已过期，请重新登录"})
		return
	}
	// 定义推荐策略的类型
	type recommendStrategy func(userID uint, limit int, tagFilter int) []models.Book

	// 推荐策略实现
	var strategies = []recommendStrategy{
		func(userID uint, limit int, tagFilter int) []models.Book {
			return recommendByCF(userID, tagFilter)
		},
		func(userID uint, limit int, tagFilter int) []models.Book {
			return recommendByEvent()
		},
		func(userID uint, limit int, tagFilter int) []models.Book {
			return models.GetPopularBooks(limit, tagFilter)
		},
	}

	realLimit := int(req.Limit) - 5
	log.Println("userID:", userID, "limit:", realLimit, "offset:", req.Offset, "tagFilter:", req.TagFilter)

	// 获取协同过滤推荐
	recommendedBooks, err = models.GetRecommendedBooks(userID, realLimit, int(req.Offset), int(req.TagFilter))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	log.Printf("推荐表中的数据量: %d recommendedBooks: %v", len(recommendedBooks), recommendedBooks)

	// 如果推荐数量不足，依次使用策略填充
	remaining := realLimit - len(recommendedBooks)

	for _, strategy := range strategies {
		if remaining <= 0 {
			break
		}
		strategyBooks := strategy(userID, remaining, int(req.TagFilter))
		if len(strategyBooks) > 0 {
			take := int(math.Min(float64(remaining), float64(len(strategyBooks))))
			recommendedBooks = append(recommendedBooks, strategyBooks[:take]...)
			remaining -= take
			log.Printf("策略 [%T] 追加: 实际追加 %d 本，剩余 %d", strategy, take, remaining)
		}
	}

	log.Printf("最终推荐数量: %d", len(recommendedBooks))

	response := gin.H{
		"code":    200,
		"message": "success",
		"data": gin.H{
			"Books": recommendedBooks,
		},
	}
	c.JSON(http.StatusOK, response)
}

// 获取用户最近的评分、关注、浏览、评论行为，综合插入新的推荐数据
// 定义用户相似度计算函数
func calculateUserSimilarity(currentUserID uint, ratings []models.Rating) []uint {
	// 获取所有用户评分数据
	allRatings, err := models.GetAllUserRatings()
	if err != nil {
		return []uint{}
	}

	// 构建评分矩阵
	userVectors := make(map[uint]map[uint]float64)
	for _, r := range allRatings {
		if userVectors[r.UserID] == nil {
			userVectors[r.UserID] = make(map[uint]float64)
		}
		userVectors[r.UserID][r.BookID] = float64(r.RatingValue)
	}

	// 获取当前用户评分向量
	currentUserVec := make(map[uint]float64)
	for _, r := range ratings {
		currentUserVec[r.BookID] = float64(r.RatingValue)
	}

	// 计算相似度
	similarities := make([]struct {
		userID uint
		score  float64
	}, 0)

	for userID, vec := range userVectors {
		if userID == currentUserID {
			continue
		}

		dotProduct := 0.0
		normA := 0.0
		normB := 0.0

		// 合并所有书籍ID
		allBooks := make(map[uint]bool)
		for bookID := range currentUserVec {
			allBooks[bookID] = true
		}
		for bookID := range vec {
			allBooks[bookID] = true
		}

		for bookID := range allBooks {
			a := currentUserVec[bookID]
			b := vec[bookID]
			dotProduct += a * b
			normA += a * a
			normB += b * b
		}

		if normA == 0 || normB == 0 {
			continue
		}
		similarity := dotProduct / (math.Sqrt(normA) * math.Sqrt(normB))
		similarities = append(similarities, struct {
			userID uint
			score  float64
		}{userID: userID, score: similarity})
	}

	// 取前5个最高相似度用户
	sort.Slice(similarities, func(i, j int) bool {
		return similarities[i].score > similarities[j].score
	})

	result := make([]uint, 0, 5)
	for i, s := range similarities {
		if i >= 5 {
			break
		}
		result = append(result, s.userID)
	}
	return result
}

// 获取已评分书籍ID列表
func getRatedBooks(ratings []models.Rating) []uint {
	bookIDs := make([]uint, 0, len(ratings))
	for _, r := range ratings {
		bookIDs = append(bookIDs, r.BookID)
	}
	return bookIDs
}

func recommendByCF(userID uint, tagFilter int) []models.Book {
	// 获取基础数据
	ratings, _ := models.GetUserRatings(userID)
	viewBookIDs, _ := models.GetUserViews(userID)
	followedUsers, _ := models.GetFollowedUsers(userID)

	// 基于评分计算用户相似度
	similarUsers := calculateUserSimilarity(userID, ratings)

	// 排除已评分书籍
	ratedBookIDs := getRatedBooks(ratings)

	// 获取相似用户集合（关注用户+评分相似用户）
	if len(similarUsers) < 5 && len(followedUsers) > 0 {
		similarUsers = append(similarUsers, followedUsers[:int(math.Min(5, float64(len(followedUsers))))]...)
	}

	// 获取相似用户的评分书籍
	// 获取未评分且高相似用户的推荐书籍
	excludedBooks := append(ratedBookIDs, viewBookIDs...)
	cfBooks, err := models.FindSimilarUsers(userID, similarUsers, excludedBooks)
	if err != nil {
		return nil
	}

	// 应用标签过滤
	if tagFilter > 0 {
		var filtered []models.Book
		for _, b := range cfBooks {
			var count int64
			models.DB.Table("book_tags").
				Where("book_id = ? AND tag_id = ?", b.BookID, tagFilter).
				Count(&count)
			if count > 0 {
				filtered = append(filtered, b)
			}
		}
		cfBooks = filtered
	}

	return cfBooks
}

func recommendByEvent() []models.Book {
	return nil
}
