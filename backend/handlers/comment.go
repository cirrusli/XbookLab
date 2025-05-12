package handlers

import (
	"net/http"
	"strconv"

	"xbooklab/models"

	"github.com/gin-gonic/gin"
)

type CreateCommentRequest struct {
	Content  string `json:"content"`
	TargetID uint   `json:"target_id"`
	Type     uint   `json:"type"`
}

type CreateCommentResponse struct {
	Code    uint   `json:"code"`
	Message string `json:"message"`
}

type GetCommentListRequest struct {
	CommentID   uint `json:"comment_id"`
	CommentType uint `json:"comment_type"`
}

type RecordCommentRequest struct {
	UserID    uint `json:"user_id"`
	CommentID uint `json:"comment_id"`
}

type RecordCommentResponse struct {
	Code    uint   `json:"code"`
	Message string `json:"message"`
}

// CreateComment 创建评论
func CreateComment(c *gin.Context) {
	var comment models.Comment
	if err := c.ShouldBindJSON(&comment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	comment.UserID = c.GetUint("userID")
	if err := models.CreateComment(&comment); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "创建评论失败"})
		return
	}

	c.JSON(http.StatusOK, comment)
}

// GetComments 获取评论列表
func GetComments(c *gin.Context) {
	targetType, _ := strconv.Atoi(c.Param("targetType"))
	targetID, _ := strconv.Atoi(c.Param("targetId"))
	if targetType == 0 || targetID == 0 {
		var comments []models.Comment
		err := models.DB.Preload("User").Find(&comments).Error
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "获取评论失败"})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"data": gin.H{
				// 所有的详情落地页都用同一个评论区
				"comments": comments,
			},
			"message": "all in one",
		})
	}
	comments, err := models.GetComments(uint(targetID), uint(targetType))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "获取评论失败"})
		return
	}

	c.JSON(http.StatusOK, comments)
}

// DeleteComment 删除评论
func DeleteComment(c *gin.Context) {
	commentID, _ := strconv.Atoi(c.Param("id"))

	if err := models.DeleteComment(uint(commentID)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "删除评论失败"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "评论删除成功"})
}
