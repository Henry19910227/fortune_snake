package settle

import (
	"errors"
	"game_server_slots_fortune_snake/internal/model/settle"
	"game_server_slots_fortune_snake/internal/model/symbol"
)

type repository struct {
	hitLines [][]int
}

func New() Repository {
	hitLines := [][]int{
		{0, 0, 0}, // Line 1
		{0, 1, 0}, // Line 2
		{0, 1, 1}, // Line 3
		{1, 1, 0}, // Line 4
		{1, 1, 1}, // Line 5
		{1, 2, 1}, // Line 6
		{1, 2, 2}, // Line 7
		{2, 2, 1}, // Line 8
		{2, 2, 2}, // Line 9
		{2, 3, 2}, // Line 10
	}
	return &repository{hitLines: hitLines}
}

func (r *repository) GetWinLines(bet int, value int, reelSet [][]*symbol.Item) ([]*settle.Line, error) {
	// 判斷軸數是否相同
	if len(reelSet) != len(r.hitLines[0]) {
		return []*settle.Line{}, errors.New("盤面格式不符")
	}
	// 依照 hitLines 中的點位組成 Line
	winLines := make([]*settle.Line, 0)
	for i := 0; i < len(r.hitLines); i++ {
		line := &settle.Line{}
		line.Positions = make([]settle.Position, 0)
		line.Symbols = make([]*symbol.Item, 0)
		line.Index = i
		for j := 0; j < len(r.hitLines[i]); j++ {
			checkPoint := r.hitLines[i][j]
			symbolItem := reelSet[j][checkPoint]
			line.Positions = append(line.Positions, settle.Position{Col: j, Row: checkPoint})
			line.Symbols = append(line.Symbols, symbolItem)
		}
		// 計算這條是否是中獎線
		isWin, winSymbol := r.CheckWinLine(line)
		if !isWin {
			continue
		}
		line.Symbol = winSymbol
		line.Score = bet * value * winSymbol.Pow
		winLines = append(winLines, line)
	}
	return winLines, nil
}

func (r *repository) CheckWinLine(line *settle.Line) (bool, *symbol.Item) {
	// 判斷軸數樣式是否相符
	if len(line.Symbols) != len(r.hitLines[0]) {
		return false, nil
	}
	var checkSymbol *symbol.Item
	for _, symbolItem := range line.Symbols {
		if checkSymbol == nil {
			checkSymbol = symbolItem
			continue
		}
		if checkSymbol.IsWild {
			checkSymbol = symbolItem
			continue
		}
		if symbolItem.ID == checkSymbol.ID || symbolItem.IsWild {
			continue
		}
		return false, nil
	}
	return true, checkSymbol
}
