package middleware

import (
	"context"
	"encoding/json"
	"game_server_slots_fortune_snake/constants"
	"game_server_slots_fortune_snake/model"
	"game_server_slots_fortune_snake/model/player/get_player_session"
	"game_server_slots_fortune_snake/server"
	playerService "game_server_slots_fortune_snake/service/player"
)

type controller struct {
	playerService playerService.Service
}

func New(playerService playerService.Service) Controller {
	return &controller{playerService: playerService}
}

func (c *controller) UnMarshalData(ctx *server.Context) {
	var req *model.MessageRequest
	if err := json.Unmarshal(ctx.Data(), &req); err != nil {
		ctx.Abort()
		return
	}
	ctx.Set("req", req)
}

func (c *controller) UnMarshalReq(ctx *server.Context) {
	//req := ctx.MustGet("req").(*model.MessageRequest)
	//switch {
	//case req.Action == "bet":
	//default:
	//	ctx.Abort()
	//	return
	//}
	//TODO implement me
	panic("implement me")
}

func (c *controller) Verify(ctx *server.Context) {
	req := ctx.MustGet("req").(*model.MessageRequest)
	grpcCtx := ctx.MustGet("ctx").(context.Context)

	input := get_player_session.Input{}
	input.Ctx = grpcCtx
	input.PlayerId = req.PlayerId
	output := c.playerService.GetPlayerSession(&input)
	if output.Code != constants.CodeSuccess {
		ctx.Abort()
		return
	}
	ctx.Set("player", output.Data)
}
