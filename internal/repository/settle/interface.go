package settle

import (
	"game_server_slots_fortune_snake/internal/model/settle"
	"game_server_slots_fortune_snake/internal/model/symbol"
)

type Repository interface {
	// PayLines 將盤面數據傳入獲取中獎賠付線
	PayLines(reelSet [][]*symbol.Item) []*settle.PayLine
}
