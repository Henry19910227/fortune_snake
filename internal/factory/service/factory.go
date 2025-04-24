package service

import (
	"game_server_slots_fortune_snake/internal/factory/repository"
	"game_server_slots_fortune_snake/internal/service/player"
)

type factory struct {
	repoFactory repository.Factory
}

func New(repoFactory repository.Factory) Factory {
	serviceFactory := &factory{repoFactory: repoFactory}
	return serviceFactory
}

func (f *factory) PlayerService() player.Service {
	playerRepo := f.repoFactory.PlayerRepository()
	return player.NewService(playerRepo)
}
