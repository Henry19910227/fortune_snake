package reels

import (
	"game_server_slots_fortune_snake/constants"
	"game_server_slots_fortune_snake/internal/model/entity/symbol"
	symbolRepo "game_server_slots_fortune_snake/internal/repository/symbol"
)

type Service interface {
	// Generate 生成一個盤面
	Generate(layout []int) [][][]*symbol.Item

	ToReels(results [][]int) [][]*symbol.Item

	ToResults(itemsList [][]*symbol.Item) [][]int

	GetMainSymbol(reels [][]*symbol.Item) *symbol.Item
}

func New(symbolRepo symbolRepo.Repository, spinMode int) Service {
	if spinMode == constants.SpinModeFree {
		return &serviceFree{symbolRepo: symbolRepo}
	}
	return &service{symbolRepo: symbolRepo}
}
