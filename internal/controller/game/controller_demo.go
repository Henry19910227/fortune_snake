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
	grpcCtx := ctx.MustGet("ctx").(context.Context)
	param := betModel.NewParam()

	// 獲取投注參數
	if err := ctx.Bind(param); err != nil {
		ctx.SendError(err)
		return
	}

	// 計算總投注額
	totalBet := param.Bet * param.Value * param.Multiplier

	// 創建紀錄
	record, err := c.betRecordService.Create(session, param, int64(totalBet))
	if err != nil {
		ctx.SendError(err)
		return
	}

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
		_ = c.betRecordService.UpdateToFailed(record)
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

	data := betModel.NewResponse(GameModeDemo)
	data.GameResult = gameResult
	data.Balance = int(session.Balance)

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
		ctx.SendError(err)
		return
	}

	// 更新投注紀錄
	record.Result = gameResult.Encode()
	record.Amount = int64(totalBet)
	record.AmountWin = int64(totalScore)
	if err = c.betRecordService.UpdateToFinished(record); err != nil {
		ctx.SendError(err)
		return
	}

	// 返回結果
	ctx.Send(CodeSuccess, "success", data)
}
