package symbol

import symbolModel "game_server_slots_fortune_snake/internal/model/symbol"

type Repository interface {
	GetSymbols() []*symbolModel.Item
	GetTotalWeight() int
}
