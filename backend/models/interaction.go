package models

import (
	"time"

	"gorm.io/gorm"
)

var MockGetPopularBooks func(int) []Book
var MockGetRandomTopics func(int) []Topic
var MockGetTopicRecommendations func(uint, int) ([]uint, error)

// 用户交互记录（包含书籍和话题）
type UserInteraction struct {
	gorm.Model
	UserID       uint    `gorm:"index"`
	ItemType     string  `gorm:"type:varchar(20);index"` // 枚举值: book/topic
	ItemID       uint    `gorm:"index"`
	Rating       float64 `gorm:"type:decimal(3,1)"`
	ViewCount    uint
	CategoryTags []string `gorm:"type:json"` // 关联的类目标签
}

// 推荐结果集
type Recommendation struct {
	gorm.Model
	UserID    uint   `gorm:"index"`
	BookIDs   []uint `gorm:"type:json"`
	TopicIDs  []uint `gorm:"type:json"`
	Algorithm string `gorm:"type:varchar(50)"` // 算法类型
	Score     float64
	ExpiredAt time.Time
}

// 获取用户相似度矩阵（包含话题数据）
func getUserSimilarityMatrix() map[uint]map[uint]float64 {
	var interactions []UserInteraction
	DB.Find(&interactions)

	matrix := make(map[uint]map[uint]float64)
	for _, i := range interactions {
		if _, ok := matrix[i.UserID]; !ok {
			matrix[i.UserID] = make(map[uint]float64)
		}
		// 综合评分和浏览行为
		score := i.Rating + float64(i.ViewCount)*0.1
		matrix[i.UserID][i.ItemID] = score
	}
	return matrix
}

// 获取话题推荐（基于标签匹配）
func GetTopicRecommendations(userID uint, limit int) ([]uint, error) {
	if MockGetTopicRecommendations != nil {
		return MockGetTopicRecommendations(userID, limit)
	}

	var userTags []string
	DB.Model(UserInteraction{}).
		Where("user_id = ?", userID).
		Pluck("DISTINCT category_tags", &userTags)

	var topics []Topic
	DB.Where("category_tags && ?", userTags).
		Order("heat_score DESC").
		Limit(limit).
		Find(&topics)

	var ids []uint
	for _, t := range topics {
		ids = append(ids, t.ID)
	}
	return ids, nil
}

// 统一推荐入口
func GenerateRecommendations(userID uint) Recommendation {
	bookIDs, _ := GetUserBasedCFRecommendations(userID, 10, &GormUserBehaviorRepository{db: DB})
	var ids []uint
	for _, b := range bookIDs {
		ids = append(ids, b.ID)
	}

	topicIDs, err := GetTopicRecommendations(userID, 6)
	if err != nil {
		topicIDs = []uint{}
	}

	return Recommendation{
		UserID:    userID,
		BookIDs:   ids,
		TopicIDs:  topicIDs,
		Algorithm: "hybrid",
		ExpiredAt: time.Now().Add(24 * time.Hour),
	}
}

// 用户行为仓库接口
type UserBehaviorRepository interface {
	GetAll() ([]UserInteraction, error)
	Create(behavior UserInteraction) error
}

// GORM用户行为仓库实现
type GormUserBehaviorRepository struct {
	db *gorm.DB
}

func (r *GormUserBehaviorRepository) GetAll() ([]UserInteraction, error) {
	var behaviors []UserInteraction
	err := r.db.Find(&behaviors).Error
	return behaviors, err
}

func (r *GormUserBehaviorRepository) Create(behavior UserInteraction) error {
	return r.db.Create(&behavior).Error
}

func (u *UserInteraction) AfterSave(tx *gorm.DB) (err error) {
	var tags []string
	// 获取关联标签
	if u.ItemType == "book" {
		var book Book
		if err := tx.First(&book, u.ItemID).Error; err != nil {
			return err
		}
		tags = book.CategoryTags
	} else if u.ItemType == "topic" {
		var topic Topic
		if err := tx.First(&topic, u.ItemID).Error; err != nil {
			return err
		}
		tags = topic.CategoryTags
	}

	// 创建行为记录
	behavior := UserInteraction{
		UserID:       u.UserID,
		ItemID:       u.ItemID,
		ItemType:     u.ItemType,
		Rating:       u.Rating,
		CategoryTags: tags,
	}
	return tx.Create(&behavior).Error
}

//go:generate mockgen -destination=mock_user_behavior_repository.go -package=models . UserBehaviorRepository
