package db

import (
	"context"
	"game_server_slots_fortune_snake/model"
	"github.com/redis/go-redis/v9"
	"log"
)

type redisDB struct {
	rdb *redis.Client
}

func NewRedisDB(config model.RedisConfig) (RedisDB, error) {
	if config.MasterName == "" || len(config.SentinelAddress) == 0 {
		log.Fatal("❌ 配置错误: 必须提供 `redis.master_name` 和 `redis.sentinel_address`")
		return nil, nil
	}

	// 连接 Redis 哨兵模式，并提供密码
	rdb := redis.NewFailoverClient(&redis.FailoverOptions{
		MasterName:       config.MasterName,
		SentinelAddrs:    config.SentinelAddress,
		SentinelPassword: config.SentinelPassword, // Sentinel 认证密码
		Password:         config.Password,         // Redis 主库认证密码
		DB:               config.Db,
	})

	// 测试连接
	ctx := context.Background()
	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		log.Fatalf("❌ Redis 连接失败: %v", err)
		return nil, err
	}
	log.Println("✅ Redis Sentinel 连接成功")
	return &redisDB{rdb: rdb}, nil
}

func (r *redisDB) RDB() *redis.Client {
	return r.rdb
}

func (r *redisDB) Close() {
	if r.rdb == nil {
		return
	}
	if err := r.rdb.Close(); err != nil {
		log.Printf("❌ 关闭 Redis 失败: %v", err)
		return
	}
	log.Println("🛑 Redis 连接已关闭")
}
