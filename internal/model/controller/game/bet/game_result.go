package bet

import "encoding/json"

type GameResult struct {
	SpinResult *SpinResult `json:"spin_result,omitempty"`
	RandSymbol int         `json:"rand_symbol"` // 金蛇模式下的隨機符號
	SpinMode   int         `json:"spin_mode"`   // 旋轉模式(0:普通/1:金蛇)
	TotalScore int         `json:"total_score"` // 獲勝總金額
	WinRate    int         `json:"win_rate"`    // 賠率
	WinType    int         `json:"win_type"`    // 獲勝類型(0:Lose/1:Win/2:Big Win/3:Mage Win/4:Super Win)
}

func NewGameResult(spinMode int) *GameResult {
	gameResult := &GameResult{}
	gameResult.SpinMode = spinMode
	return gameResult
}

func (r *GameResult) Encode() []byte {
	b, _ := json.Marshal(r)
	return b
}
