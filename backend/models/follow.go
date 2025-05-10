package models

import (
	"errors"
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

func CreateFollow(followerID, followedID uint) error {
	// 不能关注自己
	if followerID == followedID {
		return ErrCannotFollowSelf
	}

	// 检查是否已关注
	var existingFollow Follow
	if err := DB.Where("follower_id = ? AND followed_id = ?", followerID, followedID).First(&existingFollow).Error; err == nil {
		return ErrAlreadyFollowed
	}

	// 创建关注关系
	follow := Follow{
		FollowerID: followerID,
		FollowedID: followedID,
	}
	return DB.Create(&follow).Error
}

func DeleteFollow(followerID, followedID uint) error {
	return DB.Where("follower_id = ? AND followed_id = ?", followerID, followedID).Delete(&Follow{}).Error
}

func GetFollowing(userID uint) ([]Follow, error) {
	var follows []Follow
	err := DB.Preload("Followed").Where("follower_id = ?", userID).Find(&follows).Error
	return follows, err
}

func GetFollowers(userID uint) ([]Follow, error) {
	var follows []Follow
	err := DB.Preload("Follower").Where("followed_id = ?", userID).Find(&follows).Error
	return follows, err
}

var (
	ErrCannotFollowSelf = errors.New("不能关注自己")
	ErrAlreadyFollowed  = errors.New("已关注该用户")
)
