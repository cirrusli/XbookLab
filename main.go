package main

import (
	"log"

	"canigraduate/models"

	"github.com/go-redis/redis/v8"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// 初始化MySQL连接
func initMySQL() *gorm.DB {
	// 先连接无库名实例创建数据库
	createDB, _ := gorm.Open(mysql.Open("root:password@tcp(localhost:3306)/?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
	createDB.Exec("CREATE DATABASE IF NOT EXISTS book_community")

	// 连接目标数据库
	dsn := "root:password@tcp(localhost:3306)/book_community?charset=utf8mb4&parseTime=True&loc=Local"
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
	_ = initRedis()

	models.SetDB(db)

	db.AutoMigrate(&models.Book{})

	r := SetupRouter(db)
	if err := r.Run(":8000"); err != nil {
		log.Fatal("服务启动失败:", err)
	}
}
