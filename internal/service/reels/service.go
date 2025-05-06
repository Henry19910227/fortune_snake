package reels

import (
	"game_server_slots_fortune_snake/internal/model/entity/symbol"
	symbolRepo "game_server_slots_fortune_snake/internal/repository/symbol"
)

type service struct {
	symbolRepo symbolRepo.Repository
}

func New(symbolRepo symbolRepo.Repository) Service {
	return &service{symbolRepo: symbolRepo}
}

func (s *service) Generate(layout []int) [][]*symbol.Item {
	reelSet := make([][]*symbol.Item, 0)
	for col := 0; col < len(layout); col++ {
		reel := make([]*symbol.Item, 0)
		for row := 0; row < layout[col]; row++ {
			reel = append(reel, s.symbolRepo.GetRandomSymbol())
		}
		reelSet = append(reelSet, reel)
	}
	return reelSet
}
