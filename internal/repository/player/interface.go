package player

import (
	"context"
	playerModel "game_server_slots_fortune_snake/internal/model/entity/player"
)

type Repository interface {
	// FindPlayerSessionById 從 cache 獲取玩家數據
	FindPlayerSessionById(ctx context.Context, playerId uint64) (*playerModel.Session, error)
	// Balance 獲取用戶餘額
	Balance() (balance int, err error)
	// GameData 獲取用戶遊戲中的緩存數據
	GameData() (*playerModel.GameData, error)
}
