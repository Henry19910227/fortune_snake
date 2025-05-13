package random

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
	Symbols    [][]*symbol.Item
	AllSymbols [][][]*symbol.Item
}

func (o *Output) GetSymbols() [][]*symbol.Item {
	return o.Data.Symbols
}

func NewOutput(symbols [][]*symbol.Item) *Output {
	return &Output{Data: Data{Symbols: symbols}}
}
