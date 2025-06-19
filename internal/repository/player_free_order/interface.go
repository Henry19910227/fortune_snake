package player_free_order

import (
	"game_server_slots_fortune_snake/internal/model/entity/player"
	model "game_server_slots_fortune_snake/internal/model/entity/player_free_order"
)

type Repository interface {
	Create(item *model.Table) (id uint64, err error)
	Find(session *player.Session) (item *model.Table, err error)
	Update(item *model.Table) (err error)
}
