package bet

import "encoding/json"

type Param struct {
	Bonus bool `json:"bonus"`
	Bet   int  `json:"bet"`
	Value int  `json:"value"`
}

func (p *Param) ToJson() []byte {
	b, _ := json.Marshal(p)
	return b
}
