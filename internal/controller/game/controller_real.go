package game

import (
	"context"
	. "game_server_slots_fortune_snake/constants"
	betModel "game_server_slots_fortune_snake/internal/model/controller/game/bet"
	playerModel "game_server_slots_fortune_snake/internal/model/entity/player"
	"game_server_slots_fortune_snake/internal/server"
)

func (c *controller) BaseModeInReal(ctx *server.Context) {
	session := ctx.MustGet("session").(*playerModel.Session)

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
	winRate, err := c.settleService.GetRate(1, 1000, reels)
	if err != nil {
		ctx.SendError(err)
		return
	}

	// 計算中獎線
	lines, err := c.settleService.GetWinLines(1, 1000, reels)
	if err != nil {
		ctx.SendError(err)
		return
	}

	// 計算總分
	totalScore, err := c.settleService.GetTotalScore(1, 1000, reels)
	if err != nil {
		ctx.SendError(err)
		return
	}

	// 數據準備
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

	data := betModel.NewResponse(GameModeReal)
	data.GameResult = gameResult

	// 續存結果

	// 派彩

	// 返回結果
	ctx.Send(CodeSuccess, "success", data)
}

// StartFreeModeInReal 開始金蛇盤面流程
func (c *controller) StartFreeModeInReal(ctx *server.Context) {
	session := ctx.MustGet("session").(*playerModel.Session)
	grpcCtx := ctx.MustGet("ctx").(context.Context)

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

	// 回傳結果
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

	data := betModel.NewResponse(GameModeReal)
	data.GameResult = gameResult

	// 續存結果

	// 返回結果
	ctx.Send(CodeSuccess, "success", data)
}

func (c *controller) FreeModeInReal(ctx *server.Context) {
	grpcCtx := ctx.MustGet("ctx").(context.Context)
	session := ctx.MustGet("session").(*playerModel.Session)

	// 取出一筆金蛇盤面
	results, err := c.resultFreeService.GameMode(GameModeReal).PopFirstItem(grpcCtx, session.PlayerUsername)
	if err != nil {
		ctx.SendError(err)
		return
	}

	// 將盤面數據轉換為 reels
	reels := c.reelsFreeService.ToReels(results)

	// 回傳結果
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

	data := betModel.NewResponse(GameModeReal)
	data.GameResult = gameResult

	// 續存結果

	// 返回結果
	ctx.Send(CodeSuccess, "success", data)
}

// FinalFreeModeInReal 進入最後一個金蛇盤面流程
func (c *controller) FinalFreeModeInReal(ctx *server.Context) {
	grpcCtx := ctx.MustGet("ctx").(context.Context)
	session := ctx.MustGet("session").(*playerModel.Session)

	// 取出最後一筆金蛇盤面
	results, err := c.resultFreeService.GameMode(GameModeReal).PopFirstItem(grpcCtx, session.PlayerUsername)
	if err != nil {
		ctx.SendError(err)
		return
	}

	// 將盤面數據轉換為 reels
	reels := c.reelsFreeService.ToReels(results)

	// 計算賠率
	rate, err := c.settleService.GetRate(1, 1000, reels)
	if err != nil {
		ctx.SendError(err)
		return
	}

	// 計算中獎線
	lines, err := c.settleService.GetWinLines(1, 1000, reels)
	if err != nil {
		ctx.SendError(err)
		return
	}

	// 計算總分
	totalScore, err := c.settleService.GetTotalScore(1, 1000, reels)
	if err != nil {
		ctx.SendError(err)
		return
	}

	// 數據準備
	spinResult := &betModel.SpinResult{}
	spinResult.Score = totalScore
	spinResult.SetLines(lines)
	spinResult.Symbols = results
	spinResult.Times = c.settleService.GetTimes(reels)

	gameResult := betModel.NewGameResult(SpinModeFree)
	gameResult.RandSymbol = c.reelsFreeService.GetMainSymbol(reels).ID
	gameResult.SpinResult = spinResult
	gameResult.WinRate = int(rate)
	gameResult.TotalScore = totalScore
	gameResult.WinType = 0

	data := betModel.NewResponse(GameModeReal)
	data.GameResult = gameResult

	// 續存結果

	// 派彩同步

	// 返回結果
	ctx.Send(CodeSuccess, "success", data)
}
