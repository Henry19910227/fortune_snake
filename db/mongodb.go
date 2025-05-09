package db

import (
	"context"
	"fmt"
	model "game_server_slots_fortune_snake/internal/model/config/system"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

type mongoDB struct {
	client *mongo.Client
}

func NewMongoDB(config model.MongodbConfig) (MongoDB, error) {
	uri := fmt.Sprintf("mongodb://%s:%s@%s/%s?authSource=%s",
		config.Username,
		config.Password,
		config.Hosts[0],
		config.Database,
		config.AuthSource,
	)
	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		return nil, err
	}
	log.Println("✅ MongoDB 连接成功")
	return &mongoDB{client: client}, nil
}

func (m *mongoDB) Client() *mongo.Client {
	return m.client
}

func (m *mongoDB) Close() {
	if m.client == nil {
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err := m.client.Disconnect(ctx)
	if err != nil {
		return
	}
	log.Println("🛑 MongoDB 连接已关闭")
}
