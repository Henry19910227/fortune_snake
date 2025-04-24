package controller

import (
	middleController "game_server_slots_fortune_snake/internal/controller/middleware"
	playerController "game_server_slots_fortune_snake/internal/controller/player"
)

type Factory interface {
	PlayerController() playerController.Controller
	MiddleController() middleController.Controller
}
