package handlers

import (
	"net/http"
	"strconv"

	"xbooklab/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type GetTopicListRequest struct {
	Offset    uint `json:"offset"`
	Limit     uint `json:"limit"`
	TagFilter uint `json:"tag_filter"`
}

type GetTopicListResponse struct {
	Topics []Topic `json:"topics"`
	Total  uint    `json:"total"`
}

type CreateTopicRequest struct {
	Title        string `json:"title"`
	Content      string `json:"content"`
	AuthorUserID uint   `json:"author_user_id"`
}

type CreateTopicResponse struct {
	TopicID uint   `json:"topic_id"`
	Code    uint   `json:"code"`
	Message string `json:"message"`
}

type GetTopicDetailRequest struct {
	TopicID uint `json:"topic_id"`
}

type GetTopicDetailResponse struct {
	Topic Topic `json:"topic"`
}

type LikeTopicRequest struct {
	UserID  uint `json:"user_id"`
	TopicID uint `json:"topic_id"`
}

type LikeTopicResponse struct {
	Code    uint   `json:"code"`
	Message string `json:"message"`
}

type DeleteTopicRequest struct {
	TopicID uint `json:"topic_id"`
}

type DeleteTopicResponse struct {
	Code    uint   `json:"code"`
	Message string `json:"message"`
}

type Topic struct {
	TopicID      uint   `json:"topic_id"`
	Title        string `json:"title"`
	Content      string `json:"content"`
	AuthorUserID uint   `json:"author_user_id"`
	LikeCount    uint   `json:"like_count"`
	TagName      string `json:"tag_name"`
}

// 创建话题
func CreateTopic(c *gin.Context) {
	var topic models.Topic
	if err := c.ShouldBindJSON(&topic); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// if len(topic.Tag) == 0 {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": "必须选择至少一个分类标签"})
	// 	return
	// }

	if err := models.DB.Create(&topic).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "创建话题失败"})
		return
	}

	c.JSON(http.StatusCreated, topic)
}

// 获取话题列表
func GetTopics(c *gin.Context) {
	var topics []models.Topic

	if err := models.DB.Find(&topics).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "获取话题列表失败"})
		return
	}

	c.JSON(http.StatusOK, topics)
}

// 获取单个话题
func GetTopic(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的话题ID"})
		return
	}

	var topic models.Topic

	if err := models.DB.First(&topic, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "话题不存在"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "获取话题失败"})
		}
		return
	}

	c.JSON(http.StatusOK, topic)
}

// 更新话题
func UpdateTopic(c *gin.Context) {
	_, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的话题ID"})
		return
	}

	var topic models.Topic
	if err := c.ShouldBindJSON(&topic); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// if len(topic.Tag) == 0 {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": "必须选择至少一个分类标签"})
	// 	return
	// }

	// if len(topic.Tag) == 0 {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": "必须选择至少一个分类标签"})
	// 	return
	// }

	// topic.ID = uint(id)

	if err := models.DB.Save(&topic).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "更新话题失败"})
		return
	}

	c.JSON(http.StatusOK, topic)
}

// 删除话题
func DeleteTopic(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的话题ID"})
		return
	}

	if err := models.DB.Delete(&models.Topic{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "删除话题失败"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "话题删除成功"})
}
