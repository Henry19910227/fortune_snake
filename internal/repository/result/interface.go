package result

import (
	"game_server_slots_fortune_snake/internal/model/repository/result/create_items"
	"game_server_slots_fortune_snake/internal/model/repository/result/random"
)

// Repository 存取盤面結果
type Repository interface {
	// CreateItems 創建多筆盤面結果至db
	CreateItems(input *create_items.Input) (err error)
	// LoadData 將盤面結果載入至內存
	LoadData() (err error)
	// Random 以 rate 參數獲取該賠率隨機盤面
	Random(input *random.Input) (output *random.Output, err error)
}
