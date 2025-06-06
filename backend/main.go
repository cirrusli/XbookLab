package main

import (
	"context"
	"fmt"
	"log"

	"xbooklab/handlers"
	"xbooklab/middleware"
	"xbooklab/models"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	db := initMySQL()
	rdb, err := initRedis()
	if err != nil {
		log.Println("Redis not ready",err)
	}
	models.SetDB(db, rdb)

	r := InitRouter()
	if err := r.Run(":8000"); err != nil {
		log.Fatal("服务启动失败:", err)
	}
}

func InitRouter() *gin.Engine {
	r := gin.Default()
	r.Static("/static", "./assets")
	r.Use(middleware.CorsMiddleware())

	// 认证相关路由
	auth := r.Group("/api/auth")
	{
		auth.POST("/login", handlers.Login)
		auth.POST("/logout", middleware.AuthMiddleware(), handlers.Logout)
		auth.POST("/change-password", middleware.AuthMiddleware(), handlers.ChangePassword)
	}

	// 用户相关路由
	user := r.Group("/api/user")
	user.POST("/register", handlers.Register)
	user.Use(middleware.AuthMiddleware())
	{
		// 现在只能获取自己的信息
		user.POST("/profile", handlers.GetUserProfile)
		user.POST("/update", handlers.UpdateUserProfile)
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
		book.POST("/upload", handlers.UploadBookCover)
		book.GET("/:id", handlers.GetBookDetail)
		book.DELETE("/:id", handlers.DeleteBook)
		book.GET("/comments", handlers.GetComments)
	}

	// 推荐书籍路由
	r.GET("/api/recommend", middleware.AuthMiddleware(), handlers.GetRecommendedBooks)
	r.GET("/api/getGroups", handlers.GetGroups)
	r.GET("/api/getFriendsBook", handlers.GetFriendsBooks)

	// 用户行为路由
	record := r.Group("/api/record")
	record.Use(middleware.AuthMiddleware())
	{
		record.POST("/view", handlers.RecordBookView)
		record.POST("/rating", handlers.RecordBookRating)
	}

	// 话题相关路由
	topic := r.Group("/api/topic")
	topic.Use(middleware.AuthMiddleware())
	{
		topic.GET("/", handlers.GetTopics)
		topic.POST("/", handlers.CreateTopic)
		topic.GET("/:id", handlers.GetTopicDetail)
		topic.PUT("/:id", handlers.UpdateTopic)
		topic.DELETE("/:id", handlers.DeleteTopic)
	}

	// 评论相关路由
	comment := r.Group("/api/comment")
	comment.Use(middleware.AuthMiddleware())
	{
		comment.POST("/", handlers.CreateComment)
		comment.GET("/:targetType/:targetId", handlers.GetComments)
		comment.DELETE("/:id", handlers.DeleteComment)
	}

	// 点赞相关路由
	like := r.Group("/api/like")
	like.Use(middleware.AuthMiddleware())
	{
		like.POST("/topic", handlers.LikeTopic)
	}

	// 用户动态路由--not implemented
	event := r.Group("/api/event")
	event.Use(middleware.AuthMiddleware())
	{
		event.GET("/", handlers.GetUserEvents)
	}

	// 标签管理路由
	tag := r.Group("/api/tag")
	tag.Use(middleware.AuthMiddleware())
	{
		tag.GET("/", handlers.GetTagList)
		tag.POST("/", handlers.AddTag)
		tag.POST("/delete", handlers.DeleteTag)
	}

	// 用户管理路由--not implemented
	manager := r.Group("/api/manager")
	manager.Use(middleware.AuthMiddleware())
	{
		manager.GET("/users", handlers.GetUserList)
		manager.DELETE("/users/:id", handlers.DeleteUser)
	}

	return r
}

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
func initRedis() (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	// 检查连接是否正常
	if err := client.Ping(context.Background()).Err(); err != nil {
		return nil, fmt.Errorf("Redis连接失败: %v", err)
	}
	return client, nil
}
