package symbol

import (
	symbolModel "game_server_slots_fortune_snake/internal/model/symbol"
)

type repository struct {
	symbols     []*symbolModel.Item
	totalWeight int
}

func New() Repository {
	symbols := loadSymbols()
	totalWeight := 0
	for _, symbol := range symbols {
		totalWeight += symbol.Weight
	}
	return &repository{symbols: symbols, totalWeight: totalWeight}
}

func (r *repository) GetSymbols() []*symbolModel.Item {
	return r.symbols
}

func (r *repository) GetTotalWeight() int {
	return r.totalWeight
}

func loadSymbols() []*symbolModel.Item {
	symbols := make([]*symbolModel.Item, 0)
	names := []string{"百搭", "元宝", "福箱", "福袋", "红包", "橘子", "鞭炮"} // 圖案名稱
	weights := []int{32, 45, 60, 70, 70, 75, 75}                              // 圖案權重
	pows := []int{200, 100, 50, 20, 10, 5, 3}                                 // 倍率
	for i := 0; i < len(names); i++ {
		symbol := &symbolModel.Item{
			ID:     i,
			Name:   names[i],
			Weight: weights[i],
			Pow:    pows[i],
		}
		symbols = append(symbols, symbol)
	}
	return symbols
}
