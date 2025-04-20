package handlers

import (
	"net/http"
	"xbooklab/models"

	"github.com/gin-gonic/gin"
)

func RecordBookView(c *gin.Context) {
	userID := c.MustGet("userID").(uint)
	bookID := c.Param("bookId")

	var interaction models.UserInteraction
	models.DB.Where("user_id = ? AND book_id = ?", userID, bookID).FirstOrCreate(&interaction)

	interaction.ViewCount++
	models.DB.Save(&interaction)

	c.JSON(http.StatusOK, gin.H{"message": "记录成功"})
}

func RecordBookRating(c *gin.Context) {
	userID := c.MustGet("userID").(uint)
	bookID := c.Param("bookId")

	var req struct {
		Rating float64 `json:"rating" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var interaction models.UserInteraction
	models.DB.Where("user_id = ? AND book_id = ?", userID, bookID).FirstOrCreate(&interaction)

	interaction.Rating = req.Rating
	models.DB.Save(&interaction)

	c.JSON(http.StatusOK, gin.H{"message": "评分成功"})
}
