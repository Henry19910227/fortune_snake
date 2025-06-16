package bet_record

import (
	"context"
	. "game_server_slots_fortune_snake/constants"
	betRecordModel "game_server_slots_fortune_snake/internal/model/entity/bet_record"
	playerModel "game_server_slots_fortune_snake/internal/model/entity/player"
	"game_server_slots_fortune_snake/internal/server"
	betRecordService "game_server_slots_fortune_snake/internal/service/bet_record"
	gameService "game_server_slots_fortune_snake/internal/service/game"
)

type controller struct {
	betRecordService betRecordService.Service
	gameService      gameService.Service
}

func New(betRecordService betRecordService.Service, gameService gameService.Service) Controller {
	return &controller{betRecordService: betRecordService, gameService: gameService}
}

func (c *controller) CreateRecord(ctx *server.Context) {
	session := ctx.MustGet("session").(*playerModel.Session)
	condition := ctx.MustGet("condition").(string)
	grpcCtx := ctx.MustGet("ctx").(context.Context)

	if condition == FreeModeInReal || condition == FinalFreeModeInReal || condition == FreeModeInDemo || condition == FinalFreeModeInDemo {
		fatherID, err := c.gameService.GameMode(session.Mode).GetFatherID(grpcCtx, session.PlayerId)
		if err != nil {
			ctx.SendError(err)
			return
		}
		record, err := c.betRecordService.CreateByTransactionID(session, fatherID)
		if err != nil {
			ctx.SendError(err)
			return
		}
		ctx.Set("record", record)
		return
	}

	record, err := c.betRecordService.Create(session)
	if err != nil {
		ctx.SendError(err)
		return
	}
	ctx.Set("record", record)
}

func (c *controller) UpdateRecord(ctx *server.Context) {
	record := ctx.MustGet("record").(*betRecordModel.Table)
	if err := c.betRecordService.UpdateToFinished(record); err != nil {
		ctx.SendError(err)
	}
}
