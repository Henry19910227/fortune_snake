package repository

import (
	bucketRepo "game_server_slots_fortune_snake/internal/repository/bucket"
	gameRepo "game_server_slots_fortune_snake/internal/repository/game"
	playerRepo "game_server_slots_fortune_snake/internal/repository/player"
	resultFreeRepo "game_server_slots_fortune_snake/internal/repository/result_free"
	resultLoader "game_server_slots_fortune_snake/internal/repository/result_loader"
	weightRepo "game_server_slots_fortune_snake/internal/repository/weight"
)

type Factory interface {
	GameRepository() gameRepo.Repository
	PlayerRepository() playerRepo.Repository
	WeightRepository() weightRepo.Repository
	ResultLoader() resultLoader.Repository
	ResultFreeLoader() resultLoader.Repository
	ResultFreeRepository() resultFreeRepo.Repository
	BucketRepository() bucketRepo.Repository
	BucketFreeRepository() bucketRepo.Repository
}
