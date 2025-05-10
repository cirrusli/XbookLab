package models

import (
	"time"
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
	query := DB.Table("recommend").Select("book_id, score").Where("user_id = ?", userID)

	if tagFilter > 0 {
		query = query.Joins("JOIN book_tag ON recommend.book_id = book_tag.book_id").
			Where("book_tag.tag_id = ?", tagFilter)
	}

	err := query.Order("score DESC").Offset(offset).Limit(limit).Find(&books).Error
	return books, err
}

func (r *Recommendation) TableName() string {
	return "recommend"
}
