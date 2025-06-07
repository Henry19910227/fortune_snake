package load

import (
	resultLoaderService "game_server_slots_fortune_snake/internal/service/result_loader"
	weightService "game_server_slots_fortune_snake/internal/service/weight"
	"game_server_slots_fortune_snake/tool"
	"go.uber.org/zap"
)

type controller struct {
	weightService    weightService.Service
	resultLoader     resultLoaderService.Service
	resultFreeLoader resultLoaderService.Service
}

func New(weightService weightService.Service, resultLoader resultLoaderService.Service, resultFreeLoader resultLoaderService.Service) Controller {
	return &controller{weightService: weightService, resultLoader: resultLoader, resultFreeLoader: resultFreeLoader}
}

func (c *controller) Load() {
	// 載入權重
	c.weightService.Load()
	// 載入盤面結果
	if err := c.resultLoader.LoadData(); err != nil {
		tool.Log().Logs("error", true, err.Error(), zap.String("Component", "Elasticsearch"), zap.Error(err))
	}
	// 載入免費盤面結果
	if err := c.resultFreeLoader.LoadData(); err != nil {
		tool.Log().Logs("error", true, err.Error(), zap.String("Component", "Elasticsearch"), zap.Error(err))
	}
}
