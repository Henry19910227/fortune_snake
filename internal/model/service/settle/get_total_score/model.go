package get_total_score

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
	Score int
}

func (o *Output) GetScore() int {
	return o.Data.Score
}

func NewOutput(Score int) *Output {
	return &Output{Data: Data{Score: Score}}
}
