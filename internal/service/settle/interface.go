package settle

import (
	"game_server_slots_fortune_snake/internal/model/entity/line"
	"game_server_slots_fortune_snake/internal/model/entity/symbol"
	"game_server_slots_fortune_snake/internal/model/service/settle/check_win_line"
	"game_server_slots_fortune_snake/internal/model/service/settle/get_rate"
	"game_server_slots_fortune_snake/internal/model/service/settle/get_total_score"
	"game_server_slots_fortune_snake/internal/model/service/settle/get_win_lines"
)

type Service interface {
	// GetRate 獲取賠率
	GetRate(input *get_rate.Input) (output *get_rate.Output, err error)
	// GetTotalScore 獲取總分
	GetTotalScore(input *get_total_score.Input) (int, error)
	// GetWinLines 將盤面數據傳入獲取中獎賠付線
	GetWinLines(input *get_win_lines.Input) ([]*line.Item, error)
	// CheckWinLine 判斷否是中獎線
	CheckWinLine(input *check_win_line.Input) *symbol.Item
	// ToJson 將 symbol 物件轉換成 Json
	ToJson(items [][]*symbol.Item) (string, error)
}
