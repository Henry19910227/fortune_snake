package get_rate

import (
	"game_server_slots_fortune_snake/internal/model/entity/symbol"
	"game_server_slots_fortune_snake/internal/model/service/base"
)

// Input 輸入
type Input struct {
	base.Input
	Param Param
}

type Param struct {
	Bet   int
	Value int
	Reels [][]*symbol.Item
}

func NewInput(param Param) *Input {
	return &Input{Param: param}
}

// Output 輸出
type Output struct {
	base.Output
	Data Data
}

type Data struct {
	Rate float64
}

func (o *Output) GetRate() float64 {
	return o.Data.Rate
}

func NewOutput(rate float64) *Output {
	return &Output{Data: Data{Rate: rate}}
}
