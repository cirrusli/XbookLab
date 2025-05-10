package main

/*
协同过滤推荐算法实现说明

本文件实现了一个基于用户行为的协同过滤推荐系统，主要包含以下核心组件：

1. 算法流程概述：
   - 为每个用户并发计算个性化推荐
   - 结合用户评分、浏览记录、标签和社交关系等多维度数据
   - 使用加权混合策略生成最终推荐结果

2. 相似度计算方法：
   - Jaccard相似度：基于用户标签计算兴趣相似度
   - 余弦相似度：基于用户评分向量计算偏好相似度
   - 综合相似度：加权平均上述两种相似度(标签40% + 评分60%)

3. 推荐来源：
   - 用户相似度推荐：基于相似用户的高评分书籍
   - 社交关系推荐：基于关注用户的高评分书籍
   - 标签匹配推荐：基于用户兴趣标签的热门书籍

4. 评分机制：
   - 相似用户推荐权重：60%
   - 标签匹配推荐权重：40%
   - 最终得分归一化处理

5. 并发处理：
   - 使用goroutine并发处理每个用户的推荐计算
   - 通过channel收集结果
   - 使用WaitGroup同步并发任务
*/

import (
	"database/sql"
	"fmt"
	"log"
	"math"
	"sort"
	"sync"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

// 推荐结果结构体
type Recommendation struct {
	UserID  uint
	BookID  uint
	Score   float64
	Updated time.Time
}

func main() {
	// 数据库连接
	db, err := sql.Open("mysql", "root:lzq@tcp(localhost:3306)/x_book_lab_test")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// 获取所有用户ID
	userIDs := getAllUserIDs(db)

	// 使用WaitGroup等待所有goroutine完成
	var wg sync.WaitGroup
	resultChan := make(chan []Recommendation, len(userIDs))

	// 为每个用户并发计算推荐
	for _, userID := range userIDs {
		wg.Add(1)
		go func(id uint) {
			defer wg.Done()
			recommendations := calculateUserRecommendations(db, id)
			resultChan <- recommendations
		}(userID)
	}

	// 启动goroutine等待所有计算完成并关闭channel
	go func() {
		wg.Wait()
		close(resultChan)
	}()

	// 收集所有推荐结果
	var allRecommendations []Recommendation
	for recs := range resultChan {
		allRecommendations = append(allRecommendations, recs...)
	}

	// 批量插入推荐结果
	insertRecommendations(db, allRecommendations)
	fmt.Println("推荐数据计算完成")
}

// 获取所有用户ID
func getAllUserIDs(db *sql.DB) []uint {
	var userIDs []uint
	rows, err := db.Query("SELECT user_id FROM users")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id uint
		if err := rows.Scan(&id); err != nil {
			log.Fatal(err)
		}
		userIDs = append(userIDs, id)
	}
	return userIDs
}

// 计算单个用户的推荐
func calculateUserRecommendations(db *sql.DB, userID uint) []Recommendation {
	// 1. 获取用户评分数据
	userRatings := getUserRatings(db, userID)
	// 2. 获取用户浏览记录
	userViews := getUserViews(db, userID)
	// 3. 获取用户标签
	userTags := getUserTags(db, userID)
	// 4. 获取关注用户
	followedUsers := getFollowedUsers(db, userID)
	// 5. 计算相似用户
	similarUsers := calculateSimilarUsers(db, userID, userRatings, userViews, userTags)
	// 6. 合并所有潜在推荐来源
	recommendationSources := append(similarUsers, followedUsers...)
	// 7. 计算推荐书籍
	recommendedBooks := calculateRecommendedBooks(db, userID, recommendationSources, userTags)
	// 8. 生成推荐结果
	var recommendations []Recommendation
	for bookID, score := range recommendedBooks {
		recommendations = append(recommendations, Recommendation{
			UserID:  userID,
			BookID:  bookID,
			Score:   score,
			Updated: time.Now().In(time.FixedZone("CST", 8*3600)),
		})
	}

	return recommendations
}

// 获取用户评分数据
func getUserRatings(db *sql.DB, userID uint) map[uint]float64 {
	ratings := make(map[uint]float64)
	rows, err := db.Query("SELECT book_id, rating_value FROM rating WHERE user_id = ?", userID)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var bookID uint
		var rating float64
		if err := rows.Scan(&bookID, &rating); err != nil {
			log.Fatal(err)
		}
		ratings[bookID] = rating
	}
	return ratings
}

// 获取用户浏览记录
func getUserViews(db *sql.DB, userID uint) []uint {
	var views []uint
	rows, err := db.Query("SELECT DISTINCT book_id FROM book_view WHERE user_id = ?", userID)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var bookID uint
		if err := rows.Scan(&bookID); err != nil {
			log.Fatal(err)
		}
		views = append(views, bookID)
	}
	return views
}

// 获取用户标签
func getUserTags(db *sql.DB, userID uint) []uint {
	var tags []uint
	rows, err := db.Query("SELECT tag_id FROM user_tag WHERE user_id = ?", userID)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var tagID uint
		if err := rows.Scan(&tagID); err != nil {
			log.Fatal(err)
		}
		tags = append(tags, tagID)
	}
	return tags
}

// 获取关注用户
func getFollowedUsers(db *sql.DB, userID uint) []uint {
	var followed []uint
	rows, err := db.Query("SELECT followed_id FROM follow WHERE follower_id = ?", userID)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id uint
		if err := rows.Scan(&id); err != nil {
			log.Fatal(err)
		}
		followed = append(followed, id)
	}
	return followed
}

// 计算相似用户
func calculateSimilarUsers(db *sql.DB, userID uint, userRatings map[uint]float64, userViews []uint, userTags []uint) []uint {
	type UserSimilarity struct {
		UserID uint
		Score  float64
	}

	// 1. 获取所有用户ID
	allUsers := getAllUserIDs(db)
	var similarUsers []UserSimilarity

	// 2. 计算每个用户与目标用户的相似度
	for _, otherUserID := range allUsers {
		if otherUserID == userID {
			continue
		}

		// 获取其他用户的标签
		otherTags := getUserTags(db, otherUserID)

		// 计算Jaccard相似度（标签相似度）
		tagSimilarity := calculateJaccardSimilarity(userTags, otherTags)

		// 获取其他用户评分
		otherRatings := getUserRatings(db, otherUserID)

		// 计算余弦相似度（评分向量相似度）
		ratingSimilarity := calculateCosineSimilarity(userRatings, otherRatings)

		// 综合相似度（加权平均）
		combinedScore := tagSimilarity*0.4 + ratingSimilarity*0.6

		if combinedScore > 0.3 { // 相似度阈值
			similarUsers = append(similarUsers, UserSimilarity{
				UserID: otherUserID,
				Score:  combinedScore,
			})
		}
	}

	// 3. 按相似度排序并返回前10个用户
	sort.Slice(similarUsers, func(i, j int) bool {
		return similarUsers[i].Score > similarUsers[j].Score
	})

	var result []uint
	for i := 0; i < min(10, len(similarUsers)); i++ {
		result = append(result, similarUsers[i].UserID)
	}
	return result
}

// 计算Jaccard相似度
func calculateJaccardSimilarity(setA, setB []uint) float64 {
	intersection := make(map[uint]bool)
	union := make(map[uint]bool)

	for _, item := range setA {
		union[item] = true
	}

	for _, item := range setB {
		if union[item] {
			intersection[item] = true
		} else {
			union[item] = true
		}
	}

	if len(union) == 0 {
		return 0
	}
	return float64(len(intersection)) / float64(len(union))
}

// 计算余弦相似度
func calculateCosineSimilarity(vecA, vecB map[uint]float64) float64 {
	dotProduct := 0.0
	magnitudeA := 0.0
	magnitudeB := 0.0

	// 计算点积和向量A的模
	for bookID, ratingA := range vecA {
		if ratingB, exists := vecB[bookID]; exists {
			dotProduct += ratingA * ratingB
		}
		magnitudeA += ratingA * ratingA
	}

	// 计算向量B的模
	for _, ratingB := range vecB {
		magnitudeB += ratingB * ratingB
	}

	magnitudeA = math.Sqrt(magnitudeA)
	magnitudeB = math.Sqrt(magnitudeB)

	if magnitudeA == 0 || magnitudeB == 0 {
		return 0
	}
	return dotProduct / (magnitudeA * magnitudeB)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// 计算推荐书籍
func calculateRecommendedBooks(db *sql.DB, userID uint, sourceUsers []uint, userTags []uint) map[uint]float64 {
	recommendations := make(map[uint]float64)

	// 1. 从相似用户和关注用户的评分中获取推荐
	if len(sourceUsers) > 0 {
		// 将[]uint转换为逗号分隔的字符串
		userIDsStr := "("
		for i, id := range sourceUsers {
			if i > 0 {
				userIDsStr += ","
			}
			userIDsStr += fmt.Sprintf("%d", id)
		}
		userIDsStr += ")"

		rows, err := db.Query(fmt.Sprintf(`
			SELECT r.book_id, AVG(r.rating_value) as avg_rating
			FROM rating r
			WHERE r.user_id IN %s AND r.rating_value >= 7
			AND r.book_id NOT IN (
				SELECT book_id FROM rating WHERE user_id = ?
				UNION
				SELECT book_id FROM book_view WHERE user_id = ?
			)
			GROUP BY r.book_id
			ORDER BY avg_rating DESC
			LIMIT 20
		`, userIDsStr), userID, userID)
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()

		for rows.Next() {
			var bookID uint
			var rating float64
			if err := rows.Scan(&bookID, &rating); err != nil {
				log.Fatal(err)
			}
			recommendations[bookID] = rating * 0.6
		}
	}

	// 2. 从相似标签的书籍中获取推荐
	if len(userTags) > 0 {
		// 将[]uint转换为逗号分隔的字符串
		tagIDsStr := "("
		for i, id := range userTags {
			if i > 0 {
				tagIDsStr += ","
			}
			tagIDsStr += fmt.Sprintf("%d", id)
		}
		tagIDsStr += ")"

		rows, err := db.Query(fmt.Sprintf(`
			SELECT b.book_id, COUNT(*) as tag_count
			FROM book_tag b
			WHERE b.tag_id IN %s
			AND b.book_id NOT IN (
				SELECT book_id FROM rating WHERE user_id = ?
				UNION
				SELECT book_id FROM book_view WHERE user_id = ?
			)
			GROUP BY b.book_id
			ORDER BY tag_count DESC
			LIMIT 20
		`, tagIDsStr), userID, userID)
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()

		for rows.Next() {
			var bookID uint
			var count int
			if err := rows.Scan(&bookID, &count); err != nil {
				log.Fatal(err)
			}
			recommendations[bookID] += float64(count) * 0.4
		}
	}

	return recommendations
}
// 探索性推荐逻辑
func getExplorationBooks(db *sql.DB, userID uint) map[uint]float64 {
	explore := make(map[uint]float64)
	rows, err := db.Query(`
		SELECT b.book_id, r.avg_rating
		FROM book_tag bt
		JOIN (
			SELECT book_id, AVG(rating_value) as avg_rating
			FROM rating
			GROUP BY book_id
			HAVING avg_rating >= 8.0
		) r ON bt.book_id = r.book_id
		WHERE bt.tag_id NOT IN (
			SELECT tag_id FROM user_tag WHERE user_id = ?
		)
		ORDER BY r.avg_rating DESC
		LIMIT 2`, userID)
	if err == nil {
		defer rows.Close()
		for rows.Next() {
			var bookID uint
			var rating float64
			if err := rows.Scan(&bookID, &rating); err == nil {
				explore[bookID] = rating * 0.3 // 探索性权重
			}
		}
	}
	return explore
}
// 批量插入推荐结果
func insertRecommendations(db *sql.DB, recommendations []Recommendation) {
	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}

	// 先清空旧数据
	_, err = tx.Exec("TRUNCATE TABLE recommend")
	if err != nil {
		tx.Rollback()
		log.Fatal(err)
	}

	// 准备批量插入语句
	stmt, err := tx.Prepare("INSERT INTO recommend(user_id, book_id, score, update_at) VALUES(?, ?, ?, ?)")
	if err != nil {
		tx.Rollback()
		log.Fatal(err)
	}
	defer stmt.Close()

	// 执行批量插入
	for _, rec := range recommendations {
		_, err = stmt.Exec(rec.UserID, rec.BookID, rec.Score, rec.Updated)
		if err != nil {
			tx.Rollback()
			log.Fatal(err)
		}
	}

	// 提交事务
	err = tx.Commit()
	if err != nil {
		log.Fatal(err)
	}
}
