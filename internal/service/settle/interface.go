package settle

import (
	"game_server_slots_fortune_snake/internal/model/entity/line"
	"game_server_slots_fortune_snake/internal/model/entity/symbol"
	"game_server_slots_fortune_snake/internal/model/service/settle/check_win_line"
	"game_server_slots_fortune_snake/internal/model/service/settle/get_win_lines"
)

type Service interface {
	GetRate()
	// GetWinLines 將盤面數據傳入獲取中獎賠付線
	GetWinLines(input *get_win_lines.Input) ([]*line.Item, error)
	// CheckWinLine 判斷否是中獎線
	CheckWinLine(input *check_win_line.Input) *symbol.Item
}
