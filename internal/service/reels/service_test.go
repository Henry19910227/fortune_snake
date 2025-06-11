package reels

import (
	"fmt"
	gameCfg "game_server_slots_fortune_snake/config/game"
	"game_server_slots_fortune_snake/constants"
	"game_server_slots_fortune_snake/internal/model/entity/symbol"
	symbolRepo "game_server_slots_fortune_snake/internal/repository/symbol"
	"testing"
)

// TestService_Generate 產生基本模式盤面
func TestService_Generate(*testing.T) {
	cgf := gameCfg.New()
	repo := symbolRepo.New(cgf.SymbolConfig())
	svc := New(repo, constants.SpinModeBase)
	reelsList := svc.Generate([]int{3, 4, 3})
	reels := reelsList[0]
	for row := 0; row < len(reels); row++ {
		for col := 0; col < len(reels[row]); col++ {
			fmt.Println(reels[row][col])
		}
		fmt.Println("-------")
	}
}

func TestService_ToReels(*testing.T) {
	cgf := gameCfg.New()
	repo := symbolRepo.New(cgf.SymbolConfig())
	svc := New(repo, constants.SpinModeBase)
	reelSet := svc.ToReels([][]int{{1, 1, 1}, {2, 3, 4, 5}, {2, 2, 2}})
	for row := 0; row < len(reelSet); row++ {
		for col := 0; col < len(reelSet[row]); col++ {
			fmt.Println(reelSet[row][col])
		}
		fmt.Println("-------")
	}
}

func TestService_ToResults(*testing.T) {
	cgf := gameCfg.New()
	repo := symbolRepo.New(cgf.SymbolConfig())
	svc := New(repo, constants.SpinModeBase)

	itemsList := make([][]*symbol.Item, 0)
	itemsList = append(itemsList,
		[]*symbol.Item{repo.GetSymbol(1), repo.GetSymbol(2), repo.GetSymbol(3)},
		[]*symbol.Item{repo.GetSymbol(0), repo.GetSymbol(0), repo.GetSymbol(0), repo.GetSymbol(0)},
		[]*symbol.Item{repo.GetSymbol(1), repo.GetSymbol(2), repo.GetSymbol(3)},
	)
	
	results := svc.ToResults(itemsList)
	for row := 0; row < len(results); row++ {
		for col := 0; col < len(results[row]); col++ {
			fmt.Println(results[row][col])
		}
		fmt.Println("-------")
	}
}
