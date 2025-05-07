package get_win_lines

import (
	"game_server_slots_fortune_snake/internal/model"
	lineModel "game_server_slots_fortune_snake/internal/model/entity/line"
	"game_server_slots_fortune_snake/internal/model/entity/symbol"
)

// Input 輸入
type Input struct {
	model.BaseInput
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
	model.BaseOutput
	Data Data
}

type Data struct {
	Lines []*lineModel.Item
}

func (o *Output) GetLines() []*lineModel.Item {
	return o.Data.Lines
}

func NewOutput(lines []*lineModel.Item) *Output {
	return &Output{Data: Data{Lines: lines}}
}
