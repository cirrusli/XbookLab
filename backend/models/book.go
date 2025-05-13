package models

import (
	"time"
)

type Book struct {
	BookID        uint      `gorm:"primaryKey;autoIncrement"`
	Title         string    `gorm:"size:255;not null"`
	Author        string    `gorm:"size:100"`
	Cover         string    `gorm:"size:255"`
	Description   string    `gorm:"type:text"`
	AverageRating float64   `gorm:"type:decimal(3,1);default:0.0"`
	CreatedAt     time.Time `gorm:"default:CURRENT_TIMESTAMP"`
}

// GetPopularBooks 获取高分书籍
func GetPopularBooks(limit int, tagFilter int) []Book {
	var books []Book
	query := DB.Order("average_rating desc").Limit(limit)

	if tagFilter > 0 {
		query = query.Joins("JOIN book_tags ON books.book_id = book_tags.book_id").
			Where("book_tags.tag_id = ?", tagFilter)
	}

	query.Find(&books)
	return books
}

// GetBooks 获取书籍列表
func GetBooks(limit int) []Book {
	var books []Book
	DB.Limit(limit).Find(&books)
	return books
}
