package bet

import lineModel "game_server_slots_fortune_snake/internal/model/entity/line"

type Line struct {
	HitIndex int    // 賠付線的位置
	Cells    []Cell // 中獎符號位置
	Pow      int    // 賠付值
	Symbol   int    // 中獎符號
	Score    int    // 中獎金額
}

type Cell struct {
	Col    int
	Row    int
	Symbol int
}

func NewLine(item *lineModel.Item) *Line {
	line := &Line{}
	line.HitIndex = item.Index
	line.Score = item.Score
	line.Symbol = item.Symbol.ID
	line.Pow = item.Symbol.Pow

	cells := make([]Cell, 0)
	for _, pos := range item.Positions {
		cell := Cell{}
		cell.Row = pos.Row
		cell.Col = pos.Col
		cell.Symbol = pos.Symbol.ID
		cells = append(cells, cell)
	}
	line.Cells = cells

	return line
}
