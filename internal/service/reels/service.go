package reels

import (
	"game_server_slots_fortune_snake/internal/model/entity/symbol"
	symbolRepo "game_server_slots_fortune_snake/internal/repository/symbol"
	"math/rand"
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
	// 將第二軸全 wild 的機率提升為 50%，比較容易能產生出高分數盤面結果
	num := rand.Intn(100) + 1
	if num > 50 {
		for i := 0; i < len(reelSet[1]); i++ {
			reelSet[1][i] = s.symbolRepo.GetSymbol(0)
		}
	}
	return reelSet
}

func (s *service) ToReels(origin [][]int) [][]*symbol.Item {
	reelSet := make([][]*symbol.Item, 0)
	for col := 0; col < len(origin); col++ {
		reel := make([]*symbol.Item, 0)
		for row := 0; row < len(origin[col]); row++ {
			reel = append(reel, s.symbolRepo.GetSymbol(origin[col][row]))
		}
		reelSet = append(reelSet, reel)
	}
	return reelSet
}
