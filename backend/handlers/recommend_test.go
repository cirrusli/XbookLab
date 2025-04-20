package handlers

import (
	"net/http/httptest"
	"testing"
	"xbooklab/models"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func TestRecommend(t *testing.T) {
	defer func() { models.MockGetTopicRecommendations = nil }()
	dsn := "root:lzq@tcp(localhost:3306)/x_book_lab?charset=utf8mb4&parseTime=True&loc=Local"
	testDB, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	models.DB = testDB
	testDB.AutoMigrate(&models.Topic{})
	t.Run("冷启动用户推荐", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockRepo := models.NewMockUserBehaviorRepository(ctrl)
		mockRepo.EXPECT().GetAll().Return([]models.UserInteraction{}, nil)

		models.MockGetPopularBooks = func(int) []models.Book {
			return []models.Book{{Title: "三体"}, {Title: "乡土中国"}}
		}
		models.MockGetRandomTopics = func(int) []models.Topic {
			return []models.Topic{{Title: "历史话题"}, {Title: "文学话题"}}
		}
		models.MockGetTopicRecommendations = func(uint, int) ([]uint, error) { return []uint{101, 102}, nil }

		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Set("user_id", uint(999))

		Recommend(c, mockRepo)
		assert.Equal(t, 200, w.Code)
		models.MockGetTopicRecommendations = nil
	})

	t.Run("正常用户推荐", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockRepo := models.NewMockUserBehaviorRepository(ctrl)
		mockRepo.EXPECT().GetAll().Return([]models.UserInteraction{{UserID: 1}}, nil)

		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Set("user_id", uint(1))

		Recommend(c, mockRepo)
		assert.Equal(t, 200, w.Code)
	})

	t.Run("无效用户ID", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockRepo := models.NewMockUserBehaviorRepository(ctrl)

		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Set("user_id", uint(0))

		Recommend(c, mockRepo)
		assert.Equal(t, 400, w.Code)
	})
}

func TestCalculateSimilarity(t *testing.T) {
	t.Run("完全相似用户", func(t *testing.T) {
		user1 := map[uint]float64{1: 5, 2: 4}
		user2 := map[uint]float64{1: 5, 2: 4}
		sim := cosineSimilarity(user1, user2)
		assert.InDelta(t, 1.0, sim, 0.01)
	})

	t.Run("完全不同用户", func(t *testing.T) {
		user1 := map[uint]float64{1: 5}
		user2 := map[uint]float64{2: 5}
		sim := cosineSimilarity(user1, user2)
		assert.Equal(t, 0.0, sim)
	})
}
func TestGetTopKSimilarUsers(t *testing.T) {
	t.Run("正常排序", func(t *testing.T) {
		similarities := map[uint]float64{
			1: 0.9,
			2: 0.8,
			3: 0.95,
		}
		result := getTopKSimilarUsers(similarities, 2)
		assert.Equal(t, []uint{3, 1}, result)
	})
}
