package reels

import (
	"fmt"
	gameCfg "game_server_slots_fortune_snake/config/game"
	symbolRepo "game_server_slots_fortune_snake/internal/repository/symbol"
	"testing"
)

// TestService_Generate 產生基本模式盤面
func TestService_Generate(*testing.T) {
	cgf := gameCfg.New()
	repo := symbolRepo.New(cgf.SymbolConfig())
	svc := New(repo)
	reelSet := svc.Generate([]int{3, 4, 3})
	for row := 0; row < len(reelSet); row++ {
		for col := 0; col < len(reelSet[row]); col++ {
			fmt.Println(reelSet[row][col])
		}
		fmt.Println("-------")
	}
}

// TestService_Generate_Free 產生免費模式盤面
func TestService_Generate_Free(*testing.T) {
	cgf := gameCfg.New()
	repo := symbolRepo.New(cgf.SymbolConfig())
	svc := NewFree(repo)
	reelSet := svc.Generate([]int{3, 4, 3})
	for row := 0; row < len(reelSet); row++ {
		for col := 0; col < len(reelSet[row]); col++ {
			fmt.Println(reelSet[row][col])
		}
		fmt.Println("-------")
	}
}
