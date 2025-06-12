package reels

import (
	"game_server_slots_fortune_snake/internal/model/entity/symbol"
	symbolRepo "game_server_slots_fortune_snake/internal/repository/symbol"
	"math/rand"
)

type service struct {
	symbolRepo symbolRepo.Repository
}

func (s *service) Generate(layout []int) [][][]*symbol.Item {
	reels := make([][]*symbol.Item, 0)
	for col := 0; col < len(layout); col++ {
		reel := make([]*symbol.Item, 0)
		for row := 0; row < layout[col]; row++ {
			reel = append(reel, s.symbolRepo.GetRandomSymbol())
		}
		reels = append(reels, reel)
	}
	// 將第二軸全 wild 的機率提升為 50%，比較容易能產生出高分數盤面結果
	num := rand.Intn(100) + 1
	if num > 50 {
		for i := 0; i < len(reels[1]); i++ {
			reels[1][i] = s.symbolRepo.GetSymbol(0)
		}
	}
	return [][][]*symbol.Item{reels}
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

func (s *service) ToResults(itemsList [][]*symbol.Item) [][]int {
	results := make([][]int, 0, len(itemsList))
	for _, items := range itemsList {
		row := make([]int, 0, len(items))
		for _, item := range items {
			row = append(row, item.ID)
		}
		results = append(results, row)
	}
	return results
}

func (s *service) GetMainSymbol(reels [][]*symbol.Item) *symbol.Item {
	return s.symbolRepo.GetSymbol(0)
}
