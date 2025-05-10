package models

import (
	"time"
)

type Follow struct {
	FollowerID uint      `gorm:"primaryKey;column:follower_id"`
	FollowedID uint      `gorm:"primaryKey;column:followed_id"`
	CreatedAt  time.Time `gorm:"column:created_at"`

	Follower User `gorm:"foreignKey:FollowerID"`
	Followed User `gorm:"foreignKey:FollowedID"`
}

func (Follow) TableName() string {
	return "follow"
}
