package handlers

import (
	"net/http"

	"xbooklab/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// LikeTopic 点赞/取消点赞话题
func LikeTopic(c *gin.Context) {
	var req struct {
		IsLiked int `json:"is_liked"`
		TopicID int `json:"topic_id"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "参数错误"})
		return
	}

	userID := c.GetUint("userID")
	topicID := req.TopicID

	if req.IsLiked == 1 {
		// 检查是否已点赞
		var existingLike models.TopicLike
		if err := models.DB.Where("user_id = ? AND topic_id = ?", userID, topicID).First(&existingLike).Error; err == nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"code":  400,
				"error": "已点赞过该话题"})
			return
		}

		// 创建点赞记录
		like := models.TopicLike{
			UserID:  userID,
			TopicID: uint(topicID),
		}
		if err := models.DB.Create(&like).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"code":  500,
				"error": "点赞失败"})
			return
		}

		// 更新话题点赞数
		models.DB.Model(&models.Topic{}).Where("topic_id = ?", topicID).Update("like_count", gorm.Expr("like_count + 1"))

		// 记录用户动态
		event := models.Event{
			UserID:       userID,
			EventType:    3, // 点赞话题
			EventContent: "点赞了话题",
			TargetID:     uint(topicID),
			TargetType:   1, // 话题
		}
		models.DB.Create(&event)

		c.JSON(http.StatusOK, gin.H{
			"code":    200,
			"message": "点赞成功"})
	} else {
		// 查找点赞记录
		var like models.TopicLike
		if err := models.DB.Where("user_id = ? AND topic_id = ?", userID, topicID).First(&like).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"code":  404,
				"error": "未点赞过该话题"})
			return
		}

		// 删除点赞记录
		if err := models.DB.Delete(&like).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"code":  500,
				"error": "取消点赞失败"})
			return
		}

		// 更新话题点赞数
		models.DB.Model(&models.Topic{}).Where("topic_id = ?", topicID).Update("like_count", gorm.Expr("like_count - 1"))

		c.JSON(http.StatusOK, gin.H{
			"code":    200,
			"message": "取消点赞成功"})
	}
}
