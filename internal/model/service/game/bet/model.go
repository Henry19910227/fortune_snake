package bet

import (
	playerModel "game_server_slots_fortune_snake/internal/model/entity/player"
	"game_server_slots_fortune_snake/internal/model/service/base"
)

type Input struct {
	base.Input
	Session *playerModel.Session
	Param   *Param
}

type Output struct {
	Data *Data
}

type Param struct {
	Bet   int
	Value int
}

type Data struct{}
