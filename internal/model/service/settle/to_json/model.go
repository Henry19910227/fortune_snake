package to_json

import (
	"game_server_slots_fortune_snake/internal/model"
	"game_server_slots_fortune_snake/internal/model/entity/symbol"
)

// Input 輸入
type Input struct {
	model.BaseInput
	Param Param
}

type Param struct {
	Items [][]*symbol.Item
}

func NewInput(Items [][]*symbol.Item) *Input {
	return &Input{Param: Param{Items: Items}}
}

// Output 輸出
type Output struct {
	model.BaseOutput
	Data Data
}

type Data struct {
	JsonString string
}

func NewOutput(JsonString string) *Output {
	return &Output{Data: Data{JsonString: JsonString}}
}

func (o *Output) GetJson() string {
	return o.Data.JsonString
}
