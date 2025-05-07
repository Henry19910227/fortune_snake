package result

import (
	"fmt"
	resultModel "game_server_slots_fortune_snake/internal/model/entity/result"
	"game_server_slots_fortune_snake/internal/model/service/result/save_to_bucket"
	"game_server_slots_fortune_snake/internal/model/service/settle/get_rate"
	reelService "game_server_slots_fortune_snake/internal/service/reels"
	resultService "game_server_slots_fortune_snake/internal/service/result"
	settleService "game_server_slots_fortune_snake/internal/service/settle"
)

type controller struct {
	reelSvc   reelService.Service
	settleSvc settleService.Service
	resultSvc resultService.Service
}

func New(reelSvc reelService.Service, settleSvc settleService.Service, resultSvc resultService.Service) Controller {
	return &controller{reelSvc: reelSvc, settleSvc: settleSvc, resultSvc: resultSvc}
}

func (c *controller) Generate() {
	for {
		reels := c.reelSvc.Generate([]int{3, 4, 3})
		getRateInput := &get_rate.Input{}
		getRateInput.Param = get_rate.Param{
			Bet:   1,
			Value: 1000,
			Reels: reels,
		}
		getRateOut, _ := c.settleSvc.GetRate(getRateInput)
		symbolString, err := c.settleSvc.ToJson(reels)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		item := &resultModel.Item{Rate: getRateOut.Rate, Symbols: symbolString}
		input := &save_to_bucket.Input{}
		input.Param = save_to_bucket.Param{
			Result: item,
		}
		quota := c.resultSvc.SaveToBucket(input)
		fmt.Println(quota)
		if quota > 0 {
			continue
		}
		return
	}
}

func (c *controller) SaveToDatabase() {
	//TODO implement me
	panic("implement me")
}
