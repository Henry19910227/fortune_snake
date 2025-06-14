package bet_record

import (
	model "game_server_slots_fortune_snake/internal/model/entity/bet_record"
	playerModel "game_server_slots_fortune_snake/internal/model/entity/player"
)

type Service interface {
	Create(session *playerModel.Session) (item *model.Table, err error)
	Update(item *model.Table) (err error)
}
