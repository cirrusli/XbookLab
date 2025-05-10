package models

type Event struct {
	EventID      uint   `gorm:"primaryKey;column:event_id" json:"eventId"`
	UserID       uint   `gorm:"column:user_id" json:"userId"`
	EventType    uint8  `gorm:"column:event_type" json:"eventType"` // 0:阅读 1:评论 2:评分 3:点赞话题
	EventContent string `gorm:"column:event_content" json:"eventContent"`
	TargetID     uint   `gorm:"column:target_id" json:"targetId"`
	TargetType   uint8  `gorm:"column:target_type" json:"targetType"` // 0:书籍 1:话题
	CreatedAt    int64  `gorm:"column:created_at" json:"createdAt"`
	User         User   `gorm:"foreignKey:UserID" json:"user"`
}

func (Event) TableName() string {
	return "event"
}
