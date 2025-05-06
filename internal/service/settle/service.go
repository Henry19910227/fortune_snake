package settle

import (
	lineModel "game_server_slots_fortune_snake/internal/model/entity/line"
	"game_server_slots_fortune_snake/internal/model/entity/symbol"
	"game_server_slots_fortune_snake/internal/model/err"
	"game_server_slots_fortune_snake/internal/model/service/settle/check_win_line"
	"game_server_slots_fortune_snake/internal/model/service/settle/get_rate"
	"game_server_slots_fortune_snake/internal/model/service/settle/get_total_score"
	"game_server_slots_fortune_snake/internal/model/service/settle/get_win_lines"
	settleRepo "game_server_slots_fortune_snake/internal/repository/settle"
)

type service struct {
	settleRepo settleRepo.Repository
}

func New(settleRepo settleRepo.Repository) Service {
	return &service{settleRepo: settleRepo}
}

func (s *service) GetRate(input *get_rate.Input) (float64, error) {
	lines, err := s.getWinLines(get_win_lines.Param{
		Bet:   input.Param.Bet,
		Value: input.Param.Value,
		Reels: input.Param.Reels,
	})
	if err != nil {
		return 0, err
	}
	// 沒有任何符合的獎金線
	if len(lines) == 0 {
		return 0, nil
	}
	// 將百搭符號過濾
	symbolIdList := make([]int, 0)
	for _, line := range lines {
		if line.Symbol.IsWild {
			continue
		}
		symbolIdList = append(symbolIdList, line.Symbol.ID)
	}

	return 0, nil
}

func (s *service) GetTotalScore(input *get_total_score.Input) (int, error) {
	return s.getTotalScore(input.Param)
}

func (s *service) GetWinLines(input *get_win_lines.Input) ([]*lineModel.Item, error) {
	return s.getWinLines(input.Param)
}

func (s *service) CheckWinLine(input *check_win_line.Input) *symbol.Item {
	return s.checkWinLine(input.Param)
}

func (s *service) getTotalScore(param get_total_score.Param) (int, error) {
	lines, err := s.getWinLines(get_win_lines.Param{Bet: param.Bet, Value: param.Value, Reels: param.Reels})
	if err != nil {
		return 0, err
	}
	// 未中獎
	if len(lines) == 0 {
		return 0, nil
	}
	// 計算總金額
	var totalScore int
	for _, line := range lines {
		totalScore += line.Score
	}
	// 未達十條線
	if len(lines) < 10 {
		return totalScore, nil
	}
	// 將百搭符號中獎線過濾
	symbolIdList := make([]int, 0)
	for _, line := range lines {
		if line.Symbol.IsWild {
			continue
		}
		symbolIdList = append(symbolIdList, line.Symbol.ID)
	}
	if len(symbolIdList) == 0 {
		return totalScore * 10, nil
	}
	// 計算剩餘中獎線是否相同符號
	first := symbolIdList[0]
	for _, v := range symbolIdList[1:] {
		if v == first {
			continue
		}
		return totalScore, nil
	}
	// 所有中獎線相同符號獎金 * 10
	return totalScore * 10, nil
}

func (s *service) getWinLines(param get_win_lines.Param) ([]*lineModel.Item, error) {
	// 獲取參數
	reels := param.Reels
	bet := param.Bet
	value := param.Value
	// 讀取中獎線資源
	hitLines := s.settleRepo.HitLines()
	// 判斷軸數是否相同
	if len(reels) != len(hitLines[0]) {
		return []*lineModel.Item{}, err.New(400, "盤面格式不符", nil)
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
		winSymbol := s.checkWinLine(check_win_line.Param{Line: winLine})
		if winSymbol == nil {
			continue
		}
		winLine.Symbol = winSymbol
		winLine.Score = bet * value * winSymbol.Pow
		winLines = append(winLines, winLine)
	}
	return winLines, nil
}

func (s *service) checkWinLine(param check_win_line.Param) *symbol.Item {
	// 讀取中獎線資源
	hitLines := s.settleRepo.HitLines()
	// 判斷軸數樣式是否相符
	if len(param.Line.Symbols) != len(hitLines[0]) {
		return nil
	}
	var checkSymbol *symbol.Item
	for _, symbolItem := range param.Line.Symbols {
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

func (s *service) Transform(items [][]*symbol.Item) [][]int {
	symbols := make([][]int, 0)
	for i := 0; i < len(items); i++ {
		symbols = append(symbols, make([]int, 0))
	}
	for row := 0; row < len(items); row++ {
		for col := 0; col < len(items[row]); col++ {
			symbols[row] = append(symbols[row], items[row][col].ID)
		}
	}
	return symbols
}
