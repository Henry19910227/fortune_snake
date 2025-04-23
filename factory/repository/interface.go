package repository

import playerRepo "game_server_slots_fortune_snake/repository/player"

type Factory interface {
	PlayerRepository() playerRepo.Repository
}
