package bet

import (
	"encoding/json"
	"game_server_slots_fortune_snake/constants"
)

type Param struct {
	Bonus      bool `json:"bonus"`
	Bet        int  `json:"bet"`
	Value      int  `json:"value"`
	Multiplier int  `json:"multiplier"`
}

func NewParam() *Param {
	return &Param{Multiplier: constants.Multiplier}
}

func (p *Param) ToJson() []byte {
	b, _ := json.Marshal(p)
	return b
}
