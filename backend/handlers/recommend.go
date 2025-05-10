package handlers

import (
	"log"
	"net/http"
	"xbooklab/models"

	"github.com/gin-gonic/gin"
)

type GetRecommendedBooksRequest struct {
	UserID    uint `json:"user_id"`
	Limit     uint `json:"limit"`
	Offset    uint `json:"offset"`
	TagFilter uint `json:"tag_filter"`
}

type GetRecommendedBooksResponse struct {
	Books                []models.Book `json:"books"`
	Algorithm            string        `json:"algorithm"`
	RecommendationReason string        `json:"recommendation_reason"`
}

// GetRecommendedBooks 实现书籍推荐逻辑
// 首页推荐tab，要考虑未登录状态也可以访问热门书籍，登录态则可以拿到推荐书籍
func GetRecommendedBooks(c *gin.Context) {
	// 解析请求参数
	var req GetRecommendedBooksRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	log.Printf("Received request: %+v", req)
	var recommendedBooks []models.Book
	var algorithm string
	var recommendationReason string
	var err error

	if req.UserID == 0 {
		// 未登录
		recommendedBooks = models.GetPopularBooks(int(req.Limit))
		algorithm = "PopularityBased"
		recommendationReason = "Popular books for new users"
	} else {
		recommendedBooks, err = models.GetRecommendedBooks(req.UserID, int(req.Limit), int(req.Offset), int(req.TagFilter))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		// 如果推荐书籍数量不足，用热门书籍补充
		if len(recommendedBooks) < int(req.Limit) {
			remaining := int(req.Limit) - len(recommendedBooks)
			popularBooks := models.GetPopularBooks(remaining)
			recommendedBooks = append(recommendedBooks, popularBooks...)
		}
		algorithm = "CollaborativeFiltering"
		recommendationReason = "Books recommended based on user preferences"
	}

	response := GetRecommendedBooksResponse{
		Books:                recommendedBooks,
		Algorithm:            algorithm,
		RecommendationReason: recommendationReason,
	}
	c.JSON(http.StatusOK, response)
}
