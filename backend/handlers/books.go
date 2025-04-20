package handlers

import (
	"net/http"
	"strconv"
	"xbooklab/models"

	"github.com/gin-gonic/gin"
)

func GetBooks(c *gin.Context) {
	var books []models.Book
	models.DB.Find(&books)
	c.JSON(http.StatusOK, books)
}

func CreateBook(c *gin.Context) {
	var book models.Book
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if len(book.CategoryTags) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "必须选择至少一个分类标签"})
		return
	}
	models.DB.Create(&book)
	c.JSON(http.StatusCreated, book)
}

func GetBook(c *gin.Context) {
	id := c.Param("id")
	var book models.Book
	if err := models.DB.First(&book, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}
	c.JSON(http.StatusOK, book)
}

func UpdateBook(c *gin.Context) {
	id := c.Param("id")
	var book models.Book
	if err := models.DB.First(&book, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if len(book.CategoryTags) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "必须选择至少一个分类标签"})
		return
	}
	models.DB.Save(&book)
	c.JSON(http.StatusOK, book)
}

func DeleteBook(c *gin.Context) {
	id := c.Param("id")
	var book models.Book
	if err := models.DB.First(&book, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}
	models.DB.Delete(&book)
	c.JSON(http.StatusOK, gin.H{"message": "Book deleted successfully"})
}

func GetRecommendedBooks(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "10"))
	categoryID := c.Query("categoryId")

	// 从上下文中获取用户ID
	userID, _ := c.Get("userID")
	var uid uint
	if userID != nil {
		uid = userID.(uint)
	}

	books, total := models.GetRecommendedBooks(page, pageSize, categoryID, uid)
	c.JSON(http.StatusOK, gin.H{
		"data":  books,
		"total": total,
	})
}
