package bet_record

import (
	betModel "game_server_slots_fortune_snake/internal/model/controller/game/bet"
	model "game_server_slots_fortune_snake/internal/model/entity/bet_record"
	playerModel "game_server_slots_fortune_snake/internal/model/entity/player"
)

type Service interface {
	Create(session *playerModel.Session, param *betModel.Param, totalBet int64) (item *model.Table, err error)
	CreateByTransactionID(session *playerModel.Session, param *betModel.Param, totalBet int64, transactionID uint64) (item *model.Table, err error)
	Update(item *model.Table) (err error)
	UpdateToFailed(item *model.Table) (err error)
	UpdateToFinished(item *model.Table) (err error)
}
