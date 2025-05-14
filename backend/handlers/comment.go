package handlers

import (
	"log"
	"net/http"
	"strconv"

	"xbooklab/models"

	"github.com/gin-gonic/gin"
)

type CreateCommentRequest struct {
	Content  string `json:"content"`
	TargetID uint   `json:"comment_id"`
	Type     uint   `json:"comment_type"`
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
	var req CreateCommentRequest
	var comment models.Comment
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	comment.UserID = c.GetUint("userID")
	log.Println(comment.UserID)
	if comment.UserID == 0 {
		comment.UserID = 1
	}
	comment.Content = req.Content
	comment.TargetID = req.TargetID
	comment.Type = uint8(req.Type)
	if err := models.CreateComment(&comment); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "创建评论失败"})
		return
	}

	// 更新推荐评分（评论+0.5，targetID为bookID）
	if err := models.UpsertRecommendScore(comment.UserID, comment.TargetID, 0.5); err != nil {
		log.Printf("更新推荐评分失败: %v", err)
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"data":    "",
		"message": "all in one",
	})
}

// GetComments 获取评论列表
func GetComments(c *gin.Context) {
	targetType, _ := strconv.Atoi(c.Param("targetType"))
	targetID, _ := strconv.Atoi(c.Param("targetId"))
	if targetType == 0 || targetID == 0 {
		var comments []models.Comment
		err := models.DB.Preload("User").Order("created_at DESC").Find(&comments).Error
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "获取评论失败"})
			return
		}
		for i := range comments {
			if comments[i].User.Username == "" {
				comments[i].User.Username = "匿名用户"
			}
			comments[i].Author = comments[i].User.Username
			comments[i].Time = comments[i].CreatedAt.Local().Format("2006-01-02 15:04:05")

		}
		c.JSON(http.StatusOK, gin.H{
			"code":    200,
			"data":    comments,
			"message": "all in one",
		})
		return
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
