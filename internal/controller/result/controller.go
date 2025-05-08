package result

import (
	"fmt"
	resultModel "game_server_slots_fortune_snake/internal/model/entity/result"
	"game_server_slots_fortune_snake/internal/model/service/result/migrate"
	"game_server_slots_fortune_snake/internal/model/service/result/save_to_bucket"
	"game_server_slots_fortune_snake/internal/model/service/settle/get_rate"
	"game_server_slots_fortune_snake/internal/model/service/settle/to_json"
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
	fmt.Println("開始生成盤面數據至暫存區")
	for {
		// 生成一個盤面
		reels := c.reelSvc.Generate([]int{3, 4, 3})
		// 獲取賠率
		getRateInput := get_rate.NewInput(get_rate.Param{Bet: 1, Value: 1000, Reels: reels})
		getRateOutput, _ := c.settleSvc.GetRate(getRateInput)
		// 將盤面物件轉換為Json
		toJsonInput := to_json.NewInput(reels)
		toJsonOutput, _ := c.settleSvc.ToJson(toJsonInput)
		// 盤面結果 model
		result := &resultModel.Item{Rate: getRateOutput.GetRate(), Symbols: toJsonOutput.GetJson()}
		// 將盤面結果存進 bucket 暫存區
		saveInput := save_to_bucket.NewInput(result)
		saveOutput := c.resultSvc.SaveToBucket(saveInput)
		if saveOutput.GetQuota() > 0 {
			fmt.Println(saveOutput.GetQuota())
			continue
		}
		fmt.Println("盤面數據生成完成，並存入暫存區")
		break
	}
	fmt.Println("開始將暫存區數據遷移至DB")
	err := c.resultSvc.Migrate(&migrate.Input{})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("盤面數據遷移至DB完成")
}

func (c *controller) SaveToDatabase() {
	//TODO implement me
	panic("implement me")
}
