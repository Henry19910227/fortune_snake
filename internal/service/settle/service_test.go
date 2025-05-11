package settle

import (
	gameCfg "game_server_slots_fortune_snake/config/game"
	lineModel "game_server_slots_fortune_snake/internal/model/entity/line"
	"game_server_slots_fortune_snake/internal/model/entity/symbol"
	"game_server_slots_fortune_snake/internal/model/service/settle/check_win_line"
	"game_server_slots_fortune_snake/internal/model/service/settle/get_rate"
	"game_server_slots_fortune_snake/internal/model/service/settle/get_total_score"
	"game_server_slots_fortune_snake/internal/model/service/settle/get_win_lines"
	"game_server_slots_fortune_snake/internal/model/service/settle/to_json"
	settleRepo "game_server_slots_fortune_snake/internal/repository/settle"
	symbolRepo "game_server_slots_fortune_snake/internal/repository/symbol"
	"game_server_slots_fortune_snake/internal/server"
	"github.com/stretchr/testify/assert"
	"testing"
)

// [[6 3 5] [4 0 5 3] [4 5 6]]
// [[3 5 6] [4 4 6 6] [6 4 6]]
// [[6,4,0],[0,0,0,0],[1,4,0]]
// [[1,0,6],[0,0,0,0],[5,1,4]]
func TestSettleService_GetRate(t *testing.T) {
	cfg := gameCfg.New()
	symRepo := symbolRepo.New(cfg.SymbolConfig())
	stlRepo := settleRepo.New()
	reels := [][]*symbol.Item{
		{symRepo.GetSymbol(1), symRepo.GetSymbol(0), symRepo.GetSymbol(6)},
		{symRepo.GetSymbol(0), symRepo.GetSymbol(0), symRepo.GetSymbol(0), symRepo.GetSymbol(0)},
		{symRepo.GetSymbol(5), symRepo.GetSymbol(1), symRepo.GetSymbol(4)},
	}

	svc := New(stlRepo)
	input := get_rate.NewInput(get_rate.Param{Bet: 1, Value: 1000, Reels: reels})
	input.Ctx = &server.Context{}
	output, _ := svc.GetRate(input)
	assert.Equal(t, float64(5000), output.GetRate())
}

// [[6 3 5] [4 0 5 3] [4 5 6]]
// [[3 5 6] [4 4 6 6] [6 4 6]]
// [[1,0,6],[0,0,0,0],[5,1,4]]
// [[1,1,1],[99,0,0,0],[1,1,1]]
// [[1,1,1],[0,0,0,0],[1,1,1]]
func TestSettleService_GetTotalScore_1(t *testing.T) {
	cfg := gameCfg.New()
	symRepo := symbolRepo.New(cfg.SymbolConfig())
	stlRepo := settleRepo.New()
	svc := New(stlRepo)

	// 測項 1：[[0 0 0] [0 0 0 0] [0 0 0]]
	//reels := [][]*symbol.Item{
	//	{symRepo.GetSymbol(0), symRepo.GetSymbol(0), symRepo.GetSymbol(0)},
	//	{symRepo.GetSymbol(0), symRepo.GetSymbol(0), symRepo.GetSymbol(0), symRepo.GetSymbol(0)},
	//	{symRepo.GetSymbol(0), symRepo.GetSymbol(0), symRepo.GetSymbol(0)},
	//}
	//input := get_total_score.NewInput(get_total_score.Param{
	//	Bet:   1,
	//	Value: 100,
	//	Reels: reels,
	//})
	//input.Ctx = &server.Context{}
	//output, _ := svc.GetTotalScore(input)
	//assert.Equal(t, 5000000, output.GetScore())

	// 測項 2：[[1,0,6],[0,0,0,0],[5,1,4]]
	//reels = [][]*symbol.Item{
	//	{symRepo.GetSymbol(1), symRepo.GetSymbol(0), symRepo.GetSymbol(6)},
	//	{symRepo.GetSymbol(0), symRepo.GetSymbol(0), symRepo.GetSymbol(0), symRepo.GetSymbol(0)},
	//	{symRepo.GetSymbol(5), symRepo.GetSymbol(1), symRepo.GetSymbol(4)},
	//}
	//input = get_total_score.NewInput(get_total_score.Param{
	//	Bet:   1,
	//	Value: 100,
	//	Reels: reels,
	//})
	//input.Ctx = &server.Context{}
	//output, _ = svc.GetTotalScore(input)
	//assert.Equal(t, 315000, output.GetScore())

	// 測項 3：[[1,1,1],[99,0,0,0],[1,1,1]]
	reels := [][]*symbol.Item{
		{symRepo.GetSymbol(1), symRepo.GetSymbol(1), symRepo.GetSymbol(1)},
		{symRepo.GetSymbol(99), symRepo.GetSymbol(0), symRepo.GetSymbol(0), symRepo.GetSymbol(0)},
		{symRepo.GetSymbol(1), symRepo.GetSymbol(1), symRepo.GetSymbol(1)},
	}
	input := get_total_score.NewInput(get_total_score.Param{
		Bet:   1,
		Value: 100,
		Reels: reels,
	})
	input.Ctx = &server.Context{}
	output, _ := svc.GetTotalScore(input)
	assert.Equal(t, 315000, output.GetScore())
}

// [[6,4,0],[0,0,0,0],[1,4,0]]
// [[1,0,6],[0,0,0,0],[5,1,4]]
func TestSettleService_GetWinLines_1(t *testing.T) {
	cfg := gameCfg.New()
	symRepo := symbolRepo.New(cfg.SymbolConfig())
	stlRepo := settleRepo.New()
	reels := [][]*symbol.Item{
		{symRepo.GetSymbol(1), symRepo.GetSymbol(1), symRepo.GetSymbol(1)},
		{symRepo.GetSymbol(0), symRepo.GetSymbol(0), symRepo.GetSymbol(0), symRepo.GetSymbol(0)},
		{symRepo.GetSymbol(5), symRepo.GetSymbol(1), symRepo.GetSymbol(4)},
	}

	svc := New(stlRepo)
	input := get_win_lines.NewInput(get_win_lines.Param{
		Bet:   1,
		Value: 100,
		Reels: reels,
	})
	input.Ctx = &server.Context{}
	output, _ := svc.GetWinLines(input)
	assert.Equal(t, 4, len(output.GetLines()))
	assert.Equal(t, 0, output.GetLines()[0].Symbol.ID)
	assert.Equal(t, 20000, output.GetLines()[0].Score)
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
	input := check_win_line.NewInput(line)
	input.Ctx = &server.Context{}
	output, _ := svc.CheckWinLine(input)
	assert.NotNil(t, output.GetSymbol())
	assert.Equal(t, 1, output.GetSymbol().ID)

	line.Symbols = []*symbol.Item{
		{ID: 0, IsWild: true},
		{ID: 1, IsWild: false},
		{ID: 2, IsWild: false},
	}
	input.Param = check_win_line.Param{
		Line: line,
	}
	output, _ = svc.CheckWinLine(input)
	assert.Nil(t, output.GetSymbol())
}

func TestSettleRepo_Transform(t *testing.T) {
	cfg := gameCfg.New()
	symRepo := symbolRepo.New(cfg.SymbolConfig())
	stlRepo := settleRepo.New()
	reels := [][]*symbol.Item{
		{symRepo.GetSymbol(0), symRepo.GetSymbol(0), symRepo.GetSymbol(0)},
		{symRepo.GetSymbol(0), symRepo.GetSymbol(0), symRepo.GetSymbol(0), symRepo.GetSymbol(0)},
		{symRepo.GetSymbol(0), symRepo.GetSymbol(0), symRepo.GetSymbol(0)},
	}

	svc := New(stlRepo)
	output, _ := svc.ToJson(to_json.NewInput(reels))
	assert.Equal(t, "[[0,0,0],[0,0,0,0],[0,0,0]]", output.Data.JsonString)
}
