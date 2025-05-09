package reels

import (
	"game_server_slots_fortune_snake/internal/model/entity/symbol"
	symbolRepo "game_server_slots_fortune_snake/internal/repository/symbol"
	"math/rand"
)

type free struct {
	symbolRepo symbolRepo.Repository
}

func NewFree(symbolRepo symbolRepo.Repository) Free {
	return &free{symbolRepo: symbolRepo}
}

func (f *free) Generate(layout []int) [][]*symbol.Item {
	// 選擇一個隨機符號
	randomSymbol := f.symbolRepo.GetRandomSymbol()
	// 獲取一個空白符號
	spaceSymbol := f.symbolRepo.GetSymbol(99)

	reelSet := make([][]*symbol.Item, 0)
	for col := 0; col < len(layout); col++ {
		reel := make([]*symbol.Item, 0)
		for row := 0; row < layout[col]; row++ {
			num := rand.Intn(spaceSymbol.Weight + randomSymbol.Weight)
			if num < randomSymbol.Weight {
				reel = append(reel, randomSymbol)
				continue
			}
			reel = append(reel, spaceSymbol)
		}
		reelSet = append(reelSet, reel)
	}
	// 將第二軸改為全 wild + 空格符號，空格與 wild 符號出現機率各 50 %
	for i := 0; i < len(reelSet[1]); i++ {
		num := rand.Intn(100) + 1
		if num > 50 {
			reelSet[1][i] = spaceSymbol
			continue
		}
		reelSet[1][i] = f.symbolRepo.GetSymbol(0)
	}
	return reelSet
}

func (f *free) GenerateForFree(layout []int) [][][]*symbol.Item {
	//TODO implement me
	panic("implement me")
}
