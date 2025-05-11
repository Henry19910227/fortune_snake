package save

import (
	model "game_server_slots_fortune_snake/internal/model/entity/result_free"
	"game_server_slots_fortune_snake/internal/model/service/base"
)

type Input struct {
	base.Input
	Param Param
}

type Param struct {
	Results [][]*model.Item
}
