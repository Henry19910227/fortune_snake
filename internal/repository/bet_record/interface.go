package bet_record

import model "game_server_slots_fortune_snake/internal/model/entity/bet_record"

type Repository interface {
	Create(item *model.Table) (id uint64, err error)
	Update(item *model.Table) (err error)
}
