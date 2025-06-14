package result_free

import "context"

type Service interface {
	GameMode(gameMode string) Service
	// SaveItems 保存金蛇多個盤面
	SaveItems(ctx context.Context, playerID uint64, items [][][]int) error
	// PopFirstItem 彈出第一筆金蛇盤面
	PopFirstItem(ctx context.Context, playerID uint64) ([][]int, error)
	// Amount 獲取剩餘盤面數量
	Amount(ctx context.Context, playerID uint64) (int64, error)
}
