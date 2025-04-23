package controller

import (
	middleController "game_server_slots_fortune_snake/controller/middleware"
	playerController "game_server_slots_fortune_snake/controller/player"
	serviceFactory "game_server_slots_fortune_snake/factory/service"
)

type factory struct {
	serviceFactory serviceFactory.Factory
}

func New(serviceFactory serviceFactory.Factory) Factory {
	return &factory{serviceFactory: serviceFactory}
}

func (f *factory) PlayerController() playerController.Controller {
	return playerController.New()
}

func (f *factory) MiddleController() middleController.Controller {
	return middleController.New(f.serviceFactory.PlayerService())
}
