package service

import (
	"game_server_slots_fortune_snake/factory/repository"
	playerService "game_server_slots_fortune_snake/service/player"
)

type factory struct {
	repoFactory repository.Factory
}

func New(repoFactory repository.Factory) Factory {
	serviceFactory := &factory{repoFactory: repoFactory}
	return serviceFactory
}

func (f *factory) PlayerService() playerService.Service {
	playerRepo := f.repoFactory.PlayerRepository()
	return playerService.NewService(playerRepo)
}
