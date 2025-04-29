package middleware

import (
	"context"
	"game_server_slots_fortune_snake/constants"
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

// Verify 第三層 驗證 PlayerID 並獲取 Player Session 資訊
func (c *controller) Verify(ctx *server.Context) {
	playerId := ctx.MustGet("playerId").(uint64)
	grpcCtx := ctx.MustGet("ctx").(context.Context)

	input := get_player_session.Input{}
	input.Ctx = grpcCtx
	input.PlayerId = playerId
	output := c.playerService.GetPlayerSession(&input)
	if output.Code != constants.CodeSuccess {
		ctx.SendError(constants.CodeBadRequest, output.Message)
		ctx.Abort()
		return
	}
	ctx.Set("session", output.Data)
}
