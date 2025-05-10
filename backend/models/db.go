package models

import (
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

var DB *gorm.DB
var RDB *redis.Client

func SetDB(db *gorm.DB, rdb *redis.Client) {
	DB = db
	RDB = rdb
}
