package handlers

import (
	"net/http"

	"xbooklab/models"

	"github.com/gin-gonic/gin"
)

// GetUserEvents 获取用户动态
func GetUserEvents(c *gin.Context) {
	userID := c.GetUint("userID")

	events, err := models.GetEventsByUserID(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "获取用户动态失败"})
		return
	}

	c.JSON(http.StatusOK, events)
}
