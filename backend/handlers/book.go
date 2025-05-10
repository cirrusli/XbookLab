package handlers
import (
	"net/http"
	"xbooklab/models"

	"github.com/gin-gonic/gin"
)

type CreateBookRequest struct {
	BookID      uint   `json:"book_id"`
	Title       string `json:"title"`
	Author      string `json:"author"`
	Cover       string `json:"cover"`
	Description string `json:"description"`
	Rating      string `json:"rating"`
	TagName     string `json:"tag_name"`
}

type CreateBookResponse struct {
	Code    uint   `json:"code"`
	Message string `json:"message"`
}

type GetBookDetailRequest struct {
	BookID uint `json:"book_id"`
}

type GetBookDetailResponse struct {
	BookID      uint   `json:"book_id"`
	Title       string `json:"title"`
	Author      string `json:"author"`
	Cover       string `json:"cover"`
	Description string `json:"description"`
	Rating      string `json:"rating"`
	TagName     string `json:"tag_name"`
}

type SetBookRatingRequest struct {
	BookID uint `json:"book_id"`
	Rating uint `json:"rating"`
	UserID uint `json:"user_id"`
}

type SetBookRatingResponse struct {
	Message string `json:"message"`
}

type GetBookListRequest struct {
	Limit  uint `json:"limit"`
	Offset uint `json:"offset"`
}

type GetBookListResponse struct {
	Books    []models.Book `json:"book_list"`
	TotalNum uint   `json:"total_num"`
}

type DeleteBookRequest struct {
	BookID uint `json:"book_id"`
}

type DeleteBookResponse struct {
	Code    uint   `json:"code"`
	Message string `json:"message"`
}


func GetBookList(c *gin.Context) {
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
	models.DB.Create(&book)
	c.JSON(http.StatusCreated, book)
}

func GetBookDetail(c *gin.Context) {
	id := c.Param("id")
	var book models.Book
	if err := models.DB.First(&book, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}
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

// func GetRecommendedBooks(c *gin.Context) {
// 	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
// 	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "10"))
// 	categoryID := c.Query("categoryId")

// 	// 从上下文中获取用户ID
// 	userID, _ := c.Get("userID")
// 	var uid uint
// 	if userID != nil {
// 		uid = userID.(uint)
// 	}

// 	books, total := models.GetRecommendedBooks(page, pageSize, categoryID, uid)
// 	c.JSON(http.StatusOK, gin.H{
// 		"data":  books,
// 		"total": total,
// 	})
// }
