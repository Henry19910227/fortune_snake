package game

import (
	"game_server_slots_fortune_snake/internal/server"
	gameService "game_server_slots_fortune_snake/internal/service/game"
)

type controller struct {
	gameService gameService.Service
}

func New(gameService gameService.Service) Controller {
	return &controller{gameService: gameService}
}

func (c *controller) EnterGame(ctx *server.Context) {
	//TODO implement me
	panic("implement me")
}
