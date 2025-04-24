package middleware

import (
	"game_server_slots_fortune_snake/internal/server"
)

type Controller interface {
	UnMarshalData(ctx *server.Context)
	UnMarshalReq(ctx *server.Context)
	Verify(ctx *server.Context)
}
