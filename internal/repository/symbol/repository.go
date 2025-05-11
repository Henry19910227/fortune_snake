package symbol

import (
	cfg "game_server_slots_fortune_snake/internal/model/config/symbol"
	symbolModel "game_server_slots_fortune_snake/internal/model/entity/symbol"
	"math/rand"
)

// Repository Symbol 資源的存取與管理介面
type repository struct {
	symbols     []symbolModel.Item
	symbolMap   map[int]symbolModel.Item
	totalWeight int
}

func New(config []*cfg.Item) Repository {
	symbols := loadSymbols(config)
	symbolMap := loadSymbolMap(config)
	totalWeight := 0
	for _, symbol := range symbols {
		if symbol.ID == 99 {
			continue
		}
		totalWeight += symbol.Weight
	}
	return &repository{symbols: symbols, symbolMap: symbolMap, totalWeight: totalWeight}
}

// GetSymbols 返回所有定義的 symbol 物件
func (r *repository) GetSymbols() []symbolModel.Item {
	items := make([]symbolModel.Item, len(r.symbols))
	copy(items, r.symbols)
	return items
}

// GetSymbol 以 symbol id 獲取 symbol 物件
func (r *repository) GetSymbol(id int) *symbolModel.Item {
	symbol, ok := r.symbolMap[id]
	if !ok {
		return nil
	}
	return &symbol
}

// GetRandomSymbol 獲取一個隨機的 symbol
func (r *repository) GetRandomSymbol() *symbolModel.Item {
	rest := rand.Intn(r.totalWeight)
	for _, symbol := range r.symbols {
		if rest >= symbol.Weight {
			rest -= symbol.Weight
			continue
		}
		return &symbol
	}
	return nil
}

// GetTotalWeight 獲取所有 symbol 權重總和
func (r *repository) GetTotalWeight() int {
	return r.totalWeight
}

// loadSymbols Symbol 載入
func loadSymbols(config []*cfg.Item) []symbolModel.Item {
	symbols := make([]symbolModel.Item, 0)
	for _, item := range config {
		if item.ID == 99 {
			continue
		}
		symbol := symbolModel.Item{
			ID:     item.ID,
			Name:   item.Name,
			Weight: item.Weight,
			Pow:    item.Pow,
			IsWild: item.IsWild,
		}
		symbols = append(symbols, symbol)
	}
	return symbols
}

func loadSymbolMap(config []*cfg.Item) map[int]symbolModel.Item {
	symbolMap := make(map[int]symbolModel.Item)
	for _, item := range config {
		symbol := symbolModel.Item{
			ID:     item.ID,
			Name:   item.Name,
			Weight: item.Weight,
			Pow:    item.Pow,
			IsWild: item.IsWild,
		}
		symbolMap[item.ID] = symbol
	}
	return symbolMap
}
