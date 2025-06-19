package player_free_order

import model "game_server_slots_fortune_snake/internal/model/entity/player_free_order"

type Repository interface {
	Create(item *model.Table) (id uint64, err error)
}
