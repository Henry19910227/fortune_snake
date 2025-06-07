package repository

import (
	bucketRepo "game_server_slots_fortune_snake/internal/repository/bucket"
	gameRepo "game_server_slots_fortune_snake/internal/repository/game"
	playerRepo "game_server_slots_fortune_snake/internal/repository/player"
	resultRepo "game_server_slots_fortune_snake/internal/repository/result"
	weightRepo "game_server_slots_fortune_snake/internal/repository/weight"
)

type Factory interface {
	GameRepository() gameRepo.Repository
	PlayerRepository() playerRepo.Repository
	WeightRepository() weightRepo.Repository
	ResultRepository() resultRepo.Repository
	ResultFreeRepository() resultRepo.Repository
	BucketRepository() bucketRepo.Repository
	BucketFreeRepository() bucketRepo.Repository
}
