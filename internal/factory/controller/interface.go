package controller

import (
	gameController "game_server_slots_fortune_snake/internal/controller/game"
	middleController "game_server_slots_fortune_snake/internal/controller/middleware"
	playerController "game_server_slots_fortune_snake/internal/controller/player"
)

type Factory interface {
	GameController() gameController.Controller
	PlayerController() playerController.Controller
	MiddleController() middleController.Controller
}
