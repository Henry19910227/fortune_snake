package repository

import (
	"game_server_slots_fortune_snake/internal/repository/player"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type factory struct {
	db  *gorm.DB
	rdb *redis.Client
}

func New(db *gorm.DB, rdb *redis.Client) Factory {
	repoFactory := &factory{db: db, rdb: rdb}
	return repoFactory
}

func (f *factory) PlayerRepository() player.Repository {
	return player.New(f.rdb)
}
