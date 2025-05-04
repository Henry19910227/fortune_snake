package check_win_line

import (
	"game_server_slots_fortune_snake/internal/model"
	"game_server_slots_fortune_snake/internal/model/entity/line"
)

type Input struct {
	model.BaseInput
	Param Param
}

type Param struct {
	Line *line.Item
}
