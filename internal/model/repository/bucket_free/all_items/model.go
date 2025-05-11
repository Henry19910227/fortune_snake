package all_items

import (
	resultModel "game_server_slots_fortune_snake/internal/model/entity/result_free"
	"game_server_slots_fortune_snake/internal/model/repository/base"
)

// Input 輸入
type Input struct {
	base.Input
}

// Output 輸出
type Output struct {
	base.Output
	Data Data
}

type Data struct {
	Items []*resultModel.Item
}

func (o *Output) Items() []*resultModel.Item {
	return o.Data.Items
}

func NewOutput(items []*resultModel.Item) *Output {
	return &Output{Data: Data{Items: items}}
}
