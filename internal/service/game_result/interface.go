package game_result

import (
	betModel "game_server_slots_fortune_snake/internal/model/controller/game/bet"
	playerModel "game_server_slots_fortune_snake/internal/model/entity/player"
)

type Service interface {
	Create(session *playerModel.Session, result *betModel.GameResult) (id uint64, err error)
}
