package player

import (
	"context"
	playerModel "game_server_slots_fortune_snake/internal/model/entity/player"
	"gorm.io/gorm"
)

type Service interface {
	Tx(tx *gorm.DB) Service
	GetPlayerSession(ctx context.Context, playerId uint64) (output *playerModel.Session, err error)
	UpdateBalance(ctx context.Context, playerId uint64, value int64) error
}
