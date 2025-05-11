package result_free

import (
	"fmt"
	resultModel "game_server_slots_fortune_snake/internal/model/entity/result_free"
	"game_server_slots_fortune_snake/internal/model/service/result_free/save_to_bucket"
	"game_server_slots_fortune_snake/internal/model/service/settle/get_rate"
	"game_server_slots_fortune_snake/internal/model/service/settle/list_to_json"
	"game_server_slots_fortune_snake/internal/model/service/settle/to_json"
	reelsFreeService "game_server_slots_fortune_snake/internal/service/reels_free"
	resultFreeService "game_server_slots_fortune_snake/internal/service/result_free"
	settleService "game_server_slots_fortune_snake/internal/service/settle"
)

type controller struct {
	reelFreeSvc   reelsFreeService.Service
	settleSvc     settleService.Service
	resultFreeSvc resultFreeService.Service
}

func New(reelFreeSvc reelsFreeService.Service, settleSvc settleService.Service, resultFreeSvc resultFreeService.Service) Controller {
	return &controller{reelFreeSvc: reelFreeSvc, settleSvc: settleSvc, resultFreeSvc: resultFreeSvc}
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
		saveInput := save_to_bucket.NewInput(result)
		saveOutput := c.resultFreeSvc.SaveToBucket(saveInput)
		// 判斷 Bucket 剩餘空間是否還大於零，如大於零則繼續生產
		if saveOutput.GetQuota() > 0 {
			fmt.Println(saveOutput.GetQuota())
			continue
		}
		fmt.Println("盤面數據生成完成，並存入暫存區")
		break
	}
	fmt.Println("開始將暫存區數據遷移至DB")
	err := c.resultFreeSvc.Migrate()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("盤面數據遷移至DB完成")
}
