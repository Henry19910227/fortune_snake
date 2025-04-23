package player

import "game_server_slots_fortune_snake/server"

type Controller interface {
	GetPlayerSession(ctx *server.Context)
}
