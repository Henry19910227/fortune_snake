package game

import (
	"game_server_slots_fortune_snake/internal/factory/controller"
	"game_server_slots_fortune_snake/internal/server"
)

func SetRoute(group *server.RouterGroup, factory controller.Factory) {
	gameController := factory.GameController()
	middController := factory.MiddleController()
	group.EndPoint("bet", middController.Verify, gameController.PreBet, gameController.Bet)
	group.EndPoint("symbols", middController.Verify, gameController.Symbols)
}
