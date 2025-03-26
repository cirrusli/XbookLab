package main

import (
	"canigraduate/handlers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()
	r.Use(gin.Logger(), gin.Recovery())

	auth := r.Group("/api/auth")
	{
		auth.POST("/register", handlers.Register)
		auth.POST("/login", handlers.Login)
	}

	books := r.Group("/api/books")
	{
		books.GET("/", handlers.GetBooks)
		books.POST("/", handlers.CreateBook)
		books.GET("/:id", handlers.GetBook)
		books.PUT("/:id", handlers.UpdateBook)
		books.DELETE("/:id", handlers.DeleteBook)
	}

	r.GET("/api/recommend", handlers.GetRecommendedBooks)

	return r
}
