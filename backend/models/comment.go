package models

type Comment struct {
	CommentID uint   `gorm:"primaryKey;column:comment_id" json:"commentId"`
	TargetID  uint   `gorm:"column:target_id" json:"targetId"`
	Type      uint8  `gorm:"column:type" json:"type"` // 0:书籍评论 1:话题评论
	Content   string `gorm:"column:content" json:"content"`
	UserID    uint   `gorm:"column:user_id" json:"userId"`
	CreatedAt int64  `gorm:"column:created_at" json:"createdAt"`
	User      User   `gorm:"foreignKey:UserID" json:"user"`
}

func (Comment) TableName() string {
	return "comment"
}
