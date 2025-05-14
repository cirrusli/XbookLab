package models

import (
	"time"

	"gorm.io/gorm"
)

// 同tools/import_recommend.go
type Recommendation struct {
	UserID  uint
	BookID  uint
	Score   float64
	Updated time.Time
}

func GetRecommendedBooks(userID uint, limit int, offset int, tagFilter int) ([]Book, error) {
	var books []Book
	query := DB.Table("recommend").
		Select("books.*, recommend.score").
		Joins("JOIN books ON recommend.book_id = books.book_id").
		Where("recommend.user_id = ?", userID)

	if tagFilter > 0 {
		query = query.Joins("JOIN book_tags ON recommend.book_id = book_tags.book_id").
			Where("book_tags.tag_id = ?", tagFilter)
	}

	err := query.Order("recommend.score DESC").Offset(offset).Limit(limit).Find(&books).Error
	return books, err
}

func (r *Recommendation) TableName() string {
	return "recommend"
}

func UpsertRecommendScore(userID uint, bookID uint, score float64) error {
	return DB.Transaction(func(tx *gorm.DB) error {
		err := tx.Exec(`
			INSERT INTO recommend (user_id, book_id, score, updated)
			VALUES (?, ?, ?, ?)
			ON DUPLICATE KEY UPDATE score = score + ?, updated = ?
		`, userID, bookID, score, time.Now(), score, time.Now()).Error
		return err
	})
}

// 获取用户评分记录
func GetUserRatings(userID uint) ([]Rating, error) {
	var ratings []Rating
	rows, err := DB.Table("ratings").
		Select("book_id, rating_value").
		Where("user_id = ?", userID).Rows()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var bookID uint
		var rating uint
		if err := rows.Scan(&bookID, &rating); err != nil {
			return nil, err
		}
		ratings = append(ratings, Rating{BookID: bookID, RatingValue: rating})
	}
	return ratings, nil
}

// GetAllUserRatings 获取所有用户的评分数据
func GetAllUserRatings() ([]Rating, error) {
	var ratings []Rating
	err := DB.Select("user_id, book_id, rating_value").Find(&ratings).Error
	if err != nil {
		return nil, err
	}
	return ratings, nil
}

// 获取用户浏览记录
func GetUserViews(userID uint) ([]uint, error) {
	var views []uint
	err := DB.Table("book_views").
		Select("DISTINCT book_id").
		Where("user_id = ?", userID).
		Pluck("book_id", &views).Error
	return views, err
}

// 获取关注用户
func GetFollowedUsers(userID uint) ([]uint, error) {
	var users []uint
	err := DB.Table("follow").
		Select("followed_id").
		Where("follower_id = ?", userID).
		Pluck("followed_id", &users).Error
	return users, err
}

// FindSimilarUsers 实现协同过滤相似用户查询
func FindSimilarUsers(userID uint, similarUsers []uint, excludeBooks []uint) ([]Book, error) {
	var books []Book
	err := DB.Table("ratings").
		Select("DISTINCT ratings.book_id, books.*").
		Joins("JOIN books ON ratings.book_id = books.book_id").
		Where("ratings.user_id IN (?) AND ratings.rating_value >= 7", similarUsers).
		Where("ratings.book_id NOT IN (SELECT book_id FROM ratings WHERE user_id = ?)", userID).
		Where("ratings.book_id NOT IN (?)", excludeBooks).
		Order("ratings.rating_value DESC").
		Limit(20).
		Find(&books).Error
	return books, err
}
