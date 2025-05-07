package db

import (
	"github.com/elastic/go-elasticsearch/v8"
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

type ElkDB interface {
	Client() *elasticsearch.Client
}
