package result

import (
	reelService "game_server_slots_fortune_snake/internal/service/reels"
	settleService "game_server_slots_fortune_snake/internal/service/settle"
)

type controller struct {
	reelSvc   reelService.Service
	settleSvc settleService.Service
}

func New(reelSvc reelService.Service, settleSvc settleService.Service) Controller {
	return &controller{reelSvc: reelSvc, settleSvc: settleSvc}
}

func (c *controller) Generate() {}
