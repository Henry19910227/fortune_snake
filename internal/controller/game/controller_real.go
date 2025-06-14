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

	// 結算盤面
	winRate, lines, totalScore, err := c.Settle(param, reels)
	if err != nil {
		ctx.SendError(err)
		return
	}

	// 派彩更新餘額
	session.Balance += int64(totalScore)
	if err = c.playerService.UpdateBalance(grpcCtx, session.PlayerId, session.Balance); err != nil {
		ctx.SendError(err)
		return
	}

	// 生成響應數據
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
	data.Balance = int(session.Balance)

	// 續存結果
	if err = c.SavePlayerGameInfo(ctx, gameResult, param); err != nil {
		ctx.SendError(err)
		return
	}

	ctx.Send(CodeSuccess, "success", data)
}

// StartFreeModeInReal 開始金蛇盤面流程
func (c *controller) StartFreeModeInReal(ctx *server.Context) {
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
	if err := c.resultFreeService.SaveItems(grpcCtx, session.PlayerId, resultsList[1:]); err != nil {
		ctx.SendError(err)
		return
	}

	// 將盤面數據轉換為 reels
	reels := c.reelsFreeService.ToReels(results)

	// 更新餘額
	if err := c.playerService.UpdateBalance(grpcCtx, session.PlayerId, session.Balance); err != nil {
		ctx.SendError(err)
		return
	}

	// 生成響應數據
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
	data.Balance = int(session.Balance)

	// 續存結果
	if err = c.SavePlayerGameInfo(ctx, gameResult, param); err != nil {
		ctx.SendError(err)
		return
	}

	// 返回結果
	ctx.Send(CodeSuccess, "success", data)
}

func (c *controller) FreeModeInReal(ctx *server.Context) {
	grpcCtx := ctx.MustGet("ctx").(context.Context)
	session := ctx.MustGet("session").(*playerModel.Session)

	record, err := c.betRecordService.Create(session)
	if err != nil {
		ctx.SendError(err)
		return
	}

	// 載入續存的 bet 值
	param := &betModel.Param{}
	bet, err := c.gameService.GetBet(grpcCtx, session.PlayerId)
	if err != nil {
		ctx.SendError(err)
		return
	}
	param.Bet = bet

	// 載入續存的 value 值
	value, err := c.gameService.GetValue(grpcCtx, session.PlayerId)
	if err != nil {
		ctx.SendError(err)
		return
	}
	param.Value = value

	// 取出一筆金蛇盤面
	results, err := c.resultFreeService.GameMode(GameModeReal).PopFirstItem(grpcCtx, session.PlayerId)
	if err != nil {
		ctx.SendError(err)
		return
	}

	// 將盤面數據轉換為 reels
	reels := c.reelsFreeService.ToReels(results)

	// 生成響應數據
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
	data.Balance = int(session.Balance)

	// 續存結果
	if err = c.SavePlayerGameInfo(ctx, gameResult, param); err != nil {
		ctx.SendError(err)
		return
	}

	if err = c.betRecordService.Update(record); err != nil {
		ctx.SendError(err)
		return
	}

	ctx.Send(CodeSuccess, "success", data)
}

// FinalFreeModeInReal 進入最後一個金蛇盤面流程
func (c *controller) FinalFreeModeInReal(ctx *server.Context) {
	grpcCtx := ctx.MustGet("ctx").(context.Context)
	session := ctx.MustGet("session").(*playerModel.Session)

	// 載入續存的 bet 值
	param := &betModel.Param{}
	bet, err := c.gameService.GetBet(grpcCtx, session.PlayerId)
	if err != nil {
		ctx.SendError(err)
		return
	}
	param.Bet = bet

	// 載入續存的 value 值
	value, err := c.gameService.GetValue(grpcCtx, session.PlayerId)
	if err != nil {
		ctx.SendError(err)
		return
	}
	param.Value = value

	// 取出最後一筆金蛇盤面
	results, err := c.resultFreeService.GameMode(GameModeReal).PopFirstItem(grpcCtx, session.PlayerId)
	if err != nil {
		ctx.SendError(err)
		return
	}

	// 將盤面數據轉換為 reels
	reels := c.reelsFreeService.ToReels(results)

	// 結算盤面
	winRate, lines, totalScore, err := c.Settle(param, reels)
	if err != nil {
		ctx.SendError(err)
		return
	}

	// 派彩更新餘額
	session.Balance += int64(totalScore)
	if err := c.playerService.UpdateBalance(grpcCtx, session.PlayerId, session.Balance); err != nil {
		ctx.SendError(err)
		return
	}

	// 生成響應數據
	spinResult := &betModel.SpinResult{}
	spinResult.Score = totalScore
	spinResult.SetLines(lines)
	spinResult.Symbols = results
	spinResult.Times = c.settleService.GetTimes(reels)

	gameResult := betModel.NewGameResult(SpinModeFree)
	gameResult.RandSymbol = c.reelsFreeService.GetMainSymbol(reels).ID
	gameResult.SpinResult = spinResult
	gameResult.WinRate = int(winRate)
	gameResult.TotalScore = totalScore
	gameResult.WinType = 0

	data := betModel.NewResponse(GameModeReal)
	data.GameResult = gameResult
	data.Balance = int(session.Balance)

	// 續存結果
	if err = c.SavePlayerGameInfo(ctx, gameResult, param); err != nil {
		ctx.SendError(err)
		return
	}

	// 返回結果
	ctx.Send(CodeSuccess, "success", data)
}
