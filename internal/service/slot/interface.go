package slot

import (
	"game_server_slots_fortune_snake/internal/model/symbol"
)

type Service interface {
	// Generate 生成一個盤面
	Generate(layout []int) [][]*symbol.Item
}
