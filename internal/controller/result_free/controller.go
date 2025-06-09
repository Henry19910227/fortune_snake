package result_free

import (
	"fmt"
	. "game_server_slots_fortune_snake/constants"
	resultModel "game_server_slots_fortune_snake/internal/model/entity/result"
	reelsService "game_server_slots_fortune_snake/internal/service/reels_free"
	resultLoaderService "game_server_slots_fortune_snake/internal/service/result_loader"
	settleService "game_server_slots_fortune_snake/internal/service/settle"
)

type controller struct {
	reelFreeSvc      reelsService.Service
	settleSvc        settleService.Service
	resultFreeLoader resultLoaderService.Service
}

func New(reelFreeSvc reelsService.Service, settleSvc settleService.Service, resultFreeLoader resultLoaderService.Service) Controller {
	return &controller{reelFreeSvc: reelFreeSvc, settleSvc: settleSvc, resultFreeLoader: resultFreeLoader}
}

func (c *controller) Generate() {
	fmt.Println("開始生成盤面數據至暫存區")
	for {
		// 生成一個免費模式盤面
		reelsList := c.reelFreeSvc.Generate([]int{3, 4, 3})
		// 獲取最後一個盤面的賠率
		rate, _ := c.settleSvc.GetRate(1, 1000, reelsList[len(reelsList)-1])
		// 將最後一個盤面轉換為Json
		reelsString, _ := c.settleSvc.ToJson(reelsList[len(reelsList)-1])
		// 將所有盤面轉換為Json
		reelsListString, _ := c.settleSvc.ListToJson(reelsList)
		// 準備盤面結果 model
		result := &resultModel.Item{Rate: rate, Symbols: reelsString, AllSymbols: reelsListString}
		// 將盤面結果存進 bucket 暫存區
		quota := c.resultFreeLoader.SaveToBucket(result)
		// 判斷 Bucket 剩餘空間是否還大於零，如大於零則繼續生產
		if quota > 0 {
			fmt.Println(quota)
			continue
		}
		fmt.Println("盤面數據生成完成，並存入暫存區")
		break
	}
	fmt.Println("開始將暫存區數據遷移至DB")
	err := c.resultFreeLoader.SpinMode(SpinModeFree).Migrate()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("盤面數據遷移至DB完成")
}
