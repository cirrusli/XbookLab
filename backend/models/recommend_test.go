package models

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestGetUserBasedCFRecommendations(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockRepo := NewMockUserInteractionRepo(ctrl)

	t.Run("新用户冷启动推荐", func(t *testing.T) {
		mockRepo.EXPECT().GetAll().Return([]UserInteraction{}, nil)
		// MockGetPopularRecommendations = func(int) []Book {
		// 	return []Book{{ID: 1001}, {ID: 1002}}
		// }
		// defer func() { MockGetPopularRecommendations = nil }()

		result, err := GetUserBasedCFRecommendations(999, 5, mockRepo)
		assert.NoError(t, err)
		assert.NotEmpty(t, result)
	})

	t.Run("存在相似用户推荐", func(t *testing.T) {
		mockBehaviors := []UserInteraction{{
			UserID: 1,
			ItemID: 1001,
			Rating: 5,
		}, {
			UserID: 2,
			ItemID: 1001,
			Rating: 4.5,
		}, {
			UserID: 3,
			ItemID: 1002,
			Rating: 4,
		}}

		mockRepo.EXPECT().GetAll().Return(mockBehaviors, nil)
		// MockGetBookByID = func(id uint) Book {
		// 	return Book{Title: "三体"}
		// }
		// defer func() { MockGetBookByID = nil }()

		result, err := GetUserBasedCFRecommendations(1, 3, mockRepo)
		assert.NoError(t, err)
		assert.Len(t, result, 3)
	})
}

func TestCalculateSimilarity(t *testing.T) {
	t.Run("余弦相似度计算", func(t *testing.T) {
		userA := map[uint]float64{1: 5, 2: 3}
		userB := map[uint]float64{1: 4, 2: 4}

		sim := calculateSimilarity(userA, userB)
		assert.InDelta(t, 0.98, sim, 0.01)
	})
}
func TestGetTopNItems(t *testing.T) {
	t.Run("排序取TopN", func(t *testing.T) {
		items := map[uint]float64{
			1: 9.5,
			2: 8.2,
			3: 7.8,
		}
		result := getTopNItems(items, 2)
		assert.Len(t, result, 2)
		assert.Equal(t, uint(1), result[0].ID)
	})
}
