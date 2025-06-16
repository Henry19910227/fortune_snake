package service

import (
	betRecordService "game_server_slots_fortune_snake/internal/service/bet_record"
	gameService "game_server_slots_fortune_snake/internal/service/game"
	playerService "game_server_slots_fortune_snake/internal/service/player"
	reelsService "game_server_slots_fortune_snake/internal/service/reels"
	resultFreeService "game_server_slots_fortune_snake/internal/service/result_free"
	resultLoader "game_server_slots_fortune_snake/internal/service/result_loader"
	settleService "game_server_slots_fortune_snake/internal/service/settle"
	symbolService "game_server_slots_fortune_snake/internal/service/symbol"
	weightService "game_server_slots_fortune_snake/internal/service/weight"
)

type Factory interface {
	GameService() gameService.Service
	PlayerService() playerService.Service
	WeightService() weightService.Service
	ResultFreeService() resultFreeService.Service
	ResultLoader() resultLoader.Service
	ResultFreeLoader() resultLoader.Service
	ReelsService() reelsService.Service
	ReelsFreeService() reelsService.Service
	SettleService() settleService.Service
	BetRecordService() betRecordService.Service
	SymbolService() symbolService.Service
}
