package game

import (
	"game_server_slots_fortune_snake/factory/controller"
	"game_server_slots_fortune_snake/server"
)

func SetRoute(group *server.RouterGroup, factory controller.Factory) {
	group.EndPoint("bet", nil)
}
