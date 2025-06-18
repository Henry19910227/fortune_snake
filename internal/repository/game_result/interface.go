package game_result

import model "game_server_slots_fortune_snake/internal/model/entity/game_result"

type Repository interface {
	Create(item *model.Table) (id uint64, err error)
}
