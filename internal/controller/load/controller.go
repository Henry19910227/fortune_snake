package load

import (
	resultService "game_server_slots_fortune_snake/internal/service/result"
	weightService "game_server_slots_fortune_snake/internal/service/weight"
	"game_server_slots_fortune_snake/tool"
	"go.uber.org/zap"
)

type controller struct {
	weightService     weightService.Service
	resultService     resultService.Service
	resultFreeService resultService.Service
}

func New(weightService weightService.Service, resultService resultService.Service, resultFreeService resultService.Service) Controller {
	return &controller{weightService: weightService, resultService: resultService, resultFreeService: resultFreeService}
}

func (c *controller) Load() {
	// 載入權重
	c.weightService.Load()
	// 載入盤面結果
	if err := c.resultService.LoadData(); err != nil {
		tool.Log().Logs("error", true, err.Error(), zap.String("Component", "Elasticsearch"), zap.Error(err))
	}
	// 載入免費盤面結果
	if err := c.resultFreeService.LoadData(); err != nil {
		tool.Log().Logs("error", true, err.Error(), zap.String("Component", "Elasticsearch"), zap.Error(err))
	}
}
