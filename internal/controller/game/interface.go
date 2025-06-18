package game

import "game_server_slots_fortune_snake/internal/server"

type Controller interface {
	PreBet(ctx *server.Context)
	Bet(ctx *server.Context)
	Symbols(ctx *server.Context)
}
