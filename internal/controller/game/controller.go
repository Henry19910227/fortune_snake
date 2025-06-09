package game

import (
	"context"
	"fmt"
	. "game_server_slots_fortune_snake/constants"
	playerModel "game_server_slots_fortune_snake/internal/model/entity/player"
	"game_server_slots_fortune_snake/internal/model/service/game/enter_game"
	"game_server_slots_fortune_snake/internal/server"
	gameService "game_server_slots_fortune_snake/internal/service/game"
	reelsService "game_server_slots_fortune_snake/internal/service/reels"
	reelsFreeService "game_server_slots_fortune_snake/internal/service/reels_free"
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
	reelsFreeService  reelsFreeService.Service
	settleService     settleService.Service
}

func New(gameService gameService.Service, gameDemoService gameService.Service,
	weightService weightService.Service, resultFreeService resultFreeService.Service,
	resultLoader resultLoader.Service, resultFreeLoader resultLoader.Service,
	reelsService reelsService.Service, reelsFreeService reelsFreeService.Service,
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
	// 檢查redis是否有免費盤面List尚未消費(real)，如有則代表當前當前有未完成的金蛇模式
	grpcCtx := ctx.MustGet("ctx").(context.Context)
	amount, err := c.resultFreeService.GameMode(GameModeReal).Amount(grpcCtx, int(session.PlayerId))
	if err != nil {
		ctx.SendError(err)
		return
	}

	if amount > 0 {
		// 彈出一筆金蛇盤面
		result, err := c.resultFreeService.GameMode(GameModeReal).PopFirstItem(grpcCtx, int(session.PlayerId))
		if err != nil {
			ctx.SendError(err)
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
		// 獲取金蛇盤面
		results, err := c.resultFreeLoader.Random(rate)
		if err != nil {
			ctx.SendError(err)
			return
		}
		// 執行真錢環境的免費模式流程
		c.FreeModeInReal(ctx, session, results)
		return
	}
	// 執行真錢環境的普通模式流程
	c.BaseModeInReal(ctx, session)
}

func (c *controller) BetInDemo(ctx *server.Context, session *playerModel.Session) {
	// 檢查redis是否有免費盤面List尚未消費(real)，如有則代表當前當前有未完成的金蛇模式
	grpcCtx := ctx.MustGet("ctx").(context.Context)
	amount, err := c.resultFreeService.GameMode(GameModeDemo).Amount(grpcCtx, int(session.PlayerId))
	if err != nil {
		ctx.SendError(err)
		return
	}

	if amount > 0 {
		// 彈出一筆金蛇盤面
		result, err := c.resultFreeService.GameMode(GameModeDemo).PopFirstItem(grpcCtx, int(session.PlayerId))
		if err != nil {
			ctx.SendError(err)
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
		if err := c.resultFreeService.SaveItems(context.Background(), int(session.PlayerId), results[1:]); err != nil {
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

	// 結算 reels
	lines, err := c.settleService.GetWinLines(1, 1000, reels)
	if err != nil {
		ctx.SendError(err)
		return
	}
	fmt.Println(lines)

	// 續存結算的結果 LastResults key

	// 回傳結果
}

func (c *controller) FreeModeInReal(ctx *server.Context, session *playerModel.Session, results [][][]int) {
	// 結算第一個金蛇盤面

	// 續存結算的金蛇盤面結果 LastResults key

	// 回傳結果
}

func (c *controller) getRate(spinMode int, rtp float64) (float64, error) {
	// 一般模式
	if spinMode == 0 {
		return c.weightService.RandomBaseWeightRate(rtp)
	}
	// 金蛇模式
	return c.weightService.RandomFreeWeightRate(rtp)
}

func (c *controller) getGameService(mode string) gameService.Service {
	// 試玩模式
	if mode == "demo" {
		return c.gameDemoService
	}
	// 真錢模式
	return c.gameService
}

func (c *controller) getResultService(spinMode int) resultLoader.Service {
	// 試玩模式
	if spinMode == 0 {
		return c.resultFreeLoader
	}
	// 真錢模式
	return c.resultLoader
}
