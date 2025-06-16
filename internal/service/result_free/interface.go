package result_free

import "context"

type Service interface {
	// SaveItems 保存金蛇多個盤面
	SaveItems(ctx context.Context, gameMode string, playerID uint64, items [][][]int) error
	// PopFirstItem 彈出第一筆金蛇盤面
	PopFirstItem(ctx context.Context, gameMode string, playerID uint64) ([][]int, error)
	// Amount 獲取剩餘盤面數量
	Amount(ctx context.Context, gameMode string, playerID uint64) (int64, error)
}
