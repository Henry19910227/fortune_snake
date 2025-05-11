package random

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
	Rate float64
}

func NewInput(rate float64) *Input {
	return &Input{Param: Param{Rate: rate}}
}

// Output 輸出
type Output struct {
	base.Output
	Data Data
}

type Data struct {
	Item *model.Item
}

func (o *Output) GetItem() *model.Item {
	return o.Data.Item
}

func NewOutput(item *model.Item) *Output {
	return &Output{Data: Data{Item: item}}
}
