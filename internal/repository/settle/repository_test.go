package settle

import (
	"game_server_slots_fortune_snake/internal/model/settle"
	"game_server_slots_fortune_snake/internal/model/symbol"
	symbolRepo "game_server_slots_fortune_snake/internal/repository/symbol"
	"github.com/stretchr/testify/assert"
	"testing"
)

// TestSettleRepo_GetWinLines_1 測試十條百搭的盤面數據
func TestSettleRepo_GetWinLines_1(t *testing.T) {
	symRepo := symbolRepo.New()
	reelSet := [][]*symbol.Item{
		{symRepo.GetSymbol(0), symRepo.GetSymbol(0), symRepo.GetSymbol(0)},
		{symRepo.GetSymbol(0), symRepo.GetSymbol(0), symRepo.GetSymbol(0), symRepo.GetSymbol(0)},
		{symRepo.GetSymbol(0), symRepo.GetSymbol(0), symRepo.GetSymbol(0)},
	}
	repo := New()
	winLines, _ := repo.GetWinLines(1, 100, reelSet)
	assert.Equal(t, 10, len(winLines))
	assert.Equal(t, 0, winLines[0].Symbol.ID)
	assert.Equal(t, 20000, winLines[0].Score)
}

// TestSettleRepo_GetWinLines_2 測試一條中獎線的盤面數據
func TestSettleRepo_GetWinLines_2(t *testing.T) {
	symRepo := symbolRepo.New()
	repo := New()
	reelSet := [][]*symbol.Item{
		{symRepo.GetSymbol(1), symRepo.GetSymbol(2), symRepo.GetSymbol(3)},
		{symRepo.GetSymbol(0), symRepo.GetSymbol(4), symRepo.GetSymbol(5), symRepo.GetSymbol(6)},
		{symRepo.GetSymbol(1), symRepo.GetSymbol(2), symRepo.GetSymbol(2)},
	}
	winLines, _ := repo.GetWinLines(1, 100, reelSet)
	assert.Equal(t, 1, len(winLines))
	assert.Equal(t, 1, winLines[0].Symbol.ID)
	assert.Equal(t, 0, winLines[0].Index)
	assert.Equal(t, 10000, winLines[0].Score)

	reelSet = [][]*symbol.Item{
		{symRepo.GetSymbol(0), symRepo.GetSymbol(5), symRepo.GetSymbol(3)},
		{symRepo.GetSymbol(4), symRepo.GetSymbol(2), symRepo.GetSymbol(5), symRepo.GetSymbol(6)},
		{symRepo.GetSymbol(3), symRepo.GetSymbol(2), symRepo.GetSymbol(4)},
	}
	winLines, _ = repo.GetWinLines(1, 100, reelSet)
	assert.Equal(t, 1, len(winLines))
	assert.Equal(t, 2, winLines[0].Symbol.ID)
	assert.Equal(t, 2, winLines[0].Index)
	assert.Equal(t, 5000, winLines[0].Score)
}

// TestSettleRepo_CheckLine_1 測試三個百搭的情境
func TestSettleRepo_CheckLine_1(t *testing.T) {
	line := &settle.Line{}
	line.Index = 0
	line.Symbols = []*symbol.Item{
		{ID: 0, IsWild: true},
		{ID: 0, IsWild: true},
		{ID: 0, IsWild: true},
	}
	repo := New()
	isWin, winSymbol := repo.CheckWinLine(line)
	assert.Equal(t, true, isWin)
	assert.Equal(t, 0, winSymbol.ID)
}

// TestSettleRepo_CheckLine_1_1 測試第一個符號是百搭的情境
func TestSettleRepo_CheckLine_1_1(t *testing.T) {
	line := &settle.Line{}
	line.Index = 0
	line.Symbols = []*symbol.Item{
		{ID: 0, IsWild: true},
		{ID: 1, IsWild: false},
		{ID: 1, IsWild: false},
	}
	repo := New()
	isWin, winSymbol := repo.CheckWinLine(line)
	assert.Equal(t, true, isWin)
	assert.Equal(t, 1, winSymbol.ID)

	line.Symbols = []*symbol.Item{
		{ID: 0, IsWild: true},
		{ID: 1, IsWild: false},
		{ID: 2, IsWild: false},
	}
	isWin, winSymbol = repo.CheckWinLine(line)
	assert.Equal(t, false, isWin)
	assert.Nil(t, winSymbol)
}

// TestSettleRepo_CheckLine_2 測試兩個相同符號搭配一個百搭的情境
func TestSettleRepo_CheckLine_2(t *testing.T) {
	line := &settle.Line{}
	line.Index = 0
	line.Symbols = []*symbol.Item{
		{ID: 1, IsWild: false},
		{ID: 0, IsWild: true},
		{ID: 1, IsWild: false},
	}
	repo := New()
	isWin, winSymbol := repo.CheckWinLine(line)
	assert.Equal(t, true, isWin)
	assert.Equal(t, 1, winSymbol.ID)
}

// TestSettleRepo_CheckLine_3 測試三個相同符號的情境
func TestSettleRepo_CheckLine_3(t *testing.T) {
	line := &settle.Line{}
	line.Index = 0
	line.Symbols = []*symbol.Item{
		{ID: 1, IsWild: false},
		{ID: 1, IsWild: false},
		{ID: 1, IsWild: false},
	}
	repo := New()
	isWin, winSymbol := repo.CheckWinLine(line)
	assert.Equal(t, true, isWin)
	assert.Equal(t, 1, winSymbol.ID)
}

// TestSettleRepo_CheckLine_4 測試兩個相同與一個不同符號的情境
func TestSettleRepo_CheckLine_4(t *testing.T) {
	line := &settle.Line{}
	line.Index = 0
	line.Symbols = []*symbol.Item{
		{ID: 1, IsWild: false},
		{ID: 1, IsWild: false},
		{ID: 2, IsWild: false},
	}
	repo := New()
	isWin, winSymbol := repo.CheckWinLine(line)
	assert.Equal(t, false, isWin)
	assert.Nil(t, winSymbol)
}

// TestSettleRepo_CheckLine_5 測試格式錯誤的情境
func TestSettleRepo_CheckLine_5(t *testing.T) {
	line := &settle.Line{}
	line.Index = 0
	line.Symbols = []*symbol.Item{
		{ID: 1, IsWild: false},
		{ID: 1, IsWild: false},
		{ID: 1, IsWild: false},
		{ID: 1, IsWild: false},
	}
	repo := New()
	isWin, winSymbol := repo.CheckWinLine(line)
	assert.Equal(t, false, isWin)
	assert.Nil(t, winSymbol)
}
