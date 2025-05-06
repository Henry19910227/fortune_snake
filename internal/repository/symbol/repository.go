package symbol

import (
	symbolModel "game_server_slots_fortune_snake/internal/model/entity/symbol"
	"math/rand"
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

func (r *repository) GetSymbol(id int) *symbolModel.Item {
	if id < 0 || id >= len(r.symbols) {
		return nil
	}
	return r.symbols[id]
}

func (r *repository) GetRandomSymbol() *symbolModel.Item {
	rest := rand.Intn(r.totalWeight)
	for _, symbol := range r.symbols {
		if rest >= symbol.Weight {
			rest -= symbol.Weight
			continue
		}
		return symbol
	}
	return nil
}

func (r *repository) GetTotalWeight() int {
	return r.totalWeight
}

func loadSymbols() []*symbolModel.Item {
	symbols := make([]*symbolModel.Item, 0)
	names := []string{"百搭", "元宝", "金项链", "红包", "麦克风", "金币", "鞭炮"} // 圖案名稱
	weights := []int{10, 20, 30, 40, 50, 60, 70}                  // 圖案權重
	pows := []int{500, 100, 50, 20, 10, 5, 3}                     // 倍率
	for i := 0; i < len(names); i++ {
		symbol := &symbolModel.Item{
			ID:     i,
			Name:   names[i],
			Weight: weights[i],
			Pow:    pows[i],
		}
		if i == 0 {
			symbol.IsWild = true
		}
		symbols = append(symbols, symbol)
	}
	return symbols
}
