package game

import (
	gameModel "game_server_slots_fortune_snake/internal/model/service/game"
)

type repository struct {
}

func New() Repository {
	return &repository{}
}

func (r *repository) Info() (output *gameModel.Info, err error) {
	return &gameModel.Info{
		Bets:               []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		Values:             []float64{100, 1000, 4000, 20000},
		Multipler:          10,
		MultipleScoreLimit: 2000,
	}, nil
}

func (r *repository) IsSpecialMode() bool {
	return false
}
