package player

import (
	"game_server_slots_fortune_snake/internal/model/service/player/get_player_session"
)

type Service interface {
	GetPlayerSession(input *get_player_session.Input) (output *get_player_session.Output, err error)
}
