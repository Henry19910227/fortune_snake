package save

import (
	"game_server_slots_fortune_snake/internal/model"
	"game_server_slots_fortune_snake/internal/model/entity/result"
)

type Input struct {
	model.BaseInput
	Param Param
}

type Param struct {
	Results [][]*result.Item
}
