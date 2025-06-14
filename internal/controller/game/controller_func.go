package game

import (
	"context"
	betModel "game_server_slots_fortune_snake/internal/model/controller/game/bet"
	"game_server_slots_fortune_snake/internal/model/entity/line"
	playerModel "game_server_slots_fortune_snake/internal/model/entity/player"
	"game_server_slots_fortune_snake/internal/model/entity/symbol"
	"game_server_slots_fortune_snake/internal/server"
)

// SavePlayerGameInfo 續存玩家遊戲投注參數與遊戲結果
func (c *controller) SavePlayerGameInfo(ctx *server.Context, gameResult *betModel.GameResult, param *betModel.Param) (err error) {
	session := ctx.MustGet("session").(*playerModel.Session)
	grpcCtx := ctx.MustGet("ctx").(context.Context)
	// 續存 Game Result
	if err = c.gameService.GameMode(session.Mode).SaveGameResult(grpcCtx, session.PlayerId, gameResult); err != nil {
		return err
	}
	// 續存 Bet
	if err = c.gameService.GameMode(session.Mode).SaveBet(grpcCtx, session.PlayerId, param.Bet); err != nil {
		return err
	}
	// 續存 Value
	if err = c.gameService.GameMode(session.Mode).SaveValue(grpcCtx, session.PlayerId, param.Value); err != nil {
		return err
	}
	return nil
}

// Settle 結算盤面
func (c *controller) Settle(param *betModel.Param, reels [][]*symbol.Item) (winRate float64, lines []*line.Item, totalScore int, err error) {
	// 計算賠率
	winRate, err = c.settleService.GetRate(param.Bet, param.Value, reels)
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
