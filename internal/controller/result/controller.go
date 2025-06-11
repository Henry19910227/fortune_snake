package result

import (
	"fmt"
	resultModel "game_server_slots_fortune_snake/internal/model/entity/result"
	reelService "game_server_slots_fortune_snake/internal/service/reels"
	reelsFreeService "game_server_slots_fortune_snake/internal/service/reels_free"
	resultLoaderService "game_server_slots_fortune_snake/internal/service/result_loader"
	settleService "game_server_slots_fortune_snake/internal/service/settle"
)

type controller struct {
	reelsSvc     reelService.Service
	reelFreeSvc  reelsFreeService.Service
	settleSvc    settleService.Service
	resultLoader resultLoaderService.Service
}

func New(reelSvc reelService.Service, settleSvc settleService.Service, resultLoader resultLoaderService.Service) Controller {
	return &controller{reelsSvc: reelSvc, settleSvc: settleSvc, resultLoader: resultLoader}
}

func (c *controller) Generate() {
	fmt.Println("開始生成盤面數據至暫存區")
	for {
		// 生成一個盤面
		reelsList := c.reelsSvc.Generate([]int{3, 4, 3})
		reels := reelsList[0]
		// 獲取賠率
		rate, _ := c.settleSvc.GetRate(1, 1000, reels)
		// 將盤面物件轉換為Json
		JsonString, _ := c.settleSvc.ToJson(reels)
		// 準備盤面結果 model
		result := &resultModel.Item{Rate: rate, Symbols: JsonString}
		// 將盤面結果存進 bucket 暫存區
		quota := c.resultLoader.SaveToBucket(result)
		// 判斷 Bucket 剩餘空間是否還大於零，如大於零則繼續生產
		if quota > 0 {
			fmt.Println(quota)
			continue
		}
		fmt.Println("盤面數據生成完成，並存入暫存區")
		break
	}
	fmt.Println("開始將暫存區數據遷移至DB")
	err := c.resultLoader.Migrate()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("盤面數據遷移至DB完成")
}
