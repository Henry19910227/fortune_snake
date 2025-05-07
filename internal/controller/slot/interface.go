package slot

import (
	"game_server_slots_fortune_snake/internal/model/entity/symbol"
)

type Controller interface {
	// Generate 生成一個盤面
	Generate(layout []int) [][]*symbol.Item
}
