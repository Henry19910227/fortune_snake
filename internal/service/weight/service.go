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

func (s *service) RandomBaseWeightRate(rtp float64) (float64, error) {
	return s.weightRepo.RandomBaseWeightRate(rtp)
}

func (s *service) RandomBaseWeightHRate(rtp float64) (float64, error) {
	return s.weightRepo.RandomBaseWeightHRate(rtp)
}

func (s *service) RandomFreeWeightRate(rtp float64) (float64, error) {
	return s.weightRepo.RandomFreeWeightRate(rtp)
}

func (s *service) RandomFreeWeightHRate(rtp float64) (float64, error) {
	return s.weightRepo.RandomFreeWeightHRate(rtp)
}
