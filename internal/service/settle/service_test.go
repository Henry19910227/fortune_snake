package settle

import (
	"context"
	lineModel "game_server_slots_fortune_snake/internal/model/entity/line"
	"game_server_slots_fortune_snake/internal/model/entity/symbol"
	"game_server_slots_fortune_snake/internal/model/service/settle/check_win_line"
	"game_server_slots_fortune_snake/internal/model/service/settle/get_total_score"
	"game_server_slots_fortune_snake/internal/model/service/settle/get_win_lines"
	settleRepo "game_server_slots_fortune_snake/internal/repository/settle"
	symbolRepo "game_server_slots_fortune_snake/internal/repository/symbol"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSettleService_GetTotalScore_1(t *testing.T) {
	symRepo := symbolRepo.New()
	stlRepo := settleRepo.New()
	reels := [][]*symbol.Item{
		{symRepo.GetSymbol(0), symRepo.GetSymbol(0), symRepo.GetSymbol(0)},
		{symRepo.GetSymbol(0), symRepo.GetSymbol(0), symRepo.GetSymbol(0), symRepo.GetSymbol(0)},
		{symRepo.GetSymbol(0), symRepo.GetSymbol(0), symRepo.GetSymbol(0)},
	}

	svc := New(stlRepo)
	input := &get_total_score.Input{}
	input.Ctx = context.Background()
	input.Param = get_total_score.Param{
		Bet:   1,
		Value: 100,
		Reels: reels,
	}
	totalScore, _ := svc.GetTotalScore(input)
	assert.Equal(t, 20000, totalScore)
}

func TestSettleService_GetWinLines_1(t *testing.T) {
	symRepo := symbolRepo.New()
	stlRepo := settleRepo.New()
	reels := [][]*symbol.Item{
		{symRepo.GetSymbol(0), symRepo.GetSymbol(0), symRepo.GetSymbol(0)},
		{symRepo.GetSymbol(0), symRepo.GetSymbol(0), symRepo.GetSymbol(0), symRepo.GetSymbol(0)},
		{symRepo.GetSymbol(0), symRepo.GetSymbol(0), symRepo.GetSymbol(0)},
	}

	svc := New(stlRepo)
	input := &get_win_lines.Input{}
	input.Ctx = context.Background()
	input.Param = get_win_lines.Param{
		Bet:   1,
		Value: 100,
		Reels: reels,
	}
	winLines, _ := svc.GetWinLines(input)
	assert.Equal(t, 10, len(winLines))
	assert.Equal(t, 0, winLines[0].Symbol.ID)
	assert.Equal(t, 20000, winLines[0].Score)
}

// TestSettleRepo_CheckLine_1_1 測試第一個符號是百搭的情境
func TestSettleRepo_CheckLine_1_1(t *testing.T) {
	stlRepo := settleRepo.New()
	// 生成一條線
	line := &lineModel.Item{}
	line.Index = 0
	line.Symbols = []*symbol.Item{
		{ID: 0, IsWild: true},
		{ID: 1, IsWild: false},
		{ID: 1, IsWild: false},
	}

	svc := New(stlRepo)
	input := &check_win_line.Input{}
	input.Ctx = context.Background()
	input.Param = check_win_line.Param{
		Line: line,
	}
	winSymbol := svc.CheckWinLine(input)
	assert.NotNil(t, winSymbol)
	assert.Equal(t, 1, winSymbol.ID)

	line.Symbols = []*symbol.Item{
		{ID: 0, IsWild: true},
		{ID: 1, IsWild: false},
		{ID: 2, IsWild: false},
	}
	input.Param = check_win_line.Param{
		Line: line,
	}
	winSymbol = svc.CheckWinLine(input)
	assert.Nil(t, winSymbol)
}
