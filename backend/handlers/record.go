package handlers
import (
	"net/http"

	"github.com/gin-gonic/gin"
)
type RecordBookViewRequest struct {
	UserID uint `json:"user_id"`
	BookID uint `json:"book_id"`
}

type RecordBookViewResponse struct {
	Code    uint   `json:"code"`
	Message string `json:"message"`
}

type RecordBookRatingRequest struct {
	UserID uint `json:"user_id"`
	BookID uint `json:"book_id"`
	Rating uint `json:"rating"`
}

type RecordBookRatingResponse struct {
	Code    uint   `json:"code"`
	Message string `json:"message"`
}

type RecordBookCommentRequest struct {
	UserID    uint   `json:"user_id"`
	BookID    uint   `json:"book_id"`
	Content   string `json:"content"`
}
type RecordBookCommentResponse struct {
	Code    uint   `json:"code"`
	Message string `json:"message"`
}


// RecordBookView 记录用户浏览行为
func RecordBookView(c *gin.Context) {
	// 1. 获取用户ID和书籍ID
	// userID := c.MustGet("userID").(uint32)
	// bookID, err := strconv.ParseUint(c.Param("bookId"), 10, 32)
	// if err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": "invalid book id"})
	// 	return
	// }

	// // 2. 获取数据库连接
	// db := c.MustGet("db").(*sql.DB)

	// // 3. 记录浏览行为
	// err = models.RecordBookView(context.Background(), db, userID, uint32(bookID))
	// if err != nil {
	// 	c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	// 	return
	// }

	// 4. 返回成功响应
	c.JSON(http.StatusOK, gin.H{
		"message": "view recorded",
	})
}

// RecordBookRating 记录用户评分行为
func RecordBookRating(c *gin.Context) {
	// 1. 获取用户ID和书籍ID
	// userID := c.MustGet("userID").(uint32)
	// bookID, err := strconv.ParseUint(c.Param("bookId"), 10, 32)
	// if err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": "invalid book id"})
	// 	return
	// }

	// // 2. 解析请求体
	// var req struct {
	// 	Rating uint32 `json:"rating"`
	// }
	// if err := c.ShouldBindJSON(&req); err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// 	return
	// }

	// // 3. 获取数据库连接
	// db := c.MustGet("db").(*sql.DB)

	// // 4. 记录评分行为
	// err = models.RecordBookRating(context.Background(), db, userID, uint32(bookID), req.Rating)
	// if err != nil {
	// 	c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	// 	return
	// }

	// 5. 返回成功响应
	c.JSON(http.StatusOK, gin.H{
		"message": "rating recorded",
	})
}

// RecordBookComment 记录用户评论行为
func RecordBookComment(c *gin.Context) {
	// // 1. 获取用户ID
	// userID := c.MustGet("userID").(uint32)

	// // 2. 解析请求体
	// var req models.RecordCommentReq
	// if err := c.ShouldBindJSON(&req); err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// 	return
	// }
	// req.UserID = userID

	// // 3. 获取数据库连接
	// db := c.MustGet("db").(*sql.DB)

	// // 4. 记录评论行为
	// err := models.RecordBookComment(context.Background(), db, req)
	// if err != nil {
	// 	c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	// 	return
	// }

	// 5. 返回成功响应
	c.JSON(http.StatusOK, gin.H{
		"message": "comment recorded",
	})
}
