package models

type TopicLike struct {
	LikeID    uint  `gorm:"primaryKey;column:like_id" json:"like_id"`
	UserID    uint  `gorm:"column:user_id" json:"user_id"`
	TopicID   uint  `gorm:"column:topic_id" json:"topic_id"`
	CreatedAt int64 `gorm:"column:created_at" json:"created_at"`
	User      User  `gorm:"foreignKey:UserID" json:"user"`
	Topic     Topic `gorm:"foreignKey:TopicID" json:"topic"`
}

func (TopicLike) TableName() string {
	return "topic_like"
}
