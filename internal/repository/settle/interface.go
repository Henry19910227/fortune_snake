package settle

import (
	"game_server_slots_fortune_snake/internal/model/settle"
	"game_server_slots_fortune_snake/internal/model/symbol"
)

type Repository interface {
	// GetWinLines 將盤面數據傳入獲取中獎賠付線
	GetWinLines(bet int, value int, reelSet [][]*symbol.Item) ([]*settle.Line, error)
	// CheckWinLine 判斷否是中獎線
	CheckWinLine(line *settle.Line) (bool, *symbol.Item)
	// GetLineScore 取得中獎金額
	GetLineScore(line *settle.Line) int
}
