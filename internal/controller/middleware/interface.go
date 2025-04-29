package middleware

import (
	"game_server_slots_fortune_snake/internal/server"
)

type Controller interface {
	Verify(ctx *server.Context)
}
