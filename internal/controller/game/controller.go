package game

import (
	"game_server_slots_fortune_snake/constants"
	playerModel "game_server_slots_fortune_snake/internal/model/entity/player"
	"game_server_slots_fortune_snake/internal/model/service/game/enter_game"
	"game_server_slots_fortune_snake/internal/server"
	gameService "game_server_slots_fortune_snake/internal/service/game"
	resultService "game_server_slots_fortune_snake/internal/service/result"
	weightService "game_server_slots_fortune_snake/internal/service/weight"
)

type controller struct {
	gameService       gameService.Service // 真實模式 service
	gameDemoService   gameService.Service // 試玩模式 service
	weightService     weightService.Service
	resultService     resultService.Service
	resultFreeService resultService.Service
}

func New(gameService gameService.Service, gameDemoService gameService.Service,
	weightService weightService.Service, resultService resultService.Service,
	resultFreeService resultService.Service) Controller {
	return &controller{gameService: gameService, gameDemoService: gameDemoService,
		weightService: weightService, resultService: resultService,
		resultFreeService: resultFreeService}
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
	ctx.Send(constants.CodeSuccess, "success", data)
}

func (c *controller) Bet(ctx *server.Context) {
	// 取得 session 數據
	session := ctx.MustGet("session").(*playerModel.Session)

	// 進入試玩模式
	if session.Mode == "demo" {
		c.BetForDemo(ctx, float64(session.GameRtp))
		return
	}
	c.BetForReal(ctx, float64(session.GameRtp))
}

func (c *controller) BetForReal(ctx *server.Context, rtp float64) {
	// 檢查redis是否有免費盤面List尚未消費(real)，如有則代表當前當前有未完成的金蛇模式
	// results := c.gameService.RestoreResults

	// if len(results) > 0 {
	//     執行剩餘未完成的金蛇模式業務
	//     return
	// }

	// 沒有尚未消費的盤面，擇開新的一局，取得當前模式
	spinMode := c.gameService.SpinMode()

	// 進入金蛇模式
	if spinMode == 1 {
		c.FreeModeInReal(ctx, rtp)
		return
	}
	c.BaseModeInReal(ctx, rtp)
}

func (c *controller) BetForDemo(ctx *server.Context, rtp float64) {
	// 檢查redis是否有免費盤面List尚未消費(real)，如有則代表當前當前有未完成的金蛇模式
	// results := c.gameService.RestoreResults

	// if len(results) > 0 {
	//     執行剩餘未完成的金蛇模式業務
	//     return
	// }

	// 沒有尚未消費的盤面，擇開新的一局，取得當前模式
	spinMode := c.gameService.SpinMode()

	// 進入金蛇模式
	if spinMode == 1 {
		c.FreeModeInDemo(ctx, rtp)
		return
	}
	c.BaseModeInDemo(ctx, rtp)
}

func (c *controller) BaseModeInDemo(ctx *server.Context, rtp float64) {
	// 獲取賠率
	rate, err := c.weightService.RandomBaseWeightRate(rtp)
	if err != nil {
		ctx.SendError(err)
		return
	}

	// 取得隨機盤面
	_, err = c.resultService.Random(rate)
	if err != nil {
		ctx.SendError(err)
		return
	}

	// 緩存盤面 lastResults key

	// 結算盤面

	// 餘額計算

	// 回傳結果
}

func (c *controller) FreeModeInDemo(ctx *server.Context, rtp float64) {
	// 獲取賠率
	rate, err := c.weightService.RandomFreeWeightRate(rtp)
	if err != nil {
		ctx.SendError(err)
		return
	}
	// 獲取金蛇盤面
	_, err = c.resultFreeService.Random(rate)
	if err != nil {
		ctx.SendError(err)
		return
	}
	// 取出第一個金蛇盤面

	// 緩存第一個金蛇盤面 LastResults key

	// 緩存剩餘金蛇盤面至 FreeResults key

	// 結算第一個金蛇盤面

	// 回傳結果
}

func (c *controller) BaseModeInReal(ctx *server.Context, rtp float64) {
	// 獲取賠率
	rate, err := c.weightService.RandomBaseWeightRate(rtp)
	if err != nil {
		ctx.SendError(err)
		return
	}

	// 取得隨機盤面
	_, err = c.resultService.Random(rate)
	if err != nil {
		ctx.SendError(err)
		return
	}

	// 緩存盤面 lastResults key

	// 結算盤面

	// 餘額計算

	// 回傳結果
}

func (c *controller) FreeModeInReal(ctx *server.Context, rtp float64) {
	// 獲取賠率
	rate, err := c.weightService.RandomFreeWeightRate(rtp)
	if err != nil {
		ctx.SendError(err)
		return
	}
	// 獲取金蛇盤面
	_, err = c.resultFreeService.Random(rate)
	if err != nil {
		ctx.SendError(err)
		return
	}
	// 取出第一個金蛇盤面

	// 緩存第一個金蛇盤面 LastResults key

	// 緩存剩餘金蛇盤面至 FreeResults key

	// 結算第一個金蛇盤面

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

func (c *controller) getResultService(spinMode int) resultService.Service {
	// 試玩模式
	if spinMode == 0 {
		return c.resultFreeService
	}
	// 真錢模式
	return c.resultService
}
