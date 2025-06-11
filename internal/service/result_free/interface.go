package result_free

import "context"

type Service interface {
	GameMode(gameMode string) Service
	// SaveItems 保存金蛇多個盤面
	SaveItems(ctx context.Context, playerAccount string, items [][][]int) error
	// PopFirstItem 彈出第一筆金蛇盤面
	PopFirstItem(ctx context.Context, playerAccount string) ([][]int, error)
	// Amount 獲取剩餘盤面數量
	Amount(ctx context.Context, playerAccount string) (int64, error)
}
