package settle

import (
	"game_server_slots_fortune_snake/internal/model/line"
	"game_server_slots_fortune_snake/internal/model/settle/check_win_line"
	"game_server_slots_fortune_snake/internal/model/settle/get_win_lines"
	"game_server_slots_fortune_snake/internal/model/symbol"
)

type Service interface {
	// GetWinLines 將盤面數據傳入獲取中獎賠付線
	GetWinLines(input *get_win_lines.Input) ([]*line.Item, error)
	// CheckWinLine 判斷否是中獎線
	CheckWinLine(input *check_win_line.Input) *symbol.Item
}
