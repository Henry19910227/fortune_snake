package symbol

import (
	symbolModel "game_server_slots_fortune_snake/internal/model/entity/symbol"
)

type Repository interface {
	GetSymbols() []*symbolModel.Item
	GetSymbol(id int) *symbolModel.Item
	GetTotalWeight() int
	GetRandomSymbol() *symbolModel.Item
}
