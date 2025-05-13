package result

import (
	model "game_server_slots_fortune_snake/internal/model/entity/result"
	"game_server_slots_fortune_snake/internal/model/repository/result/create_items"
)

// Repository 存取盤面結果
type Repository interface {
	// CreateItems 創建多筆盤面結果至db
	CreateItems(input *create_items.Input) (err error)
	// LoadData 將盤面結果載入至內存
	LoadData() (err error)
	// Random 以 rate 參數獲取該賠率隨機盤面
	Random(rate float64) (*model.Item, error)
}
