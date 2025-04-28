package service

import (
	"game_server_slots_fortune_snake/internal/factory/repository"
	gameService "game_server_slots_fortune_snake/internal/service/game"
	playerService "game_server_slots_fortune_snake/internal/service/player"
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
	return gameService.NewService(gameRepo)
}

func (f *factory) PlayerService() playerService.Service {
	playerRepo := f.repoFactory.PlayerRepository()
	return playerService.NewService(playerRepo)
}
