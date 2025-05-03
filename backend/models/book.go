package models

import (
	"strconv"

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
	ViewCount   uint    `gorm:"default:0"`
	Tag         string  `gorm:"size:255;not null"`
}

// GetPopularBooks 获取热门书籍
func GetPopularBooks(limit int) []uint {
	var books []Book
	DB.Order("view_count desc").Limit(limit).Find(&books)

	var ids []uint
	for _, book := range books {
		ids = append(ids, book.ID)
	}
	return ids
}

func GetRecommendedBooks(page, pageSize int, categoryID string, userID uint) ([]Book, int64) {
	var books []Book
	var total int64

	// 如果有登录用户，使用协同过滤推荐
	if userID > 0 {
		recommendedBooks, err := GetUserBasedCFRecommendations(userID, pageSize, &GormUserInteractionRepo{db: DB})
		if err == nil && len(recommendedBooks) > 0 {
			books = recommendedBooks
			total = int64(len(books))
			return books, total
		}
	}

	// 如果没有登录用户或推荐失败，使用默认推荐(按评分排序)
	query := DB.Model(&Book{}).Order("rating DESC")

	if categoryID != "" {
		cid, _ := strconv.Atoi(categoryID)
		query = query.Where("category_id = ?", cid)
	}

	query.Count(&total)
	query.Offset((page - 1) * pageSize).Limit(pageSize).Find(&books)

	return books, total
}
