package models

import (
	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Title       string  `gorm:"size:255;not null"`
	Author      string  `gorm:"size:100;index"`
	Description string  `gorm:"type:text"`
	Cover       string  `gorm:"size:255"`
	Rating      float64 `gorm:"type:decimal(3,1)"`
	CategoryID  uint    `gorm:"index"`
}

func GetRecommendedBooks(page, pageSize int, categoryID string) ([]Book, int64) {
	var books []Book
	var total int64
	
	query := DB.Model(&Book{})
	
	if categoryID != "" {
		query = query.Where("category_id = ?", categoryID)
	}
	
	query.Count(&total)
	query.Offset((page - 1) * pageSize).Limit(pageSize).Find(&books)
	
	return books, total
}