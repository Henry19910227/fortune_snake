package result_free

import (
	"context"
)

// Repository 存取盤面結果
type Repository interface {
	GameMode(gameMode string) Repository

	SaveItems(ctx context.Context, playerId int, items []string) error
	// PopFirstItem 彈出第一筆金蛇盤面
	PopFirstItem(ctx context.Context, playerId int) (string, error)
}
