package bet_record

import "game_server_slots_fortune_snake/internal/server"

type Controller interface {
	UpdateRecord(ctx *server.Context)
}
