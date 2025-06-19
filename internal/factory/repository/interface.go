package repository

import (
	betRecordRepository "game_server_slots_fortune_snake/internal/repository/bet_record"
	bucketRepo "game_server_slots_fortune_snake/internal/repository/bucket"
	gameRepo "game_server_slots_fortune_snake/internal/repository/game"
	gameResultRepo "game_server_slots_fortune_snake/internal/repository/game_result"
	playerRepo "game_server_slots_fortune_snake/internal/repository/player"
	freeOrder "game_server_slots_fortune_snake/internal/repository/player_free_order"
	session "game_server_slots_fortune_snake/internal/repository/player_session"
	resultFreeRepo "game_server_slots_fortune_snake/internal/repository/result_free"
	resultLoader "game_server_slots_fortune_snake/internal/repository/result_loader"
	settleRepository "game_server_slots_fortune_snake/internal/repository/settle"
	snowFlakeRepository "game_server_slots_fortune_snake/internal/repository/snow_flake"
	symbolRepo "game_server_slots_fortune_snake/internal/repository/symbol"
	weightRepo "game_server_slots_fortune_snake/internal/repository/weight"
)

type Factory interface {
	GameRepository() gameRepo.Repository
	PlayerRepository() playerRepo.Repository
	SessionRepository() session.Repository
	WeightRepository() weightRepo.Repository
	ResultLoader() resultLoader.Repository
	ResultFreeLoader() resultLoader.Repository
	ResultFreeRepository() resultFreeRepo.Repository
	BucketRepository() bucketRepo.Repository
	BucketFreeRepository() bucketRepo.Repository
	SymbolRepository() symbolRepo.Repository
	SettleRepository() settleRepository.Repository
	SnowflakeRepository() snowFlakeRepository.Repository
	BetRecordRepository() betRecordRepository.Repository
	GameResultRepository() gameResultRepo.Repository
	FreeOrderRepository() freeOrder.Repository
}
