package reels_free

import (
	"fmt"
	gameCfg "game_server_slots_fortune_snake/config/game"
	"game_server_slots_fortune_snake/internal/model/entity/symbol"
	symbolRepo "game_server_slots_fortune_snake/internal/repository/symbol"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestService_Generate(t *testing.T) {
	cgf := gameCfg.New()
	repo := symbolRepo.New(cgf.SymbolConfig())
	svc := New(repo)
	reelsList := svc.Generate([]int{3, 4, 3})
	fmt.Println(reelsList)

}

func TestService_isEqual(t *testing.T) {
	cgf := gameCfg.New()
	repo := symbolRepo.New(cgf.SymbolConfig())
	svc := &service{symbolRepo: repo}

	reels1 := [][]*symbol.Item{
		{repo.GetSymbol(0), repo.GetSymbol(0), repo.GetSymbol(0)},
		{repo.GetSymbol(0), repo.GetSymbol(0), repo.GetSymbol(0), repo.GetSymbol(0)},
		{repo.GetSymbol(0), repo.GetSymbol(0), repo.GetSymbol(0)},
	}

	reels2 := [][]*symbol.Item{
		{repo.GetSymbol(0), repo.GetSymbol(0), repo.GetSymbol(0)},
		{repo.GetSymbol(0), repo.GetSymbol(1), repo.GetSymbol(0), repo.GetSymbol(0)},
		{repo.GetSymbol(0), repo.GetSymbol(0), repo.GetSymbol(0)},
	}
	assert.Equal(t, false, svc.isEqual(reels1, reels2))
}
