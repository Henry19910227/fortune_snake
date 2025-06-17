package game

import (
	"context"
	"fmt"
	. "game_server_slots_fortune_snake/constants"
	gameSymbol "game_server_slots_fortune_snake/internal/model/controller/game/symbol"
	playerModel "game_server_slots_fortune_snake/internal/model/entity/player"
	"game_server_slots_fortune_snake/internal/server"
	betRecordService "game_server_slots_fortune_snake/internal/service/bet_record"
	gameService "game_server_slots_fortune_snake/internal/service/game"
	playerService "game_server_slots_fortune_snake/internal/service/player"
	reelsService "game_server_slots_fortune_snake/internal/service/reels"
	resultFreeService "game_server_slots_fortune_snake/internal/service/result_free"
	resultLoader "game_server_slots_fortune_snake/internal/service/result_loader"
	settleService "game_server_slots_fortune_snake/internal/service/settle"
	symbolService "game_server_slots_fortune_snake/internal/service/symbol"
	weightService "game_server_slots_fortune_snake/internal/service/weight"
)

type controller struct {
	gameService       gameService.Service // 真實模式 service
	weightService     weightService.Service
	resultFreeService resultFreeService.Service
	resultLoader      resultLoader.Service
	resultFreeLoader  resultLoader.Service
	reelsService      reelsService.Service
	reelsFreeService  reelsService.Service
	settleService     settleService.Service
	playerService     playerService.Service
	betRecordService  betRecordService.Service
	symbolService     symbolService.Service
}

func New(gameService gameService.Service,
	weightService weightService.Service,
	resultFreeService resultFreeService.Service,
	resultLoader resultLoader.Service,
	resultFreeLoader resultLoader.Service,
	reelsService reelsService.Service,
	reelsFreeService reelsService.Service,
	settleService settleService.Service,
	playerService playerService.Service,
	betRecordService betRecordService.Service,
	symbolService symbolService.Service) Controller {
	return &controller{gameService: gameService,
		weightService: weightService, resultFreeService: resultFreeService, resultLoader: resultLoader,
		resultFreeLoader: resultFreeLoader, reelsService: reelsService, reelsFreeService: reelsFreeService,
		settleService: settleService, playerService: playerService, betRecordService: betRecordService, symbolService: symbolService}
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

func (c *controller) PreBet(ctx *server.Context) {
	// 取得 session 數據
	session := ctx.MustGet("session").(*playerModel.Session)
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
	amount, err := c.resultFreeService.Amount(grpcCtx, GameModeReal, session.PlayerId)
	if err != nil {
		ctx.SendError(err)
		return
	}

	// 金蛇中盤
	if amount > 1 {
		ctx.Set("condition", FreeModeInReal)
		return
	}

	// 進入最後一個金蛇盤面
	if amount == 1 {
		ctx.Set("condition", FinalFreeModeInReal)
		return
	}

	// 沒有尚未消費的盤面，則開新的一局，取得當前模式
	spinMode := c.gameService.SpinMode()

	// 進入金蛇模式
	if spinMode == SpinModeFree {
		ctx.Set("condition", StartFreeModeInReal)
		return
	}

	// 執行真錢環境的普通模式流程
	ctx.Set("condition", BaseModeInReal)
}

func (c *controller) BetInDemo(ctx *server.Context) {
	session := ctx.MustGet("session").(*playerModel.Session)
	grpcCtx := ctx.MustGet("ctx").(context.Context)

	// 檢查剩餘免費盤面數量
	amount, err := c.resultFreeService.Amount(grpcCtx, GameModeDemo, session.PlayerId)
	if err != nil {
		ctx.SendError(err)
		return
	}

	// 金蛇盤面進行中
	if amount > 1 {
		ctx.Set("condition", FreeModeInDemo)
		return
	}

	// 進入最後一個金蛇盤面
	if amount == 1 {
		ctx.Set("condition", FinalFreeModeInDemo)
		return
	}

	// 沒有尚未消費的盤面，則開新的一局，取得當前模式
	spinMode := c.gameService.SpinMode()

	// 開始第一局金蛇模式
	if spinMode == SpinModeFree {
		ctx.Set("condition", StartFreeModeInDemo)
		return
	}

	// 一般模式
	ctx.Set("condition", BaseModeInDemo)
}

func (c *controller) Bet(ctx *server.Context) {
	condition := ctx.MustGet("condition").(string)

	switch condition {
	case BaseModeInReal:
		c.BaseModeInReal(ctx)
	case StartFreeModeInReal:
		c.StartFreeModeInReal(ctx)
	case FreeModeInReal:
		c.FreeModeInReal(ctx)
	case FinalFreeModeInReal:
		c.FinalFreeModeInReal(ctx)
	case BaseModeInDemo:
		c.BaseModeInDemo(ctx)
	case StartFreeModeInDemo:
		c.StartFreeModeInDemo(ctx)
	case FreeModeInDemo:
		c.FreeModeInDemo(ctx)
	case FinalFreeModeInDemo:
		c.FinalFreeModeInDemo(ctx)
	default:
		ctx.SendError(fmt.Errorf("unknown condition: %s", condition))
	}
}

func (c *controller) Symbols(ctx *server.Context) {
	items := c.symbolService.GetSymbols()
	symbols := make([]*gameSymbol.Item, 0)
	for _, item := range items {
		symbol := &gameSymbol.Item{}
		symbol.ID = item.ID
		symbol.Name = item.Name
		symbol.Pow = item.Pow
		symbols = append(symbols, symbol)
	}

	data := gameSymbol.Response{}
	data.Symbols = symbols

	ctx.Send(CodeSuccess, "success", data)
}
