package handlers

import (
	"log"
	"net/http"
	"xbooklab/models"

	"github.com/gin-gonic/gin"
)

type RecordBookViewRequest struct {
	// UserID uint `json:"user_id"`
	BookID uint `json:"book_id"`
}

type RecordBookViewResponse struct {
	Code    uint   `json:"code"`
	Data    string `json:"data"`
	Message string `json:"message"`
}

type RecordBookRatingRequest struct {
	BookID uint    `json:"book_id"`
	Rating float32 `json:"rating"`
}

type RecordBookRatingResponse struct {
	Code    uint   `json:"code"`
	Message string `json:"message"`
}

type RecordBookCommentRequest struct {
	UserID  uint   `json:"user_id"`
	BookID  uint   `json:"book_id"`
	Content string `json:"content"`
}
type RecordBookCommentResponse struct {
	Code    uint   `json:"code"`
	Message string `json:"message"`
}

// RecordBookView 记录用户浏览行为
func RecordBookView(c *gin.Context) {
	var req RecordBookViewRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, RecordBookViewResponse{
			Code:    400,
			Message: "invalid request",
		})
		return
	}
	userID := c.GetUint("userID")
	if userID == 0 {
		userID = 1
	}

	if err := models.CreateBookView(userID, uint(req.BookID)); err != nil {
		c.JSON(http.StatusInternalServerError, RecordBookViewResponse{
			Code:    500,
			Message: "failed to record view",
		})
		return
	}

	// 更新推荐评分（浏览+0.1）
	if err := models.UpsertRecommendScore(userID, uint(req.BookID), 0.1); err != nil {
		log.Printf("更新推荐评分失败: %v", err)
	}

	c.JSON(http.StatusOK, RecordBookViewResponse{
		Code:    200,
		Data:    "",
		Message: "view recorded",
	})
}

// RecordBookRating 记录用户评分行为
func RecordBookRating(c *gin.Context) {
	var req RecordBookRatingRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, RecordBookRatingResponse{
			Code:    400,
			Message: "invalid request",
		})
		return
	}
	userID := c.GetUint("userID")
	if userID == 0 {
		userID = 1
	}
	log.Println(req.Rating)
	req.Rating = req.Rating * 2 // 前端传过来的是0-5，这里转换为0-10

	if err := models.CreateOrUpdateRating(userID, uint(req.BookID), uint(req.Rating)); err != nil {
		c.JSON(http.StatusInternalServerError, RecordBookRatingResponse{
			Code:    500,
			Message: "failed to record rating",
		})
		return
	}

	// 更新推荐评分（评分+0.5）
	if err := models.UpsertRecommendScore(userID, uint(req.BookID), 0.5); err != nil {
		log.Printf("更新推荐评分失败: %v", err)
	}

	c.JSON(http.StatusOK, RecordBookRatingResponse{
		Code:    200,
		Message: "rating recorded",
	})
}

// RecordBookComment 记录用户评论行为，似乎有评论表就好，不需要这个接口
// func RecordBookComment(c *gin.Context) {
// 	var req RecordBookCommentRequest
// 	if err := c.ShouldBindJSON(&req); err != nil {
// 		c.JSON(http.StatusBadRequest, RecordBookCommentResponse{
// 			Code:    400,
// 			Message: "invalid request",
// 		})
// 		return
// 	}

// 	if len(req.Content) == 0 {
// 		c.JSON(http.StatusBadRequest, RecordBookCommentResponse{
// 			Code:    400,
// 			Message: "comment content cannot be empty",
// 		})
// 		return
// 	}

// 	if err := models.CreateComment(db, req.UserID, req.BookID, req.Content); err != nil {
// 		c.JSON(http.StatusInternalServerError, RecordBookCommentResponse{
// 			Code:    500,
// 			Message: "failed to record comment",
// 		})
// 		return
// 	}

// 	c.JSON(http.StatusOK, RecordBookCommentResponse{
// 		Code:    200,
// 		Message: "comment recorded",
// 	})
// }
