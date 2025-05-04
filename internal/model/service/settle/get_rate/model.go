package get_rate

import (
	"game_server_slots_fortune_snake/internal/model"
	"game_server_slots_fortune_snake/internal/model/entity/symbol"
)

type Input struct {
	model.BaseInput
	Param Param
}

type Param struct {
	Bet   int
	Value int
	Reels [][]*symbol.Item
}
