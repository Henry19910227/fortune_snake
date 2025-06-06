package game

import (
	"context"
	gameModel "game_server_slots_fortune_snake/internal/model/entity/game"
)

type Repository interface {
	// Info 獲取遊戲配置數據
	Info() (info *gameModel.Info, err error)
	// SetSpecialMode 設置特殊模式
	SetSpecialMode(ctx context.Context, playerId uint64, specialMode bool) error
	// IsSpecialMode 是否是特殊模式
	IsSpecialMode(ctx context.Context, playerId uint64) (bool, error)

	RestoreResults()
}
