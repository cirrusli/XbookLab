package main

// 协同过滤推荐算法实现
// 涉及MySQL表说明：
// - users: 存储用户基础信息，用于获取所有用户ID（字段：user_id）
// - ratings: 存储用户对书籍的评分数据（字段：user_id, book_id, rating_value）
// - book_views: 存储用户浏览书籍的记录（字段：user_id, book_id）
// - user_tags: 存储用户兴趣标签关联关系（字段：user_id, tag_id）
// - follow: 存储用户关注关系（字段：follower_id, followed_id）
// - book_tags: 存储书籍与标签的关联关系（字段：book_id, tag_id）
// - recommend: 存储最终推荐结果（字段：user_id, book_id, score, update_at）
//
// 用户视角数据流转说明：
// 1. 用户信息获取：从users表获取所有用户ID，针对每个用户从ratings获取评分记录、从book_views获取浏览记录、从user_tags获取兴趣标签、从follow获取关注用户
// 2. 相似用户计算：基于用户标签（Jaccard相似度）和评分（余弦相似度）计算与目标用户的相似度，筛选相似度>0.3的前10名用户
// 3. 推荐书籍生成：
//    a. 合并相似用户与关注用户作为推荐来源，提取来源用户评分≥7的书籍（排除用户已评分/浏览的书籍）
//    b. 结合用户兴趣标签，从book_tags表提取标签匹配度高的书籍（排除用户已交互书籍）
//    c. 补充探索性推荐：选取用户未接触但全局评分≥8的书籍（标签与用户兴趣无交集）
// 4. 结果存储：合并多来源推荐分数，生成最终推荐结果并批量插入recommend表

import (
	"database/sql"
	"fmt"
	"log"
	"math"
	"sort"
	"strings"
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

	// 8. 整合探索性推荐
	exploreBooks := getExplorationBooks(db, userID)
	// 累加探索性推荐分数
	for bookID, score := range exploreBooks {
		recommendedBooks[bookID] += score
	}
	// 生成推荐结果
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
	rows, err := db.Query("SELECT book_id, rating_value FROM ratings WHERE user_id = ?", userID)
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
	rows, err := db.Query("SELECT DISTINCT book_id FROM book_views WHERE user_id = ?", userID)
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
	rows, err := db.Query("SELECT tag_id FROM user_tags WHERE user_id = ?", userID)
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

// 计算相似用户（协同过滤核心逻辑）
// 参数：
//
//	db: 数据库连接
//	userID: 目标用户ID
//	userRatings: 目标用户评分数据（bookID -> rating）
//	userViews: 目标用户浏览记录（bookID列表）
//	userTags: 目标用户兴趣标签（tagID列表）
//
// 返回：
//
//	高相似度用户ID列表（前10名）
//
// 算法说明：
//  1. 获取所有其他用户ID
//  2. 对每个用户计算标签相似度（Jaccard）和评分相似度（余弦）
//  3. 综合相似度 = 标签相似度*0.4 + 评分相似度*0.6
//  4. 筛选相似度>0.3的用户，按相似度排序取前10
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
		combinedScore := tagSimilarity*0.3 + ratingSimilarity*0.7

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

// 计算Jaccard相似度（用于标签相似性度量）
// 参数：
//
//	setA: 标签集合A
//	setB: 标签集合B
//
// 返回：
//
//	Jaccard相似度值（0-1之间，值越大越相似）
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

// 计算余弦相似度（用于评分向量相似性度量）
// 参数：
//
//	vecA: 评分向量A（bookID -> rating）
//	vecB: 评分向量B（bookID -> rating）
//
// 返回：
//
//	余弦相似度值（0-1之间，值越大越相似）
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
		placeholders := make([]string, len(sourceUsers))
		for i := range sourceUsers {
			placeholders[i] = "?"
		}
		// 从相似用户中获取高评分书籍（暂不排除用户已评分/浏览的书籍）
		query := fmt.Sprintf(`
			SELECT r.book_id, AVG(r.rating_value) as avg_rating
			FROM ratings r
			WHERE r.user_id IN (%s)  -- 相似用户ID列表
				AND r.rating_value >= 7  -- 仅考虑评分≥7的书籍
				-- AND r.book_id NOT IN (  -- 排除用户已交互的书籍（评分或浏览）
				--	SELECT book_id FROM ratings WHERE user_id = ?
				--	UNION
				--	SELECT book_id FROM book_views WHERE user_id = ?
				-- )
			GROUP BY r.book_id  -- 按书籍分组计算平均评分
			ORDER BY avg_rating DESC  -- 按平均评分降序排列
			LIMIT 20  -- 取前20本
		`, strings.Join(placeholders, ","))
		args := make([]interface{}, 0, len(sourceUsers)+2) // 2是上面的排除已交互的两个占位符
		for _, id := range sourceUsers {
			args = append(args, id)
		}
		// args = append(args, userID, userID) // 上面sql的注释部分的两个占位符
		rows, err := db.Query(query, args...)
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
		// 构建参数化查询避免SQL注入
		placeholders := make([]string, len(userTags))
		for i := range userTags {
			placeholders[i] = "?"
		}
		query := fmt.Sprintf(`
			SELECT b.book_id, COUNT(*) as tag_count
			FROM book_tags b
			WHERE b.tag_id IN (%s)
			AND b.book_id NOT IN (
				SELECT book_id FROM ratings WHERE user_id = ?
				UNION
				SELECT book_id FROM book_views WHERE user_id = ?
			)
			GROUP BY b.book_id
			ORDER BY tag_count DESC
			LIMIT 20
		`, strings.Join(placeholders, ","))
		args := make([]interface{}, 0, len(userTags)+2)
		for _, id := range userTags {
			args = append(args, id)
		}
		args = append(args, userID, userID)
		rows, err := db.Query(query, args...)
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
		FROM book_tags bt
		JOIN (
			SELECT book_id, AVG(rating_value) as avg_rating
			FROM ratings
			GROUP BY book_id
			HAVING avg_rating >= 8.0
		) r ON bt.book_id = r.book_id
		WHERE bt.tag_id NOT IN (
			SELECT tag_id FROM user_tags WHERE user_id = ?
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
	if len(recommendations) == 0 {
		log.Fatalln("recommendations is empty")
		return
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
