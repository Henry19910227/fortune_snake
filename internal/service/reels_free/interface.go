package reels_free

import (
	"game_server_slots_fortune_snake/internal/model/entity/symbol"
)

type Service interface {
	// Generate 生成免費模式盤面
	Generate(layout []int) [][][]*symbol.Item
}
