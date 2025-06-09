package repository

import (
	gameCfg "game_server_slots_fortune_snake/config/game"
	. "game_server_slots_fortune_snake/constants"
	bucketRepo "game_server_slots_fortune_snake/internal/repository/bucket"
	gameRepo "game_server_slots_fortune_snake/internal/repository/game"
	playerRepo "game_server_slots_fortune_snake/internal/repository/player"
	resultFreeRepo "game_server_slots_fortune_snake/internal/repository/result_free"
	resultLoader "game_server_slots_fortune_snake/internal/repository/result_loader"
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
	return gameRepo.New(f.rdb)
}

func (f *factory) PlayerRepository() playerRepo.Repository {
	return playerRepo.New(f.rdb)
}

func (f *factory) WeightRepository() weightRepo.Repository {
	return weightRepo.New(f.cfg.RTPConfig())
}

func (f *factory) ResultLoader() resultLoader.Repository {
	return resultLoader.New(f.db, SpinModeBase)
}

func (f *factory) ResultFreeLoader() resultLoader.Repository {
	return resultLoader.New(f.db, SpinModeFree)
}

func (f *factory) ResultFreeRepository() resultFreeRepo.Repository {
	return resultFreeRepo.New(f.rdb)
}

func (f *factory) BucketRepository() bucketRepo.Repository {
	return bucketRepo.New(f.cfg.BaseBucketConfig())
}

func (f *factory) BucketFreeRepository() bucketRepo.Repository {
	return bucketRepo.New(f.cfg.FreeBucketConfig())
}
