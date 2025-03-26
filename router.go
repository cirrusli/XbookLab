package main

import (
	"canigraduate/handlers"
	"canigraduate/middleware"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()
	r.Use(gin.Logger(), gin.Recovery())

	// 认证相关路由
	auth := r.Group("/api/auth")
	{
		auth.POST("/register", handlers.Register)
		auth.POST("/login", handlers.Login)
		auth.POST("/logout", handlers.Logout)
		auth.PUT("/change-password", middleware.AuthMiddleware(), handlers.ChangePassword)
	}

	// 用户相关路由 (需要认证)
	user := r.Group("/api/user")
	user.Use(middleware.AuthMiddleware())
	{
		user.GET("/profile", handlers.GetUserProfile)
		user.PUT("/profile", handlers.UpdateUserProfile)
		user.POST("/avatar", handlers.UploadAvatar)
	}

	// 书籍相关路由
	books := r.Group("/api/books")
	{
		books.GET("/", handlers.GetBooks)
		books.POST("/", handlers.CreateBook)
		books.GET("/:id", handlers.GetBook)
		books.PUT("/:id", handlers.UpdateBook)
		books.DELETE("/:id", handlers.DeleteBook)
	}

	// 推荐书籍路由
	r.GET("/api/recommend", handlers.GetRecommendedBooks)

	// 用户行为路由
	interaction := r.Group("/api/interaction")
	interaction.Use(middleware.AuthMiddleware())
	{
		interaction.POST("/view/:bookId", handlers.RecordBookView)
		interaction.POST("/rate/:bookId", handlers.RecordBookRating)
	}

	return r
}
