package controller

import (
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
	weightSvc := f.serviceFactory.WeightService()
	resultFreeSvc := f.serviceFactory.ResultFreeService()
	resultLoader := f.serviceFactory.ResultLoader()
	resultFreeLoader := f.serviceFactory.ResultFreeLoader()
	reelsSvc := f.serviceFactory.ReelsService()
	reelsFreeSvc := f.serviceFactory.ReelsFreeService()
	settleSvc := f.serviceFactory.SettleService()
	playerSvc := f.serviceFactory.PlayerService()
	return gameController.New(gameSvc, weightSvc, resultFreeSvc, resultLoader, resultFreeLoader, reelsSvc, reelsFreeSvc, settleSvc, playerSvc)
}

func (f *factory) PlayerController() player.Controller {
	return player.New(f.serviceFactory.PlayerService())
}

func (f *factory) MiddleController() middleware.Controller {
	return middleware.New(f.serviceFactory.PlayerService())
}

func (f *factory) LoadController() loadController.Controller {
	weightService := f.serviceFactory.WeightService()
	resultLoader := f.serviceFactory.ResultLoader()
	resultFreeLoader := f.serviceFactory.ResultFreeLoader()
	return loadController.New(weightService, resultLoader, resultFreeLoader)
}
