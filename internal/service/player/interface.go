package player

import (
	playerModel "game_server_slots_fortune_snake/internal/model/entity/player"
	"game_server_slots_fortune_snake/internal/server"
)

type Service interface {
	GetPlayerSession(ctx *server.Context, playerId uint64) (output *playerModel.Session, err error)
}
