package bet

type GameResult struct {
	SpinResult *SpinResult
	RandSymbol int // 金蛇模式下的隨機符號
	SpinMode   int // 旋轉模式(0:普通/1:金蛇)
	TotalScore int // 獲勝總金額
	WinRate    int // 賠率
	WinType    int // 獲勝類型(0:Lose/1:Win/2:Big Win/3:Mage Win/4:Super Win)
}

func NewGameResult(spinMode int) *GameResult {
	gameResult := &GameResult{}
	gameResult.SpinMode = spinMode
	return gameResult
}
