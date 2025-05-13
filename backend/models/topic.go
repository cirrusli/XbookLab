package models

import (
	"math/rand"
	"time"
)

// Topic 话题模型
type Topic struct {
	TopicID      uint      `gorm:"primaryKey;autoIncrement"`
	Title        string    `gorm:"size:255;not null"`
	Content      string    `gorm:"type:text"`
	AuthorUserID uint      `gorm:"index"`
	LikeCount    uint      `gorm:"default:0"`
	TagName      string    `gorm:"-"` // 不映射到数据库
	IsLiked      int      `gorm:"-" json:"isLiked"` // 用于前端显示是否已点过赞
	CreatedAt    time.Time `gorm:"default:CURRENT_TIMESTAMP"`
}

// GetRandomTopics 获取随机话题
func GetRandomTopics(limit int) []uint {
	var ids []uint

	// 先获取所有话题ID
	DB.Model(&Topic{}).Pluck("id", &ids)

	// 如果话题数量不足，直接返回
	if len(ids) <= limit {
		return ids
	}

	// 随机选择
	rand.Shuffle(len(ids), func(i, j int) {
		ids[i], ids[j] = ids[j], ids[i]
	})

	return ids[:limit]
}
