package models

import (
	"math/rand"

	"gorm.io/gorm"
)

// Topic 话题模型
type Topic struct {
	gorm.Model
	Title      string  `gorm:"size:255;not null"`
	Content    string  `gorm:"type:text;not null"`
	AuthorID   uint    `gorm:"not null"`
	ViewCount  uint    `gorm:"default:0"`
	Tag        string  `gorm:"size:255;not null"`
	HeatScore  float64 `gorm:"type:decimal(4,1);default:0.0"` // 热度评分（根据浏览、回复等计算）
	ReplyCount uint    `gorm:"default:0"`
	Status     string  `gorm:"size:20;default:'active'"`
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
