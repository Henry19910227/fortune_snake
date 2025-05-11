package settle

import (
	"encoding/json"
	"game_server_slots_fortune_snake/constants"
	lineModel "game_server_slots_fortune_snake/internal/model/entity/line"
	"game_server_slots_fortune_snake/internal/model/entity/symbol"
	errMsg "game_server_slots_fortune_snake/internal/model/err"
	"game_server_slots_fortune_snake/internal/model/service/settle/check_win_line"
	"game_server_slots_fortune_snake/internal/model/service/settle/get_rate"
	"game_server_slots_fortune_snake/internal/model/service/settle/get_total_score"
	"game_server_slots_fortune_snake/internal/model/service/settle/get_win_lines"
	"game_server_slots_fortune_snake/internal/model/service/settle/list_to_json"
	"game_server_slots_fortune_snake/internal/model/service/settle/to_json"
	settleRepo "game_server_slots_fortune_snake/internal/repository/settle"
)

type service struct {
	settleRepo settleRepo.Repository
}

func New(settleRepo settleRepo.Repository) Service {
	return &service{settleRepo: settleRepo}
}

func (s *service) GetRate(input *get_rate.Input) (output *get_rate.Output, err error) {
	totalScore, err := s.getTotalScore(get_total_score.Param{
		Bet:   input.Param.Bet,
		Value: input.Param.Value,
		Reels: input.Param.Reels,
	})
	if err != nil {
		return nil, err
	}
	realBet := input.Param.Bet * input.Param.Value * 10
	rate := float64(totalScore) / float64(realBet)
	output = get_rate.NewOutput(rate)
	return output, nil
}

func (s *service) GetTotalScore(input *get_total_score.Input) (output *get_total_score.Output, err error) {
	score, err := s.getTotalScore(input.Param)
	if err != nil {
		return nil, err
	}
	output = get_total_score.NewOutput(score)
	return output, nil
}

func (s *service) GetWinLines(input *get_win_lines.Input) (output *get_win_lines.Output, err error) {
	lines, err := s.getWinLines(input.Param)
	if err != nil {
		return nil, err
	}
	output = get_win_lines.NewOutput(lines)
	return output, nil
}

func (s *service) CheckWinLine(input *check_win_line.Input) (output *check_win_line.Output, err error) {
	item := s.checkWinLine(input.Param)
	output = check_win_line.NewOutput(item)
	return output, nil
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
	// 計算第二軸百搭個數
	var wildCount int
	for _, symbolItem := range param.Reels[1] {
		if !symbolItem.IsWild {
			continue
		}
		wildCount++
	}
	// 第二軸百搭個數小於四 則分數 * 1
	if wildCount < 4 {
		return totalScore, nil
	}
	// 第二軸百搭個數為四 則分數 * 10
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
		return []*lineModel.Item{}, errMsg.New(constants.CodeBadRequest, "盤面格式不符", nil)
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
		if symbolItem.ID == 99 {
			return nil
		}
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

func (s *service) ToJson(input *to_json.Input) (output *to_json.Output, err error) {
	items := input.Param.Items
	symbols := make([][]int, 0)
	for i := 0; i < len(items); i++ {
		symbols = append(symbols, make([]int, 0))
	}
	for row := 0; row < len(items); row++ {
		for col := 0; col < len(items[row]); col++ {
			symbols[row] = append(symbols[row], items[row][col].ID)
		}
	}
	jsonString, err := json.Marshal(symbols)
	if err != nil {
		return nil, errMsg.New(constants.CodeBadRequest, err.Error(), err)
	}
	output = to_json.NewOutput(string(jsonString))
	return output, nil
}

func (s *service) ListToJson(input *list_to_json.Input) (output *list_to_json.Output, err error) {
	list := input.Param.List
	symbolsList := make([][][]int, 0)
	for _, items := range list {
		symbols := make([][]int, 0)
		for i := 0; i < len(items); i++ {
			symbols = append(symbols, make([]int, 0))
		}
		for row := 0; row < len(items); row++ {
			for col := 0; col < len(items[row]); col++ {
				symbols[row] = append(symbols[row], items[row][col].ID)
			}
		}
		symbolsList = append(symbolsList, symbols)
	}

	jsonString, err := json.Marshal(symbolsList)
	if err != nil {
		return nil, errMsg.New(constants.CodeBadRequest, err.Error(), err)
	}
	output = list_to_json.NewOutput(string(jsonString))
	return output, nil
}
