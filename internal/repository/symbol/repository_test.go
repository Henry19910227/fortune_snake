package symbol

import (
	"fmt"
	gameCfg "game_server_slots_fortune_snake/config/game"
	"testing"
)

func TestRepository_GetBaseSymbols(*testing.T) {
	// 創建配置檔組件
	cfg := gameCfg.New()
	// 載入 base symbol
	repo := New(cfg.BaseSymbolConfig())
	for _, symbol := range repo.GetSymbols() {
		fmt.Println(symbol)
	}
	fmt.Println(repo.GetTotalWeight())
}

func TestRepository_GetFreeSymbols(*testing.T) {
	// 創建配置檔組件
	cfg := gameCfg.New()
	// 載入 free symbol
	repo := New(cfg.FreeSymbolConfig())
	for _, symbol := range repo.GetSymbols() {
		fmt.Println(symbol)
	}
	fmt.Println(repo.GetTotalWeight())
}
