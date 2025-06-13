package player

import (
	playerModel "game_server_slots_fortune_snake/internal/model/entity/player"
	"game_server_slots_fortune_snake/internal/server"
	"gorm.io/gorm"
)

type Service interface {
	Tx(tx *gorm.DB) Service
	GetPlayerSession(ctx *server.Context, playerId uint64) (output *playerModel.Session, err error)
	UpdateBalance(ctx *server.Context, playerId uint64, value int64) error
}
