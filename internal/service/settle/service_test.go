package settle

import (
	"context"
	lineModel "game_server_slots_fortune_snake/internal/model/entity/line"
	"game_server_slots_fortune_snake/internal/model/entity/symbol"
	"game_server_slots_fortune_snake/internal/model/service/settle/check_win_line"
	"game_server_slots_fortune_snake/internal/model/service/settle/get_rate"
	"game_server_slots_fortune_snake/internal/model/service/settle/get_total_score"
	"game_server_slots_fortune_snake/internal/model/service/settle/get_win_lines"
	settleRepo "game_server_slots_fortune_snake/internal/repository/settle"
	symbolRepo "game_server_slots_fortune_snake/internal/repository/symbol"
	"github.com/stretchr/testify/assert"
	"testing"
)

// [[6 3 5] [4 0 5 3] [4 5 6]]
// [[3 5 6] [4 4 6 6] [6 4 6]]
// [[6,4,0],[0,0,0,0],[1,4,0]]
// [[1,0,6],[0,0,0,0],[5,1,4]]
func TestSettleService_GetRate(t *testing.T) {
	symRepo := symbolRepo.New()
	stlRepo := settleRepo.New()
	reels := [][]*symbol.Item{
		{symRepo.GetSymbol(1), symRepo.GetSymbol(0), symRepo.GetSymbol(6)},
		{symRepo.GetSymbol(0), symRepo.GetSymbol(0), symRepo.GetSymbol(0), symRepo.GetSymbol(0)},
		{symRepo.GetSymbol(5), symRepo.GetSymbol(1), symRepo.GetSymbol(4)},
	}

	svc := New(stlRepo)
	input := &get_rate.Input{}
	input.Ctx = context.Background()
	input.Param = get_rate.Param{
		Bet:   1,
		Value: 1000,
		Reels: reels,
	}
	rate, _ := svc.GetRate(input)
	assert.Equal(t, float64(5000), rate)
}

// [[6 3 5] [4 0 5 3] [4 5 6]]
// [[3 5 6] [4 4 6 6] [6 4 6]]
// [[1,0,6],[0,0,0,0],[5,1,4]]
func TestSettleService_GetTotalScore_1(t *testing.T) {
	symRepo := symbolRepo.New()
	stlRepo := settleRepo.New()
	svc := New(stlRepo)

	// 測項 1：[[0 0 0] [0 0 0 0] [0 0 0]]
	reels := [][]*symbol.Item{
		{symRepo.GetSymbol(0), symRepo.GetSymbol(0), symRepo.GetSymbol(0)},
		{symRepo.GetSymbol(0), symRepo.GetSymbol(0), symRepo.GetSymbol(0), symRepo.GetSymbol(0)},
		{symRepo.GetSymbol(0), symRepo.GetSymbol(0), symRepo.GetSymbol(0)},
	}
	input := &get_total_score.Input{}
	input.Ctx = context.Background()
	input.Param = get_total_score.Param{
		Bet:   1,
		Value: 100,
		Reels: reels,
	}
	totalScore, _ := svc.GetTotalScore(input)
	assert.Equal(t, 5000000, totalScore)

	// 測項 2：[[1,0,6],[0,0,0,0],[5,1,4]]
	reels = [][]*symbol.Item{
		{symRepo.GetSymbol(1), symRepo.GetSymbol(0), symRepo.GetSymbol(6)},
		{symRepo.GetSymbol(0), symRepo.GetSymbol(0), symRepo.GetSymbol(0), symRepo.GetSymbol(0)},
		{symRepo.GetSymbol(5), symRepo.GetSymbol(1), symRepo.GetSymbol(4)},
	}
	input = &get_total_score.Input{}
	input.Ctx = context.Background()
	input.Param = get_total_score.Param{
		Bet:   1,
		Value: 100,
		Reels: reels,
	}
	totalScore, _ = svc.GetTotalScore(input)
	assert.Equal(t, 315000, totalScore)
}

// [[6,4,0],[0,0,0,0],[1,4,0]]
// [[1,0,6],[0,0,0,0],[5,1,4]]
func TestSettleService_GetWinLines_1(t *testing.T) {
	symRepo := symbolRepo.New()
	stlRepo := settleRepo.New()
	reels := [][]*symbol.Item{
		{symRepo.GetSymbol(1), symRepo.GetSymbol(0), symRepo.GetSymbol(6)},
		{symRepo.GetSymbol(0), symRepo.GetSymbol(0), symRepo.GetSymbol(0), symRepo.GetSymbol(0)},
		{symRepo.GetSymbol(5), symRepo.GetSymbol(1), symRepo.GetSymbol(4)},
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
	assert.Equal(t, 4, len(winLines))
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

func TestSettleRepo_Transform(t *testing.T) {
	symRepo := symbolRepo.New()
	stlRepo := settleRepo.New()
	reels := [][]*symbol.Item{
		{symRepo.GetSymbol(0), symRepo.GetSymbol(0), symRepo.GetSymbol(0)},
		{symRepo.GetSymbol(0), symRepo.GetSymbol(0), symRepo.GetSymbol(0), symRepo.GetSymbol(0)},
		{symRepo.GetSymbol(0), symRepo.GetSymbol(0), symRepo.GetSymbol(0)},
	}

	svc := New(stlRepo)
	jsonString, _ := svc.ToJson(reels)
	assert.Equal(t, "[[0,0,0],[0,0,0,0],[0,0,0]]", jsonString)
}
