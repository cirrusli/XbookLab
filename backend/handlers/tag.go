package handlers

import (
	"net/http"
	"xbooklab/models"

	"github.com/gin-gonic/gin"
)

func GetTagList(c *gin.Context) {
	tags, err := models.GetAllTags()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "获取标签列表失败"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"tags": tags})
}

func AddTag(c *gin.Context) {
	var tag models.Tag
	if err := c.ShouldBindJSON(&tag); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := models.CreateTag(&tag); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "标签添加失败"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "标签添加成功"})
}

func DeleteTag(c *gin.Context) {
	tagID := c.Param("id")
	if err := models.DeleteTag(tagID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "标签删除失败"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "标签删除成功"})
}
