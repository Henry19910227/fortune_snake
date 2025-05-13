package weight

import (
	weightModel "game_server_slots_fortune_snake/internal/model/entity/weight"
)

type Repository interface {
	LoadBaseWeight()
	LoadFreeWeight()
	LoadBaseWeightH()
	LoadFreeWeightH()

	BaseWeight() []*weightModel.Stat  // 一般權重
	BaseWeightH() []*weightModel.Stat // 高投注權重
	FreeWeight() []*weightModel.Stat
	FreeWeightH() []*weightModel.Stat

	RandomBaseWeightRate(rtp float64) (float64, error)
	RandomBaseWeightHRate(rtp float64) (float64, error)
	RandomFreeWeightRate(rtp float64) (float64, error)
	RandomFreeWeightHRate(rtp float64) (float64, error)
}
