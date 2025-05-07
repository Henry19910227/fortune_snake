package player

import (
	"context"
	"game_server_slots_fortune_snake/constants"
	"game_server_slots_fortune_snake/internal/model"
	"game_server_slots_fortune_snake/internal/model/player/get_player_session"
	"game_server_slots_fortune_snake/internal/server"
	playerService "game_server_slots_fortune_snake/internal/service/player"
)

type controller struct {
	playerService playerService.Service
}

func New(playerService playerService.Service) Controller {
	return &controller{playerService: playerService}
}

func (c *controller) GetPlayerSession(ctx *server.Context) {
	req := ctx.MustGet("req").(*model.MessageRequest)
	grpcCtx := ctx.MustGet("ctx").(context.Context)
	// 查詢 Player Session
	input := &get_player_session.Input{}
	input.Ctx = grpcCtx
	input.PlayerId = req.PlayerId
	output, err := c.playerService.GetPlayerSession(input)
	if err != nil {
		ctx.SendError(err)
		return
	}
	// 返回結果
	ctx.Send(constants.CodeSuccess, "success", output.Session)
}
