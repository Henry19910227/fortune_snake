package get_win_lines

import (
	lineModel "game_server_slots_fortune_snake/internal/model/entity/line"
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
	Lines []*lineModel.Item
}

func (o *Output) GetLines() []*lineModel.Item {
	return o.Data.Lines
}

func NewOutput(lines []*lineModel.Item) *Output {
	return &Output{Data: Data{Lines: lines}}
}
