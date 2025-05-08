package result

import (
	"game_server_slots_fortune_snake/internal/model/repository/result/create_items"
)

// Repository 存取盤面結果
type Repository interface {
	// CreateItems 創建多筆盤面結果至db
	CreateItems(input *create_items.Input) (err error)
}
