package models

type TopicLike struct {
	LikeID    uint  `gorm:"primaryKey;column:like_id" json:"likeId"`
	UserID    uint  `gorm:"column:user_id" json:"userId"`
	TopicID   uint  `gorm:"column:topic_id" json:"topicId"`
	CreatedAt int64 `gorm:"column:created_at" json:"createdAt"`
	User      User  `gorm:"foreignKey:UserID" json:"user"`
	Topic     Topic `gorm:"foreignKey:TopicID" json:"topic"`
}

func (TopicLike) TableName() string {
	return "topic_like"
}
