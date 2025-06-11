package bet

import lineModel "game_server_slots_fortune_snake/internal/model/entity/line"

type SpinResult struct {
	Symbols [][]int
	Lines   []*Line
	Score   int // 總中獎金額
	Times   int
}

func NewSpinResult(items []*lineModel.Item) *SpinResult {
	lines := make([]*Line, 0)
	for _, item := range items {
		line := NewLine(item)
		lines = append(lines, line)
	}
	spinResult := &SpinResult{}
	spinResult.Lines = lines
	return spinResult
}
