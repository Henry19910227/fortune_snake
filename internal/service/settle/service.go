package settle

import (
	"encoding/json"
	"game_server_slots_fortune_snake/constants"
	lineModel "game_server_slots_fortune_snake/internal/model/entity/line"
	"game_server_slots_fortune_snake/internal/model/entity/symbol"
	errMsg "game_server_slots_fortune_snake/internal/model/err"
	settleRepo "game_server_slots_fortune_snake/internal/repository/settle"
)

type service struct {
	settleRepo settleRepo.Repository
}

func New(settleRepo settleRepo.Repository) Service {
	return &service{settleRepo: settleRepo}
}

func (s *service) GetRate(bet int, value int, reels [][]*symbol.Item) (rate float64, err error) {
	totalScore, err := s.GetTotalScore(bet, value, reels)
	if err != nil {
		return 0, err
	}
	realBet := bet * value * 10
	rate = float64(totalScore) / float64(realBet)
	return rate, nil
}

func (s *service) GetTotalScore(bet int, value int, reels [][]*symbol.Item) (score int, err error) {
	lines, err := s.getWinLines(bet, value, reels)
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
	// 計算是否乘以十倍
	times := s.GetTimes(reels)

	return totalScore * times, nil
}

func (s *service) GetTimes(reels [][]*symbol.Item) int {
	// 計算第二軸百搭個數
	var wildCount int
	for _, symbolItem := range reels[1] {
		if !symbolItem.IsWild {
			continue
		}
		wildCount++
	}
	// 第二軸百搭個數小於四 則分數 * 1
	if wildCount < 4 {
		return 1
	}
	// 第二軸百搭個數為四 則分數 * 10
	return 10
}

func (s *service) GetWinLines(bet int, value int, reels [][]*symbol.Item) (lines []*lineModel.Item, err error) {
	lines, err = s.getWinLines(bet, value, reels)
	if err != nil {
		return []*lineModel.Item{}, err
	}
	return lines, nil
}

func (s *service) CheckWinLine(line *lineModel.Item) (symbol *symbol.Item, err error) {
	item := s.checkWinLine(line)
	return item, nil
}

func (s *service) getWinLines(bet int, value int, reels [][]*symbol.Item) ([]*lineModel.Item, error) {
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
			winLine.Positions = append(winLine.Positions, lineModel.Position{Col: j, Row: checkPoint, Symbol: symbolItem})
			winLine.Symbols = append(winLine.Symbols, symbolItem)
		}
		// 計算這條是否是中獎線
		winSymbol := s.checkWinLine(winLine)
		if winSymbol == nil {
			continue
		}
		winLine.Symbol = winSymbol
		winLine.Score = bet * value * winSymbol.Pow
		winLines = append(winLines, winLine)
	}
	return winLines, nil
}

func (s *service) checkWinLine(line *lineModel.Item) *symbol.Item {
	// 讀取中獎線資源
	hitLines := s.settleRepo.HitLines()
	// 判斷軸數樣式是否相符
	if len(line.Symbols) != len(hitLines[0]) {
		return nil
	}
	var checkSymbol *symbol.Item
	for _, symbolItem := range line.Symbols {
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

func (s *service) ToJson(items [][]*symbol.Item) (JsonString string, err error) {
	symbols := make([][]int, 0)
	for i := 0; i < len(items); i++ {
		symbols = append(symbols, make([]int, 0))
	}
	for row := 0; row < len(items); row++ {
		for col := 0; col < len(items[row]); col++ {
			symbols[row] = append(symbols[row], items[row][col].ID)
		}
	}
	result, err := json.Marshal(symbols)
	if err != nil {
		return "", errMsg.New(constants.CodeBadRequest, err.Error(), err)
	}
	JsonString = string(result)
	return JsonString, nil
}

func (s *service) ListToJson(reelsList [][][]*symbol.Item) (reelsListString string, err error) {
	symbolsList := make([][][]int, 0)
	for _, items := range reelsList {
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
		return "", errMsg.New(constants.CodeBadRequest, err.Error(), err)
	}
	reelsListString = string(jsonString)
	return reelsListString, nil
}
