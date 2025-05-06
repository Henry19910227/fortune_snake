package result

import model "game_server_slots_fortune_snake/internal/model/entity/result"

// Repository 存取盤面結果
type Repository interface {
	Create(item *model.Item) (id int64, err error)
}
