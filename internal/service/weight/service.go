package weight

import (
	weightRepo "game_server_slots_fortune_snake/internal/repository/weight"
)

type service struct {
	weightRepo weightRepo.Repository
}

func New(weightRepo weightRepo.Repository) Service {
	return &service{weightRepo: weightRepo}
}

func (s *service) Load() {
	s.weightRepo.LoadFreeWeight()
	s.weightRepo.LoadBaseWeight()
	s.weightRepo.LoadBaseWeightH()
	s.weightRepo.LoadFreeWeightH()
}

func (s *service) GetRandomRate(rtp int) {
	
}
