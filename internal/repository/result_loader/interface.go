package result

import (
	. "game_server_slots_fortune_snake/constants"
	model "game_server_slots_fortune_snake/internal/model/entity/result"
	"gorm.io/gorm"
)

// Repository 存取盤面結果
type Repository interface {
	// CreateItems 創建多筆盤面結果至db
	CreateItems(items []*model.Item) (err error)
	// LoadData 將盤面結果載入至內存
	LoadData() (err error)
	// Random 以 rate 參數獲取該賠率隨機盤面
	Random(rate float64) (*model.Item, error)
}

func New(db *gorm.DB, spinMode int) Repository {
	if spinMode == SpinModeBase {
		return &repository{db: db, resultsMap: make(map[float64][]*model.Item)}
	}
	return &repositoryFree{db: db, resultsMap: make(map[float64][]*model.Item)}
}
