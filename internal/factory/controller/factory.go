package controller

import (
	"game_server_slots_fortune_snake/internal/controller/middleware"
	"game_server_slots_fortune_snake/internal/controller/player"
	serviceFactory "game_server_slots_fortune_snake/internal/factory/service"
)

type factory struct {
	serviceFactory serviceFactory.Factory
}

func New(serviceFactory serviceFactory.Factory) Factory {
	return &factory{serviceFactory: serviceFactory}
}

func (f *factory) PlayerController() player.Controller {
	return player.New(f.serviceFactory.PlayerService())
}

func (f *factory) MiddleController() middleware.Controller {
	return middleware.New(f.serviceFactory.PlayerService())
}
