package repository

import (
	gameCfg "game_server_slots_fortune_snake/config/game"
	. "game_server_slots_fortune_snake/constants"
	"game_server_slots_fortune_snake/internal/model/config/system"
	betRecordRepo "game_server_slots_fortune_snake/internal/repository/bet_record"
	bucketRepo "game_server_slots_fortune_snake/internal/repository/bucket"
	gameRepo "game_server_slots_fortune_snake/internal/repository/game"
	gameResultRepo "game_server_slots_fortune_snake/internal/repository/game_result"
	playerRepo "game_server_slots_fortune_snake/internal/repository/player"
	freeOrder "game_server_slots_fortune_snake/internal/repository/player_free_order"
	session "game_server_slots_fortune_snake/internal/repository/player_session"
	resultFreeRepo "game_server_slots_fortune_snake/internal/repository/result_free"
	resultLoader "game_server_slots_fortune_snake/internal/repository/result_loader"
	settleRepository "game_server_slots_fortune_snake/internal/repository/settle"
	snowFlakeRepository "game_server_slots_fortune_snake/internal/repository/snow_flake"
	symbolRepo "game_server_slots_fortune_snake/internal/repository/symbol"
	weightRepository "game_server_slots_fortune_snake/internal/repository/weight"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type factory struct {
	db               *gorm.DB
	rdb              *redis.Client
	cfg              gameCfg.Config
	snowFlakeRepo    snowFlakeRepository.Repository
	weightRepo       weightRepository.Repository
	resultLoader     resultLoader.Repository
	resultFreeLoader resultLoader.Repository
}

func New(db *gorm.DB, rdb *redis.Client, cfg gameCfg.Config, config *system.Config) (Factory, error) {
	weightRepo := weightRepository.New(cfg.RTPConfig())
	resultLoad := resultLoader.New(db, SpinModeBase)
	resultFreeLoad := resultLoader.New(db, SpinModeFree)
	snowFlakeRepo, err := snowFlakeRepository.New(config.Server.ServerNode)
	if err != nil {
		return nil, err
	}
	repoFactory := &factory{db: db, rdb: rdb, cfg: cfg, snowFlakeRepo: snowFlakeRepo, weightRepo: weightRepo, resultLoader: resultLoad, resultFreeLoader: resultFreeLoad}
	return repoFactory, nil
}

func (f *factory) GameRepository() gameRepo.Repository {
	return gameRepo.New(f.rdb)
}

func (f *factory) PlayerRepository() playerRepo.Repository {
	return playerRepo.New(f.db)
}

func (f *factory) SessionRepository() session.Repository {
	return session.New(f.rdb)
}

func (f *factory) SnowflakeRepository() snowFlakeRepository.Repository {
	return f.snowFlakeRepo
}

func (f *factory) WeightRepository() weightRepository.Repository {
	return f.weightRepo
}

func (f *factory) ResultLoader() resultLoader.Repository {
	return f.resultLoader
}

func (f *factory) ResultFreeLoader() resultLoader.Repository {
	return f.resultFreeLoader
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

func (f *factory) SymbolRepository() symbolRepo.Repository {
	return symbolRepo.New(f.cfg.SymbolConfig())
}

func (f *factory) SettleRepository() settleRepository.Repository {
	return settleRepository.New()
}

func (f *factory) BetRecordRepository() betRecordRepo.Repository {
	return betRecordRepo.New(f.db)
}

func (f *factory) GameResultRepository() gameResultRepo.Repository {
	return gameResultRepo.New(f.db)
}

func (f *factory) FreeOrderRepository() freeOrder.Repository {
	return freeOrder.New(f.db)
}
