package repository

import (
	gameCfg "game_server_slots_fortune_snake/config/game"
	bucketRepo "game_server_slots_fortune_snake/internal/repository/bucket"
	bucketFreeRepo "game_server_slots_fortune_snake/internal/repository/bucket_free"
	gameRepo "game_server_slots_fortune_snake/internal/repository/game"
	playerRepo "game_server_slots_fortune_snake/internal/repository/player"
	resultRepo "game_server_slots_fortune_snake/internal/repository/result"
	weightRepo "game_server_slots_fortune_snake/internal/repository/weight"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type factory struct {
	db  *gorm.DB
	rdb *redis.Client
	cfg gameCfg.Config
}

func New(db *gorm.DB, rdb *redis.Client, cfg gameCfg.Config) Factory {
	repoFactory := &factory{db: db, rdb: rdb, cfg: cfg}
	return repoFactory
}

func (f *factory) GameRepository() gameRepo.Repository {
	return gameRepo.New()
}

func (f *factory) PlayerRepository() playerRepo.Repository {
	return playerRepo.New(f.rdb)
}

func (f *factory) WeightRepository() weightRepo.Repository {
	return weightRepo.New(f.cfg.RTPConfig())
}

func (f *factory) ResultRepository() resultRepo.Repository {
	return resultRepo.New(f.db)
}

func (f *factory) ResultFreeRepository() resultRepo.Repository {
	return resultRepo.NewFree(f.db)
}

func (f *factory) BucketRepository() bucketRepo.Repository {
	return bucketRepo.New(f.cfg.BaseBucketConfig())
}

func (f *factory) BucketFreeRepository() bucketFreeRepo.Repository {
	return bucketFreeRepo.New(f.cfg.FreeBucketConfig())
}
