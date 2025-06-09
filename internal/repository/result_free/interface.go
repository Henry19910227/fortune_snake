package result_free

import (
	"context"
)

// Repository 存取盤面結果
type Repository interface {
	GameMode(gameMode string) Repository
	// SaveItems 保存金蛇多個盤面
	SaveItems(ctx context.Context, playerId int, items []string) error
	// PopFirstItem 彈出第一筆金蛇盤面
	PopFirstItem(ctx context.Context, playerId int) (string, error)
	// Amount 獲取剩餘盤面數量
	Amount(ctx context.Context, playerId int) (int64, error)
}
