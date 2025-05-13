package controller

import (
	"game_server_slots_fortune_snake/internal/controller/game"
	gameController "game_server_slots_fortune_snake/internal/controller/game"
	loadController "game_server_slots_fortune_snake/internal/controller/load"
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

func (f *factory) GameController() gameController.Controller {
	gameSvc := f.serviceFactory.GameService()
	gameDemoSvc := f.serviceFactory.GameDemoService()
	weightSvc := f.serviceFactory.WeightService()
	resultSvc := f.serviceFactory.ResultService()
	resultFreeSvc := f.serviceFactory.ResultFreeService()
	return game.New(gameSvc, gameDemoSvc, weightSvc, resultSvc, resultFreeSvc)
}

func (f *factory) PlayerController() player.Controller {
	return player.New(f.serviceFactory.PlayerService())
}

func (f *factory) MiddleController() middleware.Controller {
	return middleware.New(f.serviceFactory.PlayerService())
}

func (f *factory) LoadController() loadController.Controller {
	weightService := f.serviceFactory.WeightService()
	resultService := f.serviceFactory.ResultService()
	resultFreeService := f.serviceFactory.ResultFreeService()
	return loadController.New(weightService, resultService, resultFreeService)
}
