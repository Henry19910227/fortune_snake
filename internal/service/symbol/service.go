package symbol

import (
	"game_server_slots_fortune_snake/internal/model/entity/symbol"
	symbolRepository "game_server_slots_fortune_snake/internal/repository/symbol"
)

type service struct {
	symbolRepo symbolRepository.Repository
}

func New(symbolRepo symbolRepository.Repository) Service {
	return &service{symbolRepo: symbolRepo}
}

func (s *service) GetSymbols() []symbol.Item {
	symbols := make([]symbol.Item, 0)
	items := s.symbolRepo.GetSymbols()
	for _, item := range items {
		if item.ID == 99 {
			continue
		}
		symbols = append(symbols, item)
	}
	return symbols
}
