package get_player_session

import (
	playerModel "game_server_slots_fortune_snake/internal/model/entity/player"
	"game_server_slots_fortune_snake/internal/model/service/base"
)

type Input struct {
	base.Input
	PlayerId uint64
}

type Output struct {
	Session *playerModel.Session
}
