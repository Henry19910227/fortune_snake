package player

import (
	"game_server_slots_fortune_snake/internal/server"
)

type Controller interface {
	GetPlayerSession(ctx *server.Context)
}
