package get_player_session

import (
	"game_server_slots_fortune_snake/model"
	playerModel "game_server_slots_fortune_snake/model/player"
)

type Input struct {
	model.BaseInput
	PlayerId uint64
}

type Output struct {
	model.BaseOutput
	Data *playerModel.Session
}
