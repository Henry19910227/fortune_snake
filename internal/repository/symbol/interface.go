package symbol

import (
	symbolModel "game_server_slots_fortune_snake/internal/model/entity/symbol"
)

// Repository Symbol 資源的存取與管理介面
type Repository interface {
	// GetSymbols 返回所有定義的 symbol 物件
	GetSymbols() []*symbolModel.Item
	// GetSymbol 以 symbol id 獲取 symbol 物件
	GetSymbol(id int) *symbolModel.Item
	// GetTotalWeight 獲取所有 symbol 權重總和
	GetTotalWeight() int
	// GetRandomSymbol 獲取一個隨機的 symbol
	GetRandomSymbol() *symbolModel.Item
}
