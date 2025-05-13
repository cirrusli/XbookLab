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
func GetRecommendedBooks(c *gin.Context) {
	// 解析URL查询参数
	var req GetRecommendedBooksRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	log.Printf("Received query params: %+v", req)
	var recommendedBooks []models.Book

	var err error
	userID := c.GetUint("userID")
	log.Println(userID)
	if userID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "登录已过期，请重新登录"})
		return
	}
	
	log.Println("userID:", userID, "limit:", req.Limit, "offset:", req.Offset, "tagFilter:", req.TagFilter)
	recommendedBooks, err = models.GetRecommendedBooks(userID, int(req.Limit), int(req.Offset), int(req.TagFilter))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	log.Println("recommendedBooks:", recommendedBooks)
	log.Println("recommendedBooks length:", len(recommendedBooks))

	if len(recommendedBooks) < int(req.Limit) {
		// 用事件推荐补充
		remaining := int(req.Limit) - len(recommendedBooks)
		newRecommendedBooks := recommendByEvent()
		if len(newRecommendedBooks) >= remaining {
			recommendedBooks = append(recommendedBooks, newRecommendedBooks[:remaining]...)
		} else {
			// 如果推荐书籍数量不足，用热门书籍补充
			popularBooks := models.GetPopularBooks(remaining, int(req.TagFilter))
			recommendedBooks = append(recommendedBooks, popularBooks...)
		}

	}

	response := gin.H{
		"code":    200,
		"message": "success",
		"data": gin.H{
			"Books": recommendedBooks,
		},
	}
	c.JSON(http.StatusOK, response)
}

// 获取用户最近的评分、关注、浏览、评论行为，综合插入新的推荐数据
func recommendByEvent() []models.Book {
	return nil
}
