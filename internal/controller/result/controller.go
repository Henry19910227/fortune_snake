package result

import (
	"fmt"
	resultModel "game_server_slots_fortune_snake/internal/model/entity/result"
	"game_server_slots_fortune_snake/internal/model/service/result/migrate"
	"game_server_slots_fortune_snake/internal/model/service/result/save_to_bucket"
	"game_server_slots_fortune_snake/internal/model/service/settle/get_rate"
	"game_server_slots_fortune_snake/internal/model/service/settle/to_json"
	reelService "game_server_slots_fortune_snake/internal/service/reels"
	reelsFreeService "game_server_slots_fortune_snake/internal/service/reels_free"
	resultService "game_server_slots_fortune_snake/internal/service/result"
	settleService "game_server_slots_fortune_snake/internal/service/settle"
)

type controller struct {
	reelSvc     reelService.Service
	reelFreeSvc reelsFreeService.Service
	settleSvc   settleService.Service
	resultSvc   resultService.Service
}

func New(reelSvc reelService.Service, reelFreeSvc reelsFreeService.Service, settleSvc settleService.Service, resultSvc resultService.Service) Controller {
	return &controller{reelSvc: reelSvc, reelFreeSvc: reelFreeSvc, settleSvc: settleSvc, resultSvc: resultSvc}
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
		// 準備盤面結果 model
		result := &resultModel.Item{Rate: getRateOutput.GetRate(), Symbols: toJsonOutput.GetJson()}
		// 將盤面結果存進 bucket 暫存區
		saveInput := save_to_bucket.NewInput(result)
		saveOutput := c.resultSvc.SaveToBucket(saveInput)
		// 判斷 Bucket 剩餘空間是否還大於零，如大於零則繼續生產
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

func (c *controller) GenerateFree() {
	// 生成一個盤面
	reelsList := c.reelFreeSvc.Generate([]int{3, 4, 3})
	// 將盤面物件轉換為Json
	for _, reels := range reelsList {
		toJsonOutput, _ := c.settleSvc.ToJson(to_json.NewInput(reels))
		fmt.Println(toJsonOutput.GetJson())
	}
}
