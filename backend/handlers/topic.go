package handlers

import (
	"net/http"
	"strconv"

	"xbooklab/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// 创建话题
func CreateTopic(c *gin.Context) {
	var topic models.Topic
	if err := c.ShouldBindJSON(&topic); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if len(topic.CategoryTags) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "必须选择至少一个分类标签"})
		return
	}

	db := models.GetDB()
	if err := db.Create(&topic).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "创建话题失败"})
		return
	}

	c.JSON(http.StatusCreated, topic)
}

// 获取话题列表
func GetTopics(c *gin.Context) {
	var topics []models.Topic
	db := models.GetDB()

	if err := db.Find(&topics).Error; err != nil {
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
	db := models.GetDB()

	if err := db.First(&topic, id).Error; err != nil {
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
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的话题ID"})
		return
	}

	var topic models.Topic
	if err := c.ShouldBindJSON(&topic); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if len(topic.CategoryTags) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "必须选择至少一个分类标签"})
		return
	}

	if len(topic.CategoryTags) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "必须选择至少一个分类标签"})
		return
	}

	topic.ID = uint(id)
	db := models.GetDB()

	if err := db.Save(&topic).Error; err != nil {
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

	db := models.GetDB()

	if err := db.Delete(&models.Topic{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "删除话题失败"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "话题删除成功"})
}
