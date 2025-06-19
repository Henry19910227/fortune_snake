package result_free

import (
	"context"
	"game_server_slots_fortune_snake/internal/model/entity/player"
	"game_server_slots_fortune_snake/internal/model/service/result_free/save_items"
)

type Service interface {
	// SaveItems 保存金蛇多個盤面
	SaveItems(param save_items.Param) error
	// PopFirstItem 彈出第一筆金蛇盤面
	PopFirstItem(ctx context.Context, session *player.Session) ([][]int, error)
	// Amount 獲取剩餘盤面數量
	Amount(ctx context.Context, session *player.Session) (int64, error)
}
