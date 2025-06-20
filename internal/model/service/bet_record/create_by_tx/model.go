package create_by_tx

import (
	betModel "game_server_slots_fortune_snake/internal/model/controller/game/bet"
	playerModel "game_server_slots_fortune_snake/internal/model/entity/player"
)

type Input struct {
	Session       *playerModel.Session
	Param         *betModel.Param
	TotalBet      int64
	TransactionID uint64
	RoundID       uint64
}
