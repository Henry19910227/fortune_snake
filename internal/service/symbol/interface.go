package symbol

import "game_server_slots_fortune_snake/internal/model/entity/symbol"

type Service interface {
	GetSymbols() []symbol.Item
}
