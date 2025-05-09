package save_to_bucket

import (
	"game_server_slots_fortune_snake/internal/model/entity/result"
	"game_server_slots_fortune_snake/internal/model/service/base"
)

// Input 輸入
type Input struct {
	base.Input
	Param Param
}

type Param struct {
	Result *result.Item
}

func NewInput(result *result.Item) *Input {
	return &Input{Param: Param{Result: result}}
}

// Output 輸出
type Output struct {
	base.Output
	Data Data
}

type Data struct {
	Quota int
}

func (o *Output) GetQuota() int {
	return o.Data.Quota
}

func NewOutput(quota int) *Output {
	return &Output{Data: Data{Quota: quota}}
}
