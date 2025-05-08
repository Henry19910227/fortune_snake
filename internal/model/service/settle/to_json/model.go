package to_json

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
	Items [][]*symbol.Item
}

func NewInput(Items [][]*symbol.Item) *Input {
	return &Input{Param: Param{Items: Items}}
}

// Output 輸出
type Output struct {
	base.Output
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
