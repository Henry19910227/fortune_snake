package create_items

import (
	model "game_server_slots_fortune_snake/internal/model/entity/result_free"
	"game_server_slots_fortune_snake/internal/model/repository/base"
)

// Input 輸入
type Input struct {
	base.Input
	Param Param
}

type Param struct {
	Items []*model.Item
}

func NewInput(items []*model.Item) *Input {
	return &Input{Param: Param{Items: items}}
}

func (i *Input) GetItems() []*model.Item {
	return i.Param.Items
}
