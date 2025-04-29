package game

import gameModel "game_server_slots_fortune_snake/internal/model/game"

type repository struct {
}

func New() Repository {
	return &repository{}
}

func (r *repository) Info() (output *gameModel.Info, err error) {
	return &gameModel.Info{}, nil
}
