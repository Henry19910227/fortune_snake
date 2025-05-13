package random

import (
	"game_server_slots_fortune_snake/internal/model/service/base"
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
	Symbols [][]int
}

func (o *Output) GetSymbols() [][]int {
	return o.Data.Symbols
}

func NewOutput(symbols [][]int) *Output {
	return &Output{Data: Data{Symbols: symbols}}
}
