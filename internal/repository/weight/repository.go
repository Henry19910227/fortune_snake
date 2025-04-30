package weight

import (
	weightModel "game_server_slots_fortune_snake/internal/model/weight"
	"github.com/360EntSecGroup-Skylar/excelize"
	"strconv"
)

type repository struct {
	baseWeights []*weightModel.Stat
}

func New() Repository {
	return &repository{baseWeights: make([]*weightModel.Stat, 0)}
}

func (r *repository) LoadBaseWeight() {
	r.load("sheet_base")
}

func (r *repository) BaseWeight() []*weightModel.Stat {
	return r.baseWeights
}

func (r *repository) BaseWeightH() []*weightModel.Stat {
	//TODO implement me
	panic("implement me")
}

func (r *repository) FreeWeight() []*weightModel.Stat {
	//TODO implement me
	panic("implement me")
}

func (r *repository) FreeWeightH() []*weightModel.Stat {
	//TODO implement me
	panic("implement me")
}

func (r *repository) load(filename string) {
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
		r.baseWeights = append(r.baseWeights, &stat)
	}
}
