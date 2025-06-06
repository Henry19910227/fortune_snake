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

	// 取得對應模式下的 game service(real/demo)
	gameSvc := c.getGameService(session.Mode)
	// 檢查redis是否有免費盤面List尚未消費(real/demo)，如有則代表當前是金蛇模式
	// results := gameSvc.RestoreResults
	// if len(results) > 0 {
	//     執行金蛇模式業務
	//     return
	// }

	// 沒有免費盤面則重新隨機出旋轉模式
	spinMode := gameSvc.SpinMode()

	// 取得賠率
	rate, err := c.getRate(spinMode, float64(session.GameRtp))
	if err != nil {
		ctx.SendError(err)
		return
	}

	// 取得隨機盤面
	resultSvc := c.getResultService(spinMode)
	_, err = resultSvc.Random(rate)
	if err != nil {
		ctx.SendError(err)
		return
	}

	// if len(results) > 0 {
	//     執行金蛇模式業務
	//     return
	// }

	// 緩存當前盤面

	// 緩存多個金蛇盤面

}

func (c *controller) betForReal(ctx *server.Context, rtp float64) {
	// 檢查redis是否有免費盤面List尚未消費(real)，如有則代表當前是金蛇模式
	// results := c.gameService.RestoreResults

	// if len(results) > 0 {
	//     執行金蛇模式業務
	//     return
	// }

	// 沒有免費盤面則重新隨機出旋轉模式
	spinMode := c.gameService.SpinMode()

	// 取得賠率
	rate, err := c.getRate(spinMode, rtp)
	if err != nil {
		ctx.SendError(err)
		return
	}

	// 取得隨機盤面
	resultSvc := c.getResultService(spinMode)
	_, err = resultSvc.Random(rate)
	if err != nil {
		ctx.SendError(err)
		return
	}
}

func (c *controller) betForDemo(ctx *server.Context) {
	// 檢查redis是否有免費盤面List尚未消費(demo)，如有則代表當前是金蛇模式
	// results := c.gameService.RestoreResults
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
