package settle

import (
	"game_server_slots_fortune_snake/internal/model/entity/line"
	"game_server_slots_fortune_snake/internal/model/entity/symbol"
)

type Service interface {
	// GetRate 獲取賠率
	GetRate(bet int, value int, reels [][]*symbol.Item) (rate float64, err error)
	// GetTotalScore 獲取總分
	GetTotalScore(bet int, value int, reels [][]*symbol.Item) (score int, err error)
	// GetWinLines 將盤面數據傳入獲取中獎賠付線
	GetWinLines(bet int, value int, reels [][]*symbol.Item) (lines []*line.Item, err error)
	// CheckWinLine 判斷否是中獎線
	CheckWinLine(line *line.Item) (symbol *symbol.Item, err error)
	// ToJson 將 symbol 物件轉換成 Json
	ToJson(items [][]*symbol.Item) (JsonString string, err error)
	// ListToJson 將 symbol items 轉換成 Json
	ListToJson(reelsList [][][]*symbol.Item) (reelsListString string, err error)
}
