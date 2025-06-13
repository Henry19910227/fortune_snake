package service

import (
	. "game_server_slots_fortune_snake/constants"
	"game_server_slots_fortune_snake/internal/factory/repository"
	gameService "game_server_slots_fortune_snake/internal/service/game"
	playerService "game_server_slots_fortune_snake/internal/service/player"
	reelsService "game_server_slots_fortune_snake/internal/service/reels"
	resultFreeService "game_server_slots_fortune_snake/internal/service/result_free"
	resultLoader "game_server_slots_fortune_snake/internal/service/result_loader"
	settleService "game_server_slots_fortune_snake/internal/service/settle"
	weightService "game_server_slots_fortune_snake/internal/service/weight"
)

type factory struct {
	repoFactory repository.Factory
}

func New(repoFactory repository.Factory) Factory {
	serviceFactory := &factory{repoFactory: repoFactory}
	return serviceFactory
}

func (f *factory) GameService() gameService.Service {
	gameRepo := f.repoFactory.GameRepository()
	sessionRepo := f.repoFactory.SessionRepository()
	weightRepo := f.repoFactory.WeightRepository()
	return gameService.NewService(gameRepo.Mode(GameModeReal), sessionRepo, weightRepo)
}

func (f *factory) GameDemoService() gameService.Service {
	gameRepo := f.repoFactory.GameRepository()
	playerRepo := f.repoFactory.PlayerRepository()
	return gameService.NewServiceDemo(gameRepo.Mode(GameModeDemo), playerRepo)
}

func (f *factory) PlayerService() playerService.Service {
	playerRepo := f.repoFactory.PlayerRepository()
	sessionRepo := f.repoFactory.SessionRepository()
	return playerService.New(playerRepo, sessionRepo)
}

func (f *factory) WeightService() weightService.Service {
	weightRepo := f.repoFactory.WeightRepository()
	return weightService.New(weightRepo)
}

func (f *factory) ResultFreeService() resultFreeService.Service {
	resultFreeRepo := f.repoFactory.ResultFreeRepository()
	return resultFreeService.New(resultFreeRepo)
}

func (f *factory) ResultLoader() resultLoader.Service {
	resultRepo := f.repoFactory.ResultLoader()
	bucketRepo := f.repoFactory.BucketRepository()
	return resultLoader.New(resultRepo, bucketRepo, SpinModeBase)
}

func (f *factory) ResultFreeLoader() resultLoader.Service {
	resultFreeRepo := f.repoFactory.ResultFreeLoader()
	bucketRepo := f.repoFactory.BucketRepository()
	return resultLoader.New(resultFreeRepo, bucketRepo, SpinModeFree)
}

func (f *factory) ReelsService() reelsService.Service {
	symbolRepo := f.repoFactory.SymbolRepository()
	return reelsService.New(symbolRepo, SpinModeBase)
}

func (f *factory) ReelsFreeService() reelsService.Service {
	symbolRepo := f.repoFactory.SymbolRepository()
	return reelsService.New(symbolRepo, SpinModeFree)
}

func (f *factory) SettleService() settleService.Service {
	settleRepo := f.repoFactory.SettleRepository()
	return settleService.New(settleRepo)
}
