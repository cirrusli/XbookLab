package main

import (
	"log"

	"xbooklab/handlers"
	"xbooklab/middleware"
	"xbooklab/models"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// 初始化MySQL连接
func initMySQL() *gorm.DB {
	// 先连接无库名实例创建数据库
	createDB, _ := gorm.Open(mysql.Open("root:lzq@tcp(localhost:3306)/?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
	createDB.Exec("CREATE DATABASE IF NOT EXISTS x_book_lab")

	// 连接目标数据库
	dsn := "root:lzq@tcp(localhost:3306)/x_book_lab_test?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("MySQL连接失败:", err)
	}
	return db
}

// 初始化Redis连接
func initRedis() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
}

func main() {
	db := initMySQL()
	models.SetDB(db)

	rdb := initRedis()

	r := InitRouter(db, rdb)
	if err := r.Run(":8000"); err != nil {
		log.Fatal("服务启动失败:", err)
	}
}
func InitRouter(db *gorm.DB, rdb *redis.Client) *gin.Engine {
	r := gin.Default()
	r.Use(gin.Logger(), gin.Recovery())

	// 认证相关路由
	auth := r.Group("/api/auth")
	{
		auth.POST("/login", handlers.Login)
		auth.POST("/logout", handlers.Logout)
		auth.PUT("/change-password", middleware.AuthMiddleware(), handlers.ChangePassword)
	}

	// 用户相关路由
	user := r.Group("/api/user")
	user.Use(middleware.AuthMiddleware())
	{
		user.POST("/register", handlers.Register)
		user.GET("/profile", handlers.GetUserProfile)
		user.PUT("/profile", handlers.UpdateUserProfile)
		user.POST("/follow/:id", handlers.FollowUser)
		user.DELETE("/follow/:id", handlers.UnfollowUser)
		user.GET("/following", handlers.GetFollowing)
		user.GET("/followers", handlers.GetFollowers)
	}

	// 书籍相关路由
	book := r.Group("/api/book")
	{
		book.GET("/", handlers.GetBookList)
		book.POST("/", handlers.CreateBook)
		book.GET("/:id", handlers.GetBookDetail)
		book.DELETE("/:id", handlers.DeleteBook)
	}

	// 推荐书籍路由
	r.GET("/api/recommend", handlers.GetRecommendedBooks)

	// 用户行为路由
	record := r.Group("/api/record")
	record.Use(middleware.AuthMiddleware())
	{
		record.POST("/view/:bookId", handlers.RecordBookView)
		record.POST("/rate/:bookId", handlers.RecordBookRating)
	}

	// 话题相关路由
	topics := r.Group("/api/topics")
	{
		topics.GET("/", handlers.GetTopics)
		topics.POST("/", handlers.CreateTopic)
		topics.GET("/:id", handlers.GetTopic)
		topics.PUT("/:id", handlers.UpdateTopic)
		topics.DELETE("/:id", handlers.DeleteTopic)
	}

	// 评论相关路由
	comments := r.Group("/api/comments")
	comments.Use(middleware.AuthMiddleware())
	{
		comments.POST("/", handlers.CreateComment)
		comments.GET("/:targetType/:targetId", handlers.GetComments)
		comments.DELETE("/:id", handlers.DeleteComment)
	}

	// 点赞相关路由
	likes := r.Group("/api/likes")
	likes.Use(middleware.AuthMiddleware())
	{
		likes.POST("/topic/:id", handlers.LikeTopic)
		likes.DELETE("/topic/:id", handlers.UnlikeTopic)
	}

	// 用户动态路由
	events := r.Group("/api/events")
	events.Use(middleware.AuthMiddleware())
	{
		events.GET("/", handlers.GetUserEvents)
	}

	return r
}
