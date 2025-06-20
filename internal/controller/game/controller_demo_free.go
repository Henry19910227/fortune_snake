package game

import (
	"context"
	. "game_server_slots_fortune_snake/constants"
	betModel "game_server_slots_fortune_snake/internal/model/controller/game/bet"
	playerModel "game_server_slots_fortune_snake/internal/model/entity/player"
	"game_server_slots_fortune_snake/internal/model/service/bet_record/create_by_tx"
	"game_server_slots_fortune_snake/internal/model/service/result_free/save_items"
	"game_server_slots_fortune_snake/internal/server"
	"game_server_slots_fortune_snake/util"
)

// StartFreeModeInDemo 開始金蛇盤面流程
func (c *controller) StartFreeModeInDemo(ctx *server.Context) {
	session := ctx.MustGet("session").(*playerModel.Session)
	grpcCtx := ctx.MustGet("ctx").(context.Context)
	param := betModel.NewParam()

	// 獲取投注參數
	if err := ctx.Bind(param); err != nil {
		ctx.SendError(err)
		return
	}

	// 計算總投注額
	totalBet := param.Bet * param.Value * param.Multiplier
	if param.Bonus {
		totalBet = int(float64(totalBet) * 1.5)
	}

	// 創建紀錄
	record, err := c.betRecordService.Create(session, param, int64(totalBet))
	if err != nil {
		ctx.SendError(err)
		return
	}

	// 儲存 father ID
	if err = c.gameService.SaveFatherID(grpcCtx, GameModeDemo, session.PlayerId, record.TransactionId); err != nil {
		_ = c.betRecordService.UpdateToFailed(record)
		ctx.SendError(err)
		return
	}

	// 儲存 round ID
	if err = c.gameService.SaveRoundID(grpcCtx, GameModeDemo, session.PlayerId, record.RoundId); err != nil {
		_ = c.betRecordService.UpdateToFailed(record)
		ctx.SendError(err)
		return
	}

	// 用戶餘額減去總投注額
	session.Balance -= int64(totalBet)

	// 獲取賠率
	rate, err := c.weightService.RandomFreeWeightRate(float64(session.GameRtp))
	if err != nil {
		_ = c.betRecordService.UpdateToFailed(record)
		ctx.SendError(err)
		return
	}

	// 從內存隨機取得多筆金蛇盤面
	resultsList, err := c.resultFreeLoader.Random(rate)
	if err != nil {
		_ = c.betRecordService.UpdateToFailed(record)
		ctx.SendError(err)
		return
	}

	// 緩存金蛇盤面數據集
	input := save_items.Param{}
	input.Ctx = grpcCtx
	input.Session = session
	input.Items = resultsList
	if err := c.resultFreeDemoService.SaveItems(input); err != nil {
		_ = c.betRecordService.UpdateToFailed(record)
		ctx.SendError(err)
		return
	}

	// 取出一筆金蛇盤面
	results, err := c.resultFreeDemoService.PopFirstItem(grpcCtx, session)
	if err != nil {
		_ = c.betRecordService.UpdateToFailed(record)
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
	gameResult.RandSymbol = util.PointerInt(c.reelsFreeService.GetMainSymbol(reels).ID)
	gameResult.SpinResult = spinResult
	gameResult.WinRate = 0
	gameResult.TotalScore = 0
	gameResult.WinType = 0

	data := betModel.NewResponse(GameModeDemo)
	data.GameResult = gameResult
	data.Balance = int(session.Balance)

	// 續存結果
	if err = c.SavePlayerGameInfo(ctx, gameResult, param); err != nil {
		_ = c.betRecordService.UpdateToFailed(record)
		ctx.SendError(err)
		return
	}

	// 更新投注紀錄
	record.Result = gameResult.Encode()
	if err = c.betRecordService.UpdateToFinished(record); err != nil {
		ctx.SendError(err)
		return
	}

	// 更新餘額
	if err := c.playerService.UpdateBalance(grpcCtx, session.PlayerId, session.Balance); err != nil {
		_ = c.betRecordService.UpdateToFailed(record)
		ctx.SendError(err)
		return
	}

	// 更新遊戲結果
	_, err = c.gameResultService.Create(session, gameResult)
	if err != nil {
		_ = c.betRecordService.UpdateToFailed(record)
		ctx.SendError(err)
		return
	}

	// 返回結果
	ctx.Send(CodeSuccess, "success", data)
}

// FreeModeInDemo 進行中的金蛇盤面流程
func (c *controller) FreeModeInDemo(ctx *server.Context) {
	session := ctx.MustGet("session").(*playerModel.Session)
	grpcCtx := ctx.MustGet("ctx").(context.Context)
	param := betModel.NewParam()

	// 恢復續存投注參數
	if err := c.RestoreParam(ctx, param); err != nil {
		ctx.SendError(err)
		return
	}

	// 獲取 father id
	fatherID, err := c.gameService.GetFatherID(grpcCtx, GameModeDemo, session.PlayerId)
	if err != nil {
		ctx.SendError(err)
		return
	}

	// 獲取 round id
	roundID, err := c.gameService.GetRoundID(grpcCtx, GameModeDemo, session.PlayerId)
	if err != nil {
		ctx.SendError(err)
		return
	}

	// 創建紀錄
	input := &create_by_tx.Input{}
	input.Session = session
	input.Param = param
	input.TotalBet = 0
	input.TransactionID = fatherID
	input.RoundID = roundID
	record, err := c.betRecordService.CreateByTransactionID(input)
	if err != nil {
		ctx.SendError(err)
		return
	}

	// 從緩存中取出一筆金蛇盤面
	results, err := c.resultFreeDemoService.PopFirstItem(grpcCtx, session)
	if err != nil {
		_ = c.betRecordService.UpdateToFailed(record)
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
	gameResult.RandSymbol = util.PointerInt(c.reelsFreeService.GetMainSymbol(reels).ID)
	gameResult.SpinResult = spinResult
	gameResult.WinRate = 0
	gameResult.TotalScore = 0
	gameResult.WinType = 0

	data := betModel.NewResponse(GameModeDemo)
	data.GameResult = gameResult
	data.Balance = int(session.Balance)

	// 續存結果
	if err = c.SavePlayerGameInfo(ctx, gameResult, param); err != nil {
		_ = c.betRecordService.UpdateToFailed(record)
		ctx.SendError(err)
		return
	}

	// 更新遊戲結果
	_, err = c.gameResultService.Create(session, gameResult)
	if err != nil {
		_ = c.betRecordService.UpdateToFailed(record)
		ctx.SendError(err)
		return
	}

	// 更新投注紀錄
	record.Result = gameResult.Encode()
	if err = c.betRecordService.UpdateToFinished(record); err != nil {
		ctx.SendError(err)
		return
	}

	// 返回結果
	ctx.Send(CodeSuccess, "success", data)
}

// FinalFreeModeInDemo 進入最後一個金蛇盤面流程
func (c *controller) FinalFreeModeInDemo(ctx *server.Context) {
	grpcCtx := ctx.MustGet("ctx").(context.Context)
	session := ctx.MustGet("session").(*playerModel.Session)
	param := betModel.NewParam()

	// 恢復續存投注參數
	if err := c.RestoreParam(ctx, param); err != nil {
		ctx.SendError(err)
		return
	}

	// 獲取 father id
	fatherID, err := c.gameService.GetFatherID(grpcCtx, GameModeDemo, session.PlayerId)
	if err != nil {
		ctx.SendError(err)
		return
	}

	// 獲取 round id
	roundID, err := c.gameService.GetRoundID(grpcCtx, GameModeDemo, session.PlayerId)
	if err != nil {
		ctx.SendError(err)
		return
	}

	// 創建紀錄
	input := &create_by_tx.Input{}
	input.Session = session
	input.Param = param
	input.TotalBet = 0
	input.TransactionID = fatherID
	input.RoundID = roundID
	record, err := c.betRecordService.CreateByTransactionID(input)
	if err != nil {
		ctx.SendError(err)
		return
	}

	// 從緩存中取出最後一筆金蛇盤面
	results, err := c.resultFreeDemoService.PopFirstItem(grpcCtx, session)
	if err != nil {
		_ = c.betRecordService.UpdateToFailed(record)
		ctx.SendError(err)
		return
	}

	// 將盤面數據轉換為 reels
	reels := c.reelsFreeService.ToReels(results)

	// 結算盤面
	winRate, lines, totalScore, err := c.Settle(param, reels)
	if err != nil {
		_ = c.betRecordService.UpdateToFailed(record)
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
	gameResult.RandSymbol = util.PointerInt(c.reelsFreeService.GetMainSymbol(reels).ID)
	gameResult.SpinResult = spinResult
	gameResult.WinRate = int(winRate)
	gameResult.TotalScore = totalScore
	gameResult.WinType = 0

	data := betModel.NewResponse(GameModeDemo)
	data.GameResult = gameResult
	data.Balance = int(session.Balance)

	// 續存結果
	if err = c.SavePlayerGameInfo(ctx, gameResult, param); err != nil {
		_ = c.betRecordService.UpdateToFailed(record)
		ctx.SendError(err)
		return
	}

	// 更新遊戲結果
	_, err = c.gameResultService.Create(session, gameResult)
	if err != nil {
		_ = c.betRecordService.UpdateToFailed(record)
		ctx.SendError(err)
		return
	}

	// 派彩更新餘額
	session.Balance += int64(totalScore)
	if err := c.playerService.UpdateBalance(grpcCtx, session.PlayerId, session.Balance); err != nil {
		_ = c.betRecordService.UpdateToFailed(record)
		ctx.SendError(err)
		return
	}

	// 更新投注紀錄
	record.Result = gameResult.Encode()
	record.AmountWin = int64(totalScore)
	if err = c.betRecordService.UpdateToFinished(record); err != nil {
		ctx.SendError(err)
		return
	}

	// 返回結果
	ctx.Send(CodeSuccess, "success", data)
}
