package bet

import lineModel "game_server_slots_fortune_snake/internal/model/entity/line"

type SpinResult struct {
	Symbols [][]int
	Lines   []*Line
	Score   int // 總中獎金額
	Times   int
}

func (s *SpinResult) SetLines(items []*lineModel.Item) {
	lines := make([]*Line, 0)
	for _, item := range items {
		line := NewLine(item)
		lines = append(lines, line)
	}
	s.Lines = lines
}
