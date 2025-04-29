package game

import (
	"context"
	"game_server_slots_fortune_snake/constants"
	"game_server_slots_fortune_snake/internal/model/game/bet"
	"game_server_slots_fortune_snake/internal/model/game/enter_game"
	playerModel "game_server_slots_fortune_snake/internal/model/player"
	"game_server_slots_fortune_snake/internal/server"
	gameService "game_server_slots_fortune_snake/internal/service/game"
)

type controller struct {
	gameService     gameService.Service // 真實模式 service
	gameDemoService gameService.Service // 試玩模式 service
}

func New(gameService gameService.Service) Controller {
	return &controller{gameService: gameService}
}

func (c *controller) EnterGame(ctx *server.Context) {
	// 取得中間層處理好的數據
	grpcCtx := ctx.MustGet("ctx").(context.Context)
	session := ctx.MustGet("session").(*playerModel.Session)
	// 以 mode 獲取對應的 game service
	service := c.getGameService(session.Mode)
	// 執行 Enter Game 業務邏輯
	input := &enter_game.Input{}
	input.Ctx = grpcCtx
	input.Session = session
	output := service.EnterGame(input)
	// 返回結果
	ctx.Send(output.Code, output.Message, output.Data)
}

func (c *controller) Bet(ctx *server.Context) {
	// 取得請求數據
	grpcCtx := ctx.MustGet("ctx").(context.Context)
	session := ctx.MustGet("session").(*playerModel.Session)
	// 獲取參數
	var param bet.Param
	if err := ctx.Bind(&param); err != nil {
		ctx.SendError(constants.CodeBadRequest, err.Error())
		ctx.Abort()
		return
	}
	// 處理輸入參數
	input := &bet.Input{}
	input.Ctx = grpcCtx
	input.Session = session
	input.Param = &param
	// 以 mode 獲取對應的 game service
	service := c.getGameService(session.Mode)
	// 執行 Bet 業務邏輯
	output := service.Bet(input)
	ctx.Send(output.Code, output.Message, output.Data)
}

func (c *controller) getGameService(mode string) gameService.Service {
	if mode == "demo" {
		return c.gameDemoService
	} // 試玩模式業務層
	return c.gameService // 真實模式業務層
}
