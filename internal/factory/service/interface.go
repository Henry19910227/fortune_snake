package service

import (
	playerService "game_server_slots_fortune_snake/internal/service/player"
)

type Factory interface {
	PlayerService() playerService.Service
}
