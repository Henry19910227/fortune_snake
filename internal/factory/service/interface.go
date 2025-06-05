package service

import (
	gameService "game_server_slots_fortune_snake/internal/service/game"
	playerService "game_server_slots_fortune_snake/internal/service/player"
	resultService "game_server_slots_fortune_snake/internal/service/result"
	weightService "game_server_slots_fortune_snake/internal/service/weight"
)

type Factory interface {
	GameService() gameService.Service
	GameDemoService() gameService.Service
	PlayerService() playerService.Service
	WeightService() weightService.Service
	ResultService() resultService.Service
	ResultFreeService() resultService.Service
}
