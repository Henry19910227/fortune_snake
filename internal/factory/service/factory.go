package service

import (
	. "game_server_slots_fortune_snake/constants"
	"game_server_slots_fortune_snake/internal/factory/repository"
	gameService "game_server_slots_fortune_snake/internal/service/game"
	playerService "game_server_slots_fortune_snake/internal/service/player"
	resultLoader "game_server_slots_fortune_snake/internal/service/result_loader"
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
	playerRepo := f.repoFactory.PlayerRepository()
	weightRepo := f.repoFactory.WeightRepository()
	return gameService.NewService(gameRepo.Mode(GameModeReal), playerRepo, weightRepo)
}

func (f *factory) GameDemoService() gameService.Service {
	gameRepo := f.repoFactory.GameRepository()
	playerRepo := f.repoFactory.PlayerRepository()
	return gameService.NewServiceDemo(gameRepo.Mode(GameModeDemo), playerRepo)
}

func (f *factory) PlayerService() playerService.Service {
	playerRepo := f.repoFactory.PlayerRepository()
	return playerService.NewService(playerRepo)
}

func (f *factory) WeightService() weightService.Service {
	weightRepo := f.repoFactory.WeightRepository()
	return weightService.New(weightRepo)
}

func (f *factory) ResultLoader() resultLoader.Service {
	resultRepo := f.repoFactory.ResultRepository()
	bucketRepo := f.repoFactory.BucketRepository()
	return resultLoader.SpinMode(SpinModeNormal).Init(resultRepo, bucketRepo)
}

func (f *factory) ResultFreeLoader() resultLoader.Service {
	resultFreeRepo := f.repoFactory.ResultFreeRepository()
	bucketRepo := f.repoFactory.BucketRepository()
	return resultLoader.SpinMode(SpinModeNormal).Init(resultFreeRepo, bucketRepo)
}
