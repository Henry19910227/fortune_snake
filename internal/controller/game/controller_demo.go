package game

import (
	"context"
	. "game_server_slots_fortune_snake/constants"
	betModel "game_server_slots_fortune_snake/internal/model/controller/game/bet"
	playerModel "game_server_slots_fortune_snake/internal/model/entity/player"
	"game_server_slots_fortune_snake/internal/server"
)

func (c *controller) BaseModeInDemo(ctx *server.Context) {
	session := ctx.MustGet("session").(*playerModel.Session)
	param := &betModel.Param{}
	if err := ctx.Bind(param); err != nil {
		ctx.SendError(err)
		return
	}

	// 計算總投注額
	totalBet := param.Bet * param.Value * 10

	// 用戶餘額減去總投注額
	session.Balance -= int64(totalBet)

	// 獲取賠率
	rate, err := c.weightService.RandomBaseWeightRate(float64(session.GameRtp))
	if err != nil {
		ctx.SendError(err)
		return
	}

	// 從內存隨機取得一筆普通盤面
	resultsList, err := c.resultLoader.Random(rate)
	if err != nil {
		ctx.SendError(err)
		return
	}

	// 取出唯一一筆普通盤面
	results := resultsList[0]

	// 將盤面數據轉換為 reels
	reels := c.reelsFreeService.ToReels(results)

	// 計算賠率
	winRate, err := c.settleService.GetRate(param.Bet, param.Value, reels)
	if err != nil {
		ctx.SendError(err)
		return
	}

	// 計算中獎線
	lines, err := c.settleService.GetWinLines(param.Bet, param.Value, reels)
	if err != nil {
		ctx.SendError(err)
		return
	}

	// 計算總分
	totalScore, err := c.settleService.GetTotalScore(param.Bet, param.Value, reels)
	if err != nil {
		ctx.SendError(err)
		return
	}

	// 派彩更新餘額
	session.Balance += int64(totalScore)
	if err := c.playerService.UpdateBalance(ctx, session.PlayerId, session.Balance); err != nil {
		ctx.SendError(err)
		return
	}

	// 返回結果
	spinResult := &betModel.SpinResult{}
	spinResult.Score = totalScore
	spinResult.SetLines(lines)
	spinResult.Symbols = c.reelsFreeService.ToResults(reels)
	spinResult.Times = 0

	gameResult := betModel.NewGameResult(SpinModeBase)
	gameResult.SpinResult = spinResult
	gameResult.WinRate = int(winRate)
	gameResult.TotalScore = totalScore
	gameResult.WinType = 0

	data := betModel.NewResponse(GameModeDemo)
	data.GameResult = gameResult
	data.Balance = int(session.Balance)

	ctx.Send(CodeSuccess, "success", data)
}

// StartFreeModeInDemo 開始金蛇盤面流程
func (c *controller) StartFreeModeInDemo(ctx *server.Context) {
	session := ctx.MustGet("session").(*playerModel.Session)
	grpcCtx := ctx.MustGet("ctx").(context.Context)
	param := &betModel.Param{}
	if err := ctx.Bind(param); err != nil {
		ctx.SendError(err)
		return
	}

	// 計算總投注額
	totalBet := param.Bet * param.Value * 10

	// 用戶餘額減去總投注額
	session.Balance -= int64(totalBet)

	// 獲取賠率
	rate, err := c.weightService.RandomFreeWeightRate(float64(session.GameRtp))
	if err != nil {
		ctx.SendError(err)
		return
	}

	// 從內存隨機取得多筆金蛇盤面
	resultsList, err := c.resultFreeLoader.Random(rate)
	if err != nil {
		ctx.SendError(err)
		return
	}

	// 取出第一個金蛇盤面
	results := resultsList[0]

	// 緩存剩餘金蛇盤面
	if err := c.resultFreeService.SaveItems(grpcCtx, session.PlayerUsername, resultsList[1:]); err != nil {
		ctx.SendError(err)
		return
	}

	// 將盤面數據轉換為 reels
	reels := c.reelsFreeService.ToReels(results)

	// 更新餘額
	if err := c.playerService.UpdateBalance(ctx, session.PlayerId, session.Balance); err != nil {
		ctx.SendError(err)
		return
	}

	// 返回結果
	spinResult := &betModel.SpinResult{}
	spinResult.Score = 0
	spinResult.Lines = []*betModel.Line{}
	spinResult.Symbols = results
	spinResult.Times = 0

	gameResult := betModel.NewGameResult(SpinModeFree)
	gameResult.RandSymbol = c.reelsFreeService.GetMainSymbol(reels).ID
	gameResult.SpinResult = spinResult
	gameResult.WinRate = 0
	gameResult.TotalScore = 0
	gameResult.WinType = 0

	data := betModel.NewResponse(GameModeDemo)
	data.GameResult = gameResult
	data.Balance = int(session.Balance)

	ctx.Send(CodeSuccess, "success", data)
}

// FreeModeInDemo 進行中的金蛇盤面流程
func (c *controller) FreeModeInDemo(ctx *server.Context) {
	session := ctx.MustGet("session").(*playerModel.Session)
	grpcCtx := ctx.MustGet("ctx").(context.Context)

	// 從緩存中取出一筆金蛇盤面
	results, err := c.resultFreeService.GameMode(GameModeDemo).PopFirstItem(grpcCtx, session.PlayerUsername)
	if err != nil {
		ctx.SendError(err)
		return
	}

	// 將盤面數據轉換為 reels
	reels := c.reelsFreeService.ToReels(results)

	// 返回結果
	spinResult := &betModel.SpinResult{}
	spinResult.Score = 0
	spinResult.Lines = []*betModel.Line{}
	spinResult.Symbols = results
	spinResult.Times = 0

	gameResult := betModel.NewGameResult(SpinModeFree)
	gameResult.RandSymbol = c.reelsFreeService.GetMainSymbol(reels).ID
	gameResult.SpinResult = spinResult
	gameResult.WinRate = 0
	gameResult.TotalScore = 0
	gameResult.WinType = 0

	data := betModel.NewResponse(GameModeDemo)
	data.GameResult = gameResult
	data.Balance = int(session.Balance)

	ctx.Send(CodeSuccess, "success", data)
}

// FinalFreeModeInDemo 進入最後一個金蛇盤面流程
func (c *controller) FinalFreeModeInDemo(ctx *server.Context) {
	grpcCtx := ctx.MustGet("ctx").(context.Context)
	session := ctx.MustGet("session").(*playerModel.Session)
	param := &betModel.Param{}
	if err := ctx.Bind(param); err != nil {
		ctx.SendError(err)
		return
	}

	// 從緩存中取出最後一筆金蛇盤面
	results, err := c.resultFreeService.GameMode(GameModeDemo).PopFirstItem(grpcCtx, session.PlayerUsername)
	if err != nil {
		ctx.SendError(err)
		return
	}

	// 將盤面數據轉換為 reels
	reels := c.reelsFreeService.ToReels(results)

	// 計算賠率
	rate, err := c.settleService.GetRate(param.Bet, param.Value, reels)
	if err != nil {
		ctx.SendError(err)
		return
	}

	// 計算中獎線
	lines, err := c.settleService.GetWinLines(param.Bet, param.Value, reels)
	if err != nil {
		ctx.SendError(err)
		return
	}

	// 計算總分
	totalScore, err := c.settleService.GetTotalScore(param.Bet, param.Value, reels)
	if err != nil {
		ctx.SendError(err)
		return
	}

	// 派彩更新餘額
	session.Balance += int64(totalScore)
	if err := c.playerService.UpdateBalance(ctx, session.PlayerId, session.Balance); err != nil {
		ctx.SendError(err)
		return
	}

	// 返回結果
	spinResult := &betModel.SpinResult{}
	spinResult.Score = totalScore
	spinResult.SetLines(lines)
	spinResult.Symbols = c.reelsFreeService.ToResults(reels)
	spinResult.Times = c.settleService.GetTimes(reels)

	gameResult := betModel.NewGameResult(SpinModeFree)
	gameResult.RandSymbol = c.reelsFreeService.GetMainSymbol(reels).ID
	gameResult.SpinResult = spinResult
	gameResult.WinRate = int(rate)
	gameResult.TotalScore = totalScore
	gameResult.WinType = 0

	data := betModel.NewResponse(GameModeDemo)
	data.GameResult = gameResult
	data.Balance = int(session.Balance)

	ctx.Send(CodeSuccess, "success", data)
}
