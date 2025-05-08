package middleware

import (
	"context"
	"errors"
	"fmt"
	"game_server_slots_fortune_snake/constants"
	errMsg "game_server_slots_fortune_snake/internal/model/err"
	"game_server_slots_fortune_snake/internal/model/service/player/get_player_session"
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
	output, err := c.playerService.GetPlayerSession(&input)
	if err != nil {
		ctx.SendError(err)
		ctx.Abort()
		return
	}
	ctx.Set("session", output.Session)
}

// Recover panic 恢復層
func (c *controller) Recover(ctx *server.Context) {
	defer func() {
		if r := recover(); r != nil {
			var err *errMsg.Error
			switch e := r.(type) {
			case string:
				err = errMsg.New(constants.CodeInternalError, e, errors.New(e))
			case error:
				err = errMsg.New(constants.CodeInternalError, e.Error(), err)
			default:
				errorMsg := fmt.Errorf("未知錯誤: %v", e)
				err = errMsg.New(constants.CodeInternalError, errorMsg.Error(), errorMsg)
			}
			ctx.SendError(err)
			ctx.Abort()
		}
	}()
	ctx.Next()
}
