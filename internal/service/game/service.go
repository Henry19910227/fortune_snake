package game

import (
	"game_server_slots_fortune_snake/internal/model/game/enter_game"
	gameRepo "game_server_slots_fortune_snake/internal/repository/game"
)

type service struct {
	gameRepo gameRepo.Repository
}

func NewService(gameRepo gameRepo.Repository) Service {
	return &service{gameRepo: gameRepo}
}

func (s *service) EnterGame(input *enter_game.Input) (output *enter_game.Output) {
	//TODO implement me
	panic("implement me")
}
