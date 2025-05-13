package models

import "time"

type Tag struct {
	TagID     uint      `gorm:"primaryKey;autoIncrement"`
	TagName   string    `gorm:"size:50;not null;unique" json:"tag_name"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP"`
}

type BookTag struct {
	BookID uint `gorm:"not null"`
	TagID  uint `gorm:"not null"`
}

type TopicTag struct {
	TopicID uint `gorm:"not null"`
	TagID   uint `gorm:"not null"`
}

func GetAllTags() ([]Tag, error) {
	var tags []Tag
	if err := DB.Find(&tags).Error; err != nil {
		return nil, err
	}
	return tags, nil
}

func CreateTag(tag *Tag) error {
	if err := DB.Create(tag).Error; err != nil {
		return err
	}
	return nil
}

func DeleteTag(id string) error {
	if err := DB.Where("tag_id = ?", id).Delete(&Tag{}).Error; err != nil {
		return err
	}
	return nil
}

// 书籍类目
// var AllBookTags = []Tag{
// 	{1, "小说"},
// 	{2, "散文"},
// 	{3, "诗歌"},
// 	{4, "计算机"},
// 	{5, "工程"},
// 	{6, "医学"},
// 	{7, "中国史"},
// 	{8, "世界史"},
// 	{9, "考古"},
// 	{10, "绘画"},
// 	{11, "音乐"},
// 	{12, "建筑"},
// 	{13, "中国哲学"},
// 	{14, "西方哲学"},
// 	{15, "伦理学"},
// 	{16, "其他"},
// }
