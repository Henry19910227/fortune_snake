package check_win_line

import (
	"game_server_slots_fortune_snake/internal/model"
	"game_server_slots_fortune_snake/internal/model/entity/line"
	"game_server_slots_fortune_snake/internal/model/entity/symbol"
)

// Input 輸入
type Input struct {
	model.BaseInput
	Param Param
}

type Param struct {
	Line *line.Item
}

func NewInput(line *line.Item) *Input {
	return &Input{Param: Param{Line: line}}
}

// Output 輸出
type Output struct {
	model.BaseOutput
	Data Data
}

type Data struct {
	Symbol *symbol.Item
}

func (o *Output) GetSymbol() *symbol.Item {
	return o.Data.Symbol
}

func NewOutput(symbol *symbol.Item) *Output {
	return &Output{Data: Data{Symbol: symbol}}
}
