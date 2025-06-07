package result

import (
	"game_server_slots_fortune_snake/internal/model/entity/result"
)

type Service interface {
	// SaveToBucket 將 result 數據存到本地 bucket 中
	SaveToBucket(result *result.Item) (quota int)
	// Migrate 將 bucket 中的數據存至 DB
	Migrate() (err error)
	// LoadData 將盤面結果載入至內存
	LoadData() (err error)
	// Random 獲取隨機盤面結果
	Random(rate float64) ([][][]int, error)
}
