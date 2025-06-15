package bet

import "encoding/json"

type Param struct {
	Bonus bool
	Bet   int
	Value int
}

func (p *Param) ToJson() []byte {
	b, _ := json.Marshal(p)
	return b
}
