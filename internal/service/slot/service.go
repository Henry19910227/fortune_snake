package slot

import (
	"game_server_slots_fortune_snake/internal/model/slot"
	symbolRepo "game_server_slots_fortune_snake/internal/repository/symbol"
)

type service struct {
	symbolRepo symbolRepo.Repository
}

func NewService(symbolRepo symbolRepo.Repository) Service {
	return &service{symbolRepo: symbolRepo}
}

func (s *service) Generate(layout []int) *slot.Result {
	result := &slot.Result{}
	result.Symbols = make([][]int, 0)
	for col := 0; col < len(layout); col++ {
		reel := make([]int, 0)
		for row := 0; row < layout[col]; row++ {
			reel = append(reel, s.symbolRepo.GetRandomSymbol().ID)
		}
		result.Symbols = append(result.Symbols, reel)
	}
	return result
}
