package models

import (
	"time"
)

// BookView 浏览记录模型
type BookView struct {
	ViewID    uint      `gorm:"primaryKey;autoIncrement"`
	UserID    uint      `gorm:"index"`
	BookID    uint      `gorm:"not null;index"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP"`
}

// Rating 评分模型
type Rating struct {
	RatingID    uint      `gorm:"primaryKey;autoIncrement"`
	UserID      uint      `gorm:"uniqueIndex:user_book"`
	BookID      uint      `gorm:"uniqueIndex:user_book"`
	RatingValue uint      `gorm:"not null"`
	CreatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"`
}

// CreateBookView 创建浏览记录
func CreateBookView(userID, bookID uint) error {
	view := BookView{
		UserID: userID,
		BookID: bookID,
	}
	return DB.Create(&view).Error
}

// CreateOrUpdateRating 创建或更新评分
func CreateOrUpdateRating(userID, bookID, ratingValue uint) error {
	rating := Rating{
		UserID:      userID,
		BookID:      bookID,
		RatingValue: ratingValue,
	}
	return DB.Where(Rating{UserID: userID, BookID: bookID}).
		Assign(Rating{RatingValue: ratingValue}).
		FirstOrCreate(&rating).Error
}
