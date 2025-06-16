package bet

import (
	"encoding/json"
	lineModel "game_server_slots_fortune_snake/internal/model/entity/line"
)

type SpinResult struct {
	Symbols [][]int `json:"symbols"`
	Lines   []*Line `json:"lines"`
	Score   int     `json:"score"` // 總中獎金額
	Times   int     `json:"times"`
}

func (s *SpinResult) SetLines(items []*lineModel.Item) {
	lines := make([]*Line, 0)
	for _, item := range items {
		line := NewLine(item)
		lines = append(lines, line)
	}
	s.Lines = lines
}

func (s *SpinResult) ToJson() []byte {
	b, _ := json.Marshal(s)
	return b
}
