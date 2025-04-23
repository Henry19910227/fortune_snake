package repository

import (
	playerRepo "game_server_slots_fortune_snake/repository/player"
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

func (f *factory) PlayerRepository() playerRepo.Repository {
	return playerRepo.New(f.rdb)
}
