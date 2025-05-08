package info

import (
	gameModel "game_server_slots_fortune_snake/internal/model/entity/game"
	"game_server_slots_fortune_snake/internal/model/repository/base"
)

// Input 輸入
type Input struct {
	base.Input
}

func NewInput() *Input {
	return &Input{}
}

// Output 輸出
type Output struct {
	base.Output
	Data Data
}

type Data struct {
	Info *gameModel.Info
}

func (o *Output) GetInfo() *gameModel.Info {
	return o.Data.Info
}

func NewOutput(info *gameModel.Info) *Output {
	return &Output{Data: Data{Info: info}}
}
