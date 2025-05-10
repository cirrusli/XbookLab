package handlers
import (
	"net/http"
	"strconv"

	"xbooklab/models"

	"github.com/gin-gonic/gin"
)
type CreateCommentRequest struct {
	Comment models.Comment `json:"comment"`
}

type CreateCommentResponse struct {
	Code    uint   `json:"code"`
	Message string `json:"message"`
}

type GetCommentListRequest struct {
	CommentID   uint `json:"comment_id"`
	CommentType uint `json:"comment_type"`
}

type GetCommentListResponse struct {
	Comments []models.Comment `json:"comments"`
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

	userID := c.GetUint("userID")
	comment.UserID = userID

	if err := models.DB.Create(&comment).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "创建评论失败"})
		return
	}

	// 记录用户动态
	event := models.Event{
		UserID:       userID,
		EventType:    1, // 评论
		EventContent: "发表了评论",
		TargetID:     comment.TargetID,
		TargetType:   comment.Type,
	}
	models.DB.Create(&event)

	c.JSON(http.StatusOK, comment)
}

// GetComments 获取评论列表
func GetComments(c *gin.Context) {
	targetType, _ := strconv.Atoi(c.Param("targetType"))
	targetID, _ := strconv.Atoi(c.Param("targetId"))

	var comments []models.Comment
	if err := models.DB.Where("target_id = ? AND type = ?", targetID, targetType).
		Preload("User").Order("created_at DESC").Find(&comments).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "获取评论失败"})
		return
	}

	c.JSON(http.StatusOK, comments)
}

// DeleteComment 删除评论
func DeleteComment(c *gin.Context) {
	commentID, _ := strconv.Atoi(c.Param("id"))
	userID := c.GetUint("userID")

	var comment models.Comment
	if err := models.DB.Where("comment_id = ? AND user_id = ?", commentID, userID).First(&comment).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "评论不存在或无权删除"})
		return
	}

	if err := models.DB.Delete(&comment).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "删除评论失败"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "评论删除成功"})
}
