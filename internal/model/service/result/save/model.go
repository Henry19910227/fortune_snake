package save

import (
	"game_server_slots_fortune_snake/internal/model/entity/result"
	"game_server_slots_fortune_snake/internal/model/service/base"
)

type Input struct {
	base.Input
	Param Param
}

type Param struct {
	Results [][]*result.Item
}
