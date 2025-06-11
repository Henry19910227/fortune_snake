package game

import (
	"context"
	. "game_server_slots_fortune_snake/constants"
	betModel "game_server_slots_fortune_snake/internal/model/controller/game/bet"
	playerModel "game_server_slots_fortune_snake/internal/model/entity/player"
	"game_server_slots_fortune_snake/internal/model/service/game/enter_game"
	"game_server_slots_fortune_snake/internal/server"
	gameService "game_server_slots_fortune_snake/internal/service/game"
	reelsService "game_server_slots_fortune_snake/internal/service/reels"
	resultFreeService "game_server_slots_fortune_snake/internal/service/result_free"
	resultLoader "game_server_slots_fortune_snake/internal/service/result_loader"
	settleService "game_server_slots_fortune_snake/internal/service/settle"
	weightService "game_server_slots_fortune_snake/internal/service/weight"
)

type controller struct {
	gameService       gameService.Service // 真實模式 service
	gameDemoService   gameService.Service // 試玩模式 service
	weightService     weightService.Service
	resultFreeService resultFreeService.Service
	resultLoader      resultLoader.Service
	resultFreeLoader  resultLoader.Service
	reelsService      reelsService.Service
	reelsFreeService  reelsService.Service
	settleService     settleService.Service
}

func New(gameService gameService.Service, gameDemoService gameService.Service,
	weightService weightService.Service, resultFreeService resultFreeService.Service,
	resultLoader resultLoader.Service, resultFreeLoader resultLoader.Service,
	reelsService reelsService.Service, reelsFreeService reelsService.Service,
	settleService settleService.Service) Controller {
	return &controller{gameService: gameService, gameDemoService: gameDemoService,
		weightService: weightService, resultFreeService: resultFreeService, resultLoader: resultLoader,
		resultFreeLoader: resultFreeLoader, reelsService: reelsService, reelsFreeService: reelsFreeService,
		settleService: settleService}
}

func (c *controller) EnterGame(ctx *server.Context) {
	// 取得中間層處理好的數據
	session := ctx.MustGet("session").(*playerModel.Session)
	// 以 mode 獲取對應的 game service
	service := c.getGameService(session.Mode)
	// 執行 Enter Game 業務邏輯
	input := &enter_game.Input{}
	input.Ctx = ctx
	input.Session = session
	data, err := service.EnterGame(input)
	if err != nil {
		ctx.SendError(err)
		return
	}
	// 返回結果
	ctx.Send(CodeSuccess, "success", data)
}

func (c *controller) Bet(ctx *server.Context) {
	// 取得 session 數據
	session := ctx.MustGet("session").(*playerModel.Session)

	// 進入試玩模式
	if session.Mode == "demo" {
		c.BetInDemo(ctx, session)
		return
	}
	c.BetInReal(ctx, session)
}

func (c *controller) BetInReal(ctx *server.Context, session *playerModel.Session) {
	grpcCtx := ctx.MustGet("ctx").(context.Context)

	// 檢查剩餘免費盤面數量
	amount, err := c.resultFreeService.GameMode(GameModeReal).Amount(grpcCtx, session.PlayerUsername)
	if err != nil {
		ctx.SendError(err)
		return
	}

	// 檢查redis是否有免費盤面List尚未消費(real)，如有則代表當前當前有未完成的金蛇模式
	if amount > 0 {
		// 取出一筆金蛇盤面
		result, err := c.resultFreeService.GameMode(GameModeReal).PopFirstItem(grpcCtx, session.PlayerUsername)
		if err != nil {
			ctx.SendError(err)
			return
		}
		// 執行試玩環境的免費模式流程
		c.FreeModeInDemo(ctx, session, result)
		return
	}

	// 沒有尚未消費的盤面，則開新的一局，取得當前模式
	spinMode := c.gameService.SpinMode()

	// 進入金蛇模式
	if spinMode == SpinModeFree {
		// 獲取賠率
		rate, err := c.weightService.RandomFreeWeightRate(float64(session.GameRtp))
		if err != nil {
			ctx.SendError(err)
			return
		}
		// 生成隨機金蛇盤面
		results, err := c.resultFreeLoader.Random(rate)
		if err != nil {
			ctx.SendError(err)
			return
		}
		// 取出第一個金蛇盤面
		result := results[0]

		// 緩存剩餘金蛇盤面
		if err := c.resultFreeService.SaveItems(context.Background(), session.PlayerUsername, results[1:]); err != nil {
			ctx.SendError(err)
			return
		}
		// 執行真錢環境的免費模式流程
		c.FreeModeInReal(ctx, session, result)
		return
	}
	// 執行真錢環境的普通模式流程
	c.BaseModeInReal(ctx, session)
}

func (c *controller) BetInDemo(ctx *server.Context, session *playerModel.Session) {
	// 檢查redis是否有免費盤面List尚未消費(real)，如有則代表當前當前有未完成的金蛇模式
	grpcCtx := ctx.MustGet("ctx").(context.Context)
	amount, err := c.resultFreeService.GameMode(GameModeDemo).Amount(grpcCtx, session.PlayerUsername)
	if err != nil {
		ctx.SendError(err)
		return
	}

	if amount > 0 {
		// 彈出一筆金蛇盤面
		result, err := c.resultFreeService.GameMode(GameModeDemo).PopFirstItem(grpcCtx, session.PlayerUsername)
		if err != nil {
			ctx.SendError(err)
			return
		}
		// 執行試玩環境的免費模式流程
		c.FreeModeInDemo(ctx, session, result)
		return
	}

	// 沒有尚未消費的盤面，則開新的一局，取得當前模式
	spinMode := c.gameService.SpinMode()

	// 進入金蛇模式
	if spinMode == SpinModeFree {
		// 獲取賠率
		rate, err := c.weightService.RandomFreeWeightRate(float64(session.GameRtp))
		if err != nil {
			ctx.SendError(err)
			return
		}
		// 生成隨機金蛇盤面
		results, err := c.resultFreeLoader.Random(rate)
		if err != nil {
			ctx.SendError(err)
			return
		}
		// 取出第一個金蛇盤面
		result := results[0]

		// 緩存剩餘金蛇盤面
		if err := c.resultFreeService.SaveItems(context.Background(), session.PlayerUsername, results[1:]); err != nil {
			ctx.SendError(err)
			return
		}

		// 執行試玩環境的免費模式流程
		c.FreeModeInDemo(ctx, session, result)
		return
	}
	c.BaseModeInDemo(ctx, session)
}

func (c *controller) BaseModeInDemo(ctx *server.Context, session *playerModel.Session) {
	// 獲取賠率
	rate, err := c.weightService.RandomBaseWeightRate(float64(session.GameRtp))
	if err != nil {
		ctx.SendError(err)
		return
	}

	// 取得隨機盤面
	_, err = c.resultLoader.Random(rate)
	if err != nil {
		ctx.SendError(err)
		return
	}

	// 結算盤面

	// 餘額計算

	// 緩存盤面結果 lastResults key

	// 回傳結果
}

func (c *controller) BaseModeInReal(ctx *server.Context, session *playerModel.Session) {
	// 獲取賠率
	rate, err := c.weightService.RandomBaseWeightRate(float64(session.GameRtp))
	if err != nil {
		ctx.SendError(err)
		return
	}

	// 取得隨機盤面
	_, err = c.resultLoader.Random(rate)
	if err != nil {
		ctx.SendError(err)
		return
	}
	// 結算盤面

	// 續存結算的盤面結果 LastResults key

	// 餘額計算

	// 回傳結果
}

func (c *controller) FreeModeInDemo(ctx *server.Context, session *playerModel.Session, results [][]int) {
	// 將盤面數據轉換為 reels
	reels := c.reelsFreeService.ToReels(results)

	// 計算中獎線
	lines, err := c.settleService.GetWinLines(1, 1000, reels)
	if err != nil {
		ctx.SendError(err)
		return
	}

	// 計算賠率
	rate, err := c.settleService.GetRate(1, 1000, reels)
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

	// 續存結算的結果 LastResults key

	// 同步餘額

	// 回傳結果
	spinResult := betModel.NewSpinResult(lines)
	spinResult.Score = totalScore
	spinResult.Symbols = c.reelsFreeService.ToResults(reels)
	spinResult.Times = c.settleService.GetTimes(reels)

	gameResult := betModel.NewGameResult(SpinModeFree)
	gameResult.SpinResult = spinResult
	gameResult.WinRate = int(rate)
	gameResult.TotalScore = totalScore
	gameResult.WinType = 0

	data := betModel.NewResponse(GameModeDemo)
	data.GameResult = gameResult
	data.ScoreTry = 10000

	// 返回結果
	ctx.Send(CodeSuccess, "success", data)
}

func (c *controller) FreeModeInReal(ctx *server.Context, session *playerModel.Session, results [][]int) {
	// 將盤面數據轉換為 reels
	reels := c.reelsFreeService.ToReels(results)

	// 計算中獎線
	lines, err := c.settleService.GetWinLines(1, 1000, reels)
	if err != nil {
		ctx.SendError(err)
		return
	}

	// 計算賠率
	rate, err := c.settleService.GetRate(1, 1000, reels)
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

	// 續存結算的結果 LastResults key

	// 同步餘額

	// 回傳結果
	spinResult := betModel.NewSpinResult(lines)
	spinResult.Score = totalScore
	spinResult.Symbols = c.reelsFreeService.ToResults(reels)
	spinResult.Times = c.settleService.GetTimes(reels)

	gameResult := betModel.NewGameResult(SpinModeFree)
	gameResult.SpinResult = spinResult
	gameResult.WinRate = int(rate)
	gameResult.TotalScore = totalScore
	gameResult.WinType = 0

	data := betModel.NewResponse(GameModeReal)
	data.GameResult = gameResult
	data.ScoreTry = 10000

	// 返回結果
	ctx.Send(CodeSuccess, "success", data)
}

func (c *controller) getGameService(mode string) gameService.Service {
	// 試玩模式
	if mode == "demo" {
		return c.gameDemoService
	}
	// 真錢模式
	return c.gameService
}
