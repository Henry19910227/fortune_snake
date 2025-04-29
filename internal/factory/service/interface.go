package service

import (
	gameService "game_server_slots_fortune_snake/internal/service/game"
	playerService "game_server_slots_fortune_snake/internal/service/player"
)

type Factory interface {
	GameService() gameService.Service
	GameDemoService() gameService.Service
	PlayerService() playerService.Service
}
