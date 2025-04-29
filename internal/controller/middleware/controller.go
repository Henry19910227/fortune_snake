package middleware

import (
	"context"
	"encoding/json"
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

// UnMarshalData 第一層 將整包 json 字串轉為 MessageRequest model
func (c *controller) UnMarshalData(ctx *server.Context) {
	var req *model.MessageRequest
	if err := json.Unmarshal(ctx.Data(), &req); err != nil {
		ctx.SendError(constants.CodeBadRequest, err.Error())
		ctx.Abort()
		return
	}
	ctx.Set("req", req)
}

// UnMarshalReq 第二層 將 MessageRequest model 的 data json 字段轉換為對應 model
func (c *controller) UnMarshalReq(ctx *server.Context) {
	req := ctx.MustGet("req").(*model.MessageRequest)
	switch {
	case req.Action == "player":
		req.Data = ""
	case req.Action == "enter_game":
		req.Data = ""
	default:
		ctx.SendError(constants.CodeBadRequest, "action is not valid")
		ctx.Abort()
		return
	}
	ctx.Set("req", req)
}

// Verify 第三層 驗證 PlayerID 並獲取 Player Session 資訊
func (c *controller) Verify(ctx *server.Context) {
	req := ctx.MustGet("req").(*model.MessageRequest)
	grpcCtx := ctx.MustGet("ctx").(context.Context)

	input := get_player_session.Input{}
	input.Ctx = grpcCtx
	input.PlayerId = req.PlayerId
	output := c.playerService.GetPlayerSession(&input)
	if output.Code != constants.CodeSuccess {
		ctx.SendError(constants.CodeBadRequest, output.Message)
		ctx.Abort()
		return
	}
	ctx.Set("session", output.Data)
}
