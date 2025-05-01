package slot

import "game_server_slots_fortune_snake/internal/model/slot"

type Service interface {
	// Generate 生成一個盤面結果
	Generate(layout []int) *slot.Result
}
