package bet_record

import (
	betRecordModel "game_server_slots_fortune_snake/internal/model/entity/bet_record"
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

func (c *controller) UpdateRecord(ctx *server.Context) {
	record := ctx.MustGet("record").(*betRecordModel.Table)
	if err := c.betRecordService.UpdateToFinished(record); err != nil {
		ctx.SendError(err)
	}
}
