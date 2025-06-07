package result_free

import (
	"fmt"
	resultModel "game_server_slots_fortune_snake/internal/model/entity/result"
	"game_server_slots_fortune_snake/internal/model/service/settle/get_rate"
	"game_server_slots_fortune_snake/internal/model/service/settle/list_to_json"
	"game_server_slots_fortune_snake/internal/model/service/settle/to_json"
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
		reels := c.reelFreeSvc.Generate([]int{3, 4, 3})
		// 獲取賠率
		getRateInput := get_rate.NewInput(get_rate.Param{Bet: 1, Value: 1000, Reels: reels[len(reels)-1]})
		getRateOutput, _ := c.settleSvc.GetRate(getRateInput)
		// 將盤面物件轉換為Json
		toJsonOutput, _ := c.settleSvc.ToJson(to_json.NewInput(reels[len(reels)-1]))
		listToJsonOutput, _ := c.settleSvc.ListToJson(list_to_json.NewInput(reels))
		// 準備盤面結果 model
		result := &resultModel.Item{Rate: getRateOutput.GetRate(), Symbols: toJsonOutput.GetJson(), AllSymbols: listToJsonOutput.GetJson()}
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
	err := c.resultFreeLoader.Migrate()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("盤面數據遷移至DB完成")
}
