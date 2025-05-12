package bet

import lineModel "game_server_slots_fortune_snake/internal/model/entity/line"

type Response struct {
	Balance     int // 餘額
	BalanceDemo int // 試玩餘額
	GameResult  GameResult
}

type GameResult struct {
	SpinResult SpinResult
	SpinMode   int // 旋轉模式(0:普通/1:金蛇)
	TotalScore int // 獲勝總金額
	WinType    int // 獲勝類型(0:Lose/1:Win/2:Big Win/3:Mage Win/4:Super Win)
}

type SpinResult struct {
	Symbols [][]int
	Lines   []*lineModel.Item
	Times   int
}

type Line struct {
	Index     int        // 賠付線的位置
	Positions []Position // 中獎符號位置
	Symbols   []int      // 線的組成符號
	Symbol    int        // 中獎符號
	Score     int        // 中獎金額
}

type Position struct {
	Col int
	Row int
}
