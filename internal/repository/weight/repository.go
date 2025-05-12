package weight

import (
	weightModel "game_server_slots_fortune_snake/internal/model/entity/weight"
	"github.com/360EntSecGroup-Skylar/excelize"
	"strconv"
)

type repository struct {
	baseWeights  []*weightModel.Stat
	freeWeights  []*weightModel.Stat
	baseWeightsH []*weightModel.Stat
	freeWeightsH []*weightModel.Stat
}

func New() Repository {
	baseWeights := make([]*weightModel.Stat, 0)
	freeWeights := make([]*weightModel.Stat, 0)
	baseWeightsH := make([]*weightModel.Stat, 0)
	freeWeightsH := make([]*weightModel.Stat, 0)
	return &repository{baseWeights: baseWeights, freeWeights: freeWeights, baseWeightsH: baseWeightsH, freeWeightsH: freeWeightsH}
}

func (r *repository) LoadBaseWeight() {
	r.load("sheet_base", &r.baseWeights)
}

func (r *repository) LoadFreeWeight() {
	r.load("sheet_free", &r.freeWeights)
}

func (r *repository) LoadBaseWeightH() {
	r.load("sheet_base_H", &r.baseWeightsH)
}

func (r *repository) LoadFreeWeightH() {
	r.load("sheet_free_h", &r.freeWeightsH)
}

func (r *repository) BaseWeight() []*weightModel.Stat {
	return r.baseWeights
}

func (r *repository) BaseWeightH() []*weightModel.Stat {
	return r.baseWeightsH
}

func (r *repository) FreeWeight() []*weightModel.Stat {
	return r.freeWeights
}

func (r *repository) FreeWeightH() []*weightModel.Stat {
	return r.freeWeightsH
}

func (r *repository) load(filename string, list *[]*weightModel.Stat) {
	xlsx, err := excelize.OpenFile("standard_slotsnum.xlsx")
	if err != nil {
		return
	}
	sumWeights := make([]int64, 20)
	rows := xlsx.GetRows(filename)
	for row := 1; row < len(rows); row++ {
		fromWeights := make([]int64, 0)
		toWeights := make([]int64, 0)
		weightAccount := make([]int64, 0)
		var rate float64
		for col := 0; col < len(rows[row]); col++ {
			value := rows[row][col]
			if col == 0 {
				rate, _ = strconv.ParseFloat(value, 64)
				continue
			}
			weight, _ := strconv.ParseFloat(value, 64)
			weightAccount = append(weightAccount, int64(weight))
			fromWeights = append(fromWeights, sumWeights[col])
			sumWeights[col] += int64(weight)
			toWeights = append(toWeights, sumWeights[col])
		}
		stat := weightModel.Stat{Rate: rate, FromWeights: fromWeights, WeightAccount: weightAccount, ToWeights: toWeights}
		*list = append(*list, &stat)
	}
}
