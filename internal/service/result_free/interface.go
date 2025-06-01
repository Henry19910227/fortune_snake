package result

import (
	model "game_server_slots_fortune_snake/internal/model/entity/result_free"
	"game_server_slots_fortune_snake/internal/model/service/result_free/save_to_bucket"
)

type Service interface {
	// SaveToBucket 將 result 數據存到本地 bucket 中
	SaveToBucket(input *save_to_bucket.Input) (output *save_to_bucket.Output)
	// Migrate 將 bucket 中的數據存至 DB
	Migrate() (err error)
	// LoadData 將盤面結果載入至內存
	LoadData() (err error)
	// Random 獲取隨機盤面結果
	Random(rate float64) (output *model.Item, err error)
}
