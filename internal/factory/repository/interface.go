package repository

import (
	gameRepo "game_server_slots_fortune_snake/internal/repository/game"
	playerRepo "game_server_slots_fortune_snake/internal/repository/player"
	weightRepo "game_server_slots_fortune_snake/internal/repository/weight"
)

type Factory interface {
	GameRepository() gameRepo.Repository
	PlayerRepository() playerRepo.Repository
	WeightRepository() weightRepo.Repository
}
