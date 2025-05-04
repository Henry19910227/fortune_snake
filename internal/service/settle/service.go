package settle

import (
	"errors"
	lineModel "game_server_slots_fortune_snake/internal/model/line"
	"game_server_slots_fortune_snake/internal/model/settle/check_win_line"
	"game_server_slots_fortune_snake/internal/model/settle/get_win_lines"
	"game_server_slots_fortune_snake/internal/model/symbol"
	settleRepo "game_server_slots_fortune_snake/internal/repository/settle"
)

type service struct {
	settleRepo settleRepo.Repository
}

func New(settleRepo settleRepo.Repository) Service {
	return &service{settleRepo: settleRepo}
}

func (s *service) GetWinLines(input *get_win_lines.Input) ([]*lineModel.Item, error) {
	// 獲取參數
	reels := input.Param.Reels
	bet := input.Param.Bet
	value := input.Param.Value
	// 讀取中獎線資源
	hitLines := s.settleRepo.HitLines()
	// 判斷軸數是否相同
	if len(reels) != len(hitLines[0]) {
		return []*lineModel.Item{}, errors.New("盤面格式不符")
	}
	// 依照 hitLines 中的點位組成 Line
	winLines := make([]*lineModel.Item, 0)
	for i := 0; i < len(hitLines); i++ {
		winLine := &lineModel.Item{}
		winLine.Positions = make([]lineModel.Position, 0)
		winLine.Symbols = make([]*symbol.Item, 0)
		winLine.Index = i
		for j := 0; j < len(hitLines[i]); j++ {
			checkPoint := hitLines[i][j]
			symbolItem := reels[j][checkPoint]
			winLine.Positions = append(winLine.Positions, lineModel.Position{Col: j, Row: checkPoint})
			winLine.Symbols = append(winLine.Symbols, symbolItem)
		}
		// 計算這條是否是中獎線
		in := &check_win_line.Input{}
		in.Ctx = input.Ctx
		in.Param = check_win_line.Param{Line: winLine}
		winSymbol := s.CheckWinLine(in)
		if winSymbol == nil {
			continue
		}
		winLine.Symbol = winSymbol
		winLine.Score = bet * value * winSymbol.Pow
		winLines = append(winLines, winLine)
	}
	return winLines, nil
}

func (s *service) CheckWinLine(input *check_win_line.Input) *symbol.Item {
	// 獲取參數
	line := input.Param.Line
	// 讀取中獎線資源
	hitLines := s.settleRepo.HitLines()
	// 判斷軸數樣式是否相符
	if len(line.Symbols) != len(hitLines[0]) {
		return nil
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
		return nil
	}
	return checkSymbol
}
