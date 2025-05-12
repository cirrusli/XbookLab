package models

import "time"

type Comment struct {
	CommentID uint      `gorm:"primaryKey;column:comment_id" json:"comment_id"`
	TargetID  uint      `gorm:"column:target_id" json:"target_id"`
	Type      uint8     `gorm:"column:type" json:"type"` // 0:书籍评论 1:话题评论
	Content   string    `gorm:"column:content" json:"content"`
	UserID    uint      `gorm:"column:user_id" json:"user_id"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
	User      User      `gorm:"foreignKey:UserID" json:"user"`
}

func (Comment) TableName() string {
	return "comment"
}

func CreateComment(comment *Comment) error {
	if err := DB.Create(comment).Error; err != nil {
		return err
	}

	// 记录用户动态
	event := Event{
		UserID:       comment.UserID,
		EventType:    1, // 评论
		EventContent: "发表了评论",
		TargetID:     comment.TargetID,
		TargetType:   comment.Type,
	}
	return DB.Create(&event).Error
}

func GetComments(targetID, targetType uint) ([]Comment, error) {
	var comments []Comment
	err := DB.Where("target_id = ? AND type = ?", targetID, targetType).
		Preload("User").Order("created_at DESC").Find(&comments).Error
	return comments, err
}

func DeleteComment(commentID uint) error {
	var comment Comment
	if err := DB.Where("comment_id = ?", commentID).First(&comment).Error; err != nil {
		return err
	}
	return DB.Delete(&comment).Error
}
