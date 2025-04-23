package player

import (
	"context"
	playerModel "game_server_slots_fortune_snake/model/player"
)

type Repository interface {
	// FindPlayerSessionById 從 cache 獲取玩家數據
	FindPlayerSessionById(ctx context.Context, playerId uint64) (*playerModel.Session, error)
}
