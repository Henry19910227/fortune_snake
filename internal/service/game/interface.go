package game

import (
	"game_server_slots_fortune_snake/internal/model/service/game/bet"
	"game_server_slots_fortune_snake/internal/model/service/game/enter_game"
)

type Service interface {
	EnterGame(input *enter_game.Input) (output *enter_game.Output, err error)
	Bet(input *bet.Input) (output *bet.Output, err error)
}
