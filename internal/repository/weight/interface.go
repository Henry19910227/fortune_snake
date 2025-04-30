package weight

import weightModel "game_server_slots_fortune_snake/internal/model/weight"

type Repository interface {
	LoadBaseWeight()

	BaseWeight() []*weightModel.Stat  // 一般權重
	BaseWeightH() []*weightModel.Stat // 高投注權重
	FreeWeight() []*weightModel.Stat
	FreeWeightH() []*weightModel.Stat
}
