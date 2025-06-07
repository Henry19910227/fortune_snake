package game

import (
	"context"
	gameModel "game_server_slots_fortune_snake/internal/model/entity/game"
)

type Repository interface {
	Mode(gameMode string) Repository
	// Info 獲取遊戲配置數據
	Info() (info *gameModel.Info, err error)
	// SetSpecialMode 設置特殊模式
	SetSpecialMode(ctx context.Context, playerId uint64, specialMode bool) error
	// IsSpecialMode 是否是特殊模式
	IsSpecialMode(ctx context.Context, playerId uint64) (bool, error)
	// SaveFreeResults 保存金蛇多個盤面
	SaveFreeResults(ctx context.Context, playerId int, items []string) error
	// PopFirstResult 獲取第一筆金蛇盤面
	PopFirstResult(ctx context.Context, playerId int) (string, error)
}
