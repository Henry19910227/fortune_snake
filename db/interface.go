package db

import (
	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
)

type MysqlDB interface {
	DB() *gorm.DB
	Close()
}

type RedisDB interface {
	RDB() *redis.Client
	Close()
}

type MongoDB interface {
	Client() *mongo.Client
	Close()
}
