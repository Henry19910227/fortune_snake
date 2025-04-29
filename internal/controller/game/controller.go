package game

import (
	"context"
	"game_server_slots_fortune_snake/internal/model"
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
	_ = ctx.MustGet("req").(*model.MessageRequest)
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

func (c *controller) getGameService(mode string) gameService.Service {
	if mode == "demo" {
		return c.gameDemoService
	}
	return c.gameService
}
