package game

import (
	"game_server_slots_fortune_snake/internal/factory/controller"
	"game_server_slots_fortune_snake/internal/server"
)

func SetRoute(group *server.RouterGroup, factory controller.Factory) {
	playerController := factory.PlayerController()
	gameController := factory.GameController()
	middController := factory.MiddleController()
	group.EndPoint("enter_game", middController.Verify, gameController.EnterGame)
	group.EndPoint("spin", middController.Verify, gameController.Bet)
	group.EndPoint("player", playerController.GetPlayerSession) // 測試
}
