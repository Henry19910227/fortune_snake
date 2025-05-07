package result

import (
	"game_server_slots_fortune_snake/internal/model/service/result/save"
	"game_server_slots_fortune_snake/internal/model/service/result/save_to_bucket"
)

type Service interface {
	// Save 儲存 result 數據
	Save(input *save.Input) error
	// SaveToBucket 將 result 數據存到本地 bucket 中
	SaveToBucket(input *save_to_bucket.Input) (output *save_to_bucket.Output)
}
