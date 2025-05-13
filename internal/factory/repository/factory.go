package repository

import (
	gameRepo "game_server_slots_fortune_snake/internal/repository/game"
	playerRepo "game_server_slots_fortune_snake/internal/repository/player"
	weightRepo "game_server_slots_fortune_snake/internal/repository/weight"
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

func (f *factory) GameRepository() gameRepo.Repository {
	return gameRepo.New()
}

func (f *factory) PlayerRepository() playerRepo.Repository {
	return playerRepo.New(f.rdb)
}

func (f *factory) WeightRepository() weightRepo.Repository {
	return weightRepo.New()
}
