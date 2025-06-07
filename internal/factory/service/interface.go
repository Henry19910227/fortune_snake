package service

import (
	gameService "game_server_slots_fortune_snake/internal/service/game"
	playerService "game_server_slots_fortune_snake/internal/service/player"
	resultLoader "game_server_slots_fortune_snake/internal/service/result_loader"
	weightService "game_server_slots_fortune_snake/internal/service/weight"
)

type Factory interface {
	GameService() gameService.Service
	GameDemoService() gameService.Service
	PlayerService() playerService.Service
	WeightService() weightService.Service
	ResultLoader() resultLoader.Service
	ResultFreeLoader() resultLoader.Service
}
