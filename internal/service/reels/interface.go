package reels

import (
	"game_server_slots_fortune_snake/internal/model/entity/symbol"
)

type Service interface {
	// Generate 生成一個盤面
	Generate(layout []int) [][]*symbol.Item
}

type Free interface {
	Service
	// GenerateForFree 生成免費模式盤面
	GenerateForFree(layout []int) [][][]*symbol.Item
}
