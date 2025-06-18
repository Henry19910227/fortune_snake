package game

import (
	"context"
	betModel "game_server_slots_fortune_snake/internal/model/controller/game/bet"
	"game_server_slots_fortune_snake/internal/model/entity/line"
	playerModel "game_server_slots_fortune_snake/internal/model/entity/player"
	"game_server_slots_fortune_snake/internal/model/entity/symbol"
	"game_server_slots_fortune_snake/internal/model/service/game/save"
	"game_server_slots_fortune_snake/internal/server"
	"game_server_slots_fortune_snake/util"
)

// SavePlayerGameInfo 續存玩家遊戲投注參數與遊戲結果
func (c *controller) SavePlayerGameInfo(ctx *server.Context, gameResult *betModel.GameResult, param *betModel.Param) (err error) {
	session := ctx.MustGet("session").(*playerModel.Session)
	grpcCtx := ctx.MustGet("ctx").(context.Context)

	err = c.gameService.Save(save.Param{
		Ctx:        grpcCtx,
		GameMode:   session.Mode,
		PlayerID:   session.PlayerId,
		GameResult: gameResult,
		Bet:        util.PointerInt(param.Bet),
		Value:      util.PointerInt(param.Value),
		Bonus:      util.PointerBool(param.Bonus),
	})
	return err
}

// RestoreParam 恢復續存參數
func (c *controller) RestoreParam(ctx *server.Context, param *betModel.Param) (err error) {
	session := ctx.MustGet("session").(*playerModel.Session)
	grpcCtx := ctx.MustGet("ctx").(context.Context)
	output, err := c.gameService.GetParam(grpcCtx, session.Mode, session.PlayerId)
	if err != nil {
		return err
	}
	param.Bet = output.Bet
	param.Value = output.Value
	param.Bonus = output.Bonus
	return nil
}

// Settle 結算盤面
func (c *controller) Settle(param *betModel.Param, reels [][]*symbol.Item) (winRate float64, lines []*line.Item, totalScore int, err error) {
	// 計算賠率
	winRate, err = c.settleService.GetRate(param.Bet, param.Value, param.Multiplier, reels)
	if err != nil {
		return 0, []*line.Item{}, 0, err
	}

	// 計算中獎線
	lines, err = c.settleService.GetWinLines(param.Bet, param.Value, reels)
	if err != nil {
		return 0, []*line.Item{}, 0, err
	}

	// 計算總分
	totalScore, err = c.settleService.GetTotalScore(param.Bet, param.Value, reels)
	if err != nil {
		return 0, []*line.Item{}, 0, err
	}
	return winRate, lines, totalScore, nil
}
