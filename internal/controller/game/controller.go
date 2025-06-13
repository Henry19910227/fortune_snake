package game

import (
	"context"
	. "game_server_slots_fortune_snake/constants"
	betModel "game_server_slots_fortune_snake/internal/model/controller/game/bet"
	playerModel "game_server_slots_fortune_snake/internal/model/entity/player"
	"game_server_slots_fortune_snake/internal/server"
	gameService "game_server_slots_fortune_snake/internal/service/game"
	playerService "game_server_slots_fortune_snake/internal/service/player"
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
	playerService     playerService.Service
}

func New(gameService gameService.Service, gameDemoService gameService.Service,
	weightService weightService.Service, resultFreeService resultFreeService.Service,
	resultLoader resultLoader.Service, resultFreeLoader resultLoader.Service,
	reelsService reelsService.Service, reelsFreeService reelsService.Service,
	settleService settleService.Service, playerService playerService.Service) Controller {
	return &controller{gameService: gameService, gameDemoService: gameDemoService,
		weightService: weightService, resultFreeService: resultFreeService, resultLoader: resultLoader,
		resultFreeLoader: resultFreeLoader, reelsService: reelsService, reelsFreeService: reelsFreeService,
		settleService: settleService, playerService: playerService}
}

func (c *controller) EnterGame(ctx *server.Context) {
	//// 取得中間層處理好的數據
	//session := ctx.MustGet("session").(*playerModel.Session)
	//// 以 mode 獲取對應的 game service
	//service := c.getGameService(session.Mode)
	//// 執行 Enter Game 業務邏輯
	//input := &enter_game.Input{}
	//input.Ctx = ctx
	//input.Session = session
	//data, err := service.EnterGame(input)
	//if err != nil {
	//	ctx.SendError(err)
	//	return
	//}
	//// 返回結果
	//ctx.Send(CodeSuccess, "success", data)
}

func (c *controller) Bet(ctx *server.Context) {
	// 取得 session 數據
	session := ctx.MustGet("session").(*playerModel.Session)

	param := &betModel.Param{}
	if err := ctx.Bind(param); err != nil {
		ctx.SendError(err)
		return
	}

	// 進入試玩模式
	if session.Mode == "demo" {
		c.BetInDemo(ctx)
		return
	}
	c.BetInReal(ctx)
}

func (c *controller) BetInReal(ctx *server.Context) {
	session := ctx.MustGet("session").(*playerModel.Session)
	grpcCtx := ctx.MustGet("ctx").(context.Context)

	// 檢查剩餘免費盤面數量
	amount, err := c.resultFreeService.GameMode(GameModeReal).Amount(grpcCtx, session.PlayerUsername)
	if err != nil {
		ctx.SendError(err)
		return
	}

	// 金蛇中盤
	if amount > 1 {
		c.FreeModeInReal(ctx)
		return
	}

	// 進入最後一個金蛇盤面
	if amount == 1 {
		c.FinalFreeModeInReal(ctx)
		return
	}

	// 沒有尚未消費的盤面，則開新的一局，取得當前模式
	spinMode := c.gameService.SpinMode()

	// 扣除投注額

	// 進入金蛇模式
	if spinMode == SpinModeFree {
		c.StartFreeModeInReal(ctx)
		return
	}

	// 執行真錢環境的普通模式流程
	c.BaseModeInReal(ctx)
}

func (c *controller) BetInDemo(ctx *server.Context) {
	session := ctx.MustGet("session").(*playerModel.Session)
	grpcCtx := ctx.MustGet("ctx").(context.Context)

	// 檢查redis是否有免費盤面List尚未消費(real)，如有則代表當前當前有未完成的金蛇模式
	amount, err := c.resultFreeService.GameMode(GameModeDemo).Amount(grpcCtx, session.PlayerUsername)
	if err != nil {
		ctx.SendError(err)
		return
	}

	// 金蛇盤面進行中
	if amount > 1 {
		c.FreeModeInDemo(ctx)
		return
	}

	// 進入最後一個金蛇盤面
	if amount == 1 {
		c.FinalFreeModeInDemo(ctx)
		return
	}

	// 沒有尚未消費的盤面，則開新的一局，取得當前模式
	spinMode := c.gameService.SpinMode()

	// 扣除試玩投注額

	// 開始第一局金蛇模式
	if spinMode == SpinModeFree {
		c.StartFreeModeInDemo(ctx)
		return
	}

	// 一般模式
	c.BaseModeInDemo(ctx)
}

func (c *controller) Settle(ctx *server.Context) {

}

func (c *controller) getGameService(mode string) gameService.Service {
	// 試玩模式
	if mode == "demo" {
		return c.gameDemoService
	}
	// 真錢模式
	return c.gameService
}
