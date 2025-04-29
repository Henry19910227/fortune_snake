package bet

import (
	"game_server_slots_fortune_snake/internal/model"
	playerModel "game_server_slots_fortune_snake/internal/model/player"
)

type Input struct {
	model.BaseInput
	Session *playerModel.Session
	Param   *Param
}

type Output struct {
	model.BaseOutput
	Data *Data
}

type Param struct {
	Bet   int
	Value int
}

type Data struct{}
