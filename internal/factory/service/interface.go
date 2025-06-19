package service

import (
	betRecordService "game_server_slots_fortune_snake/internal/service/bet_record"
	gameService "game_server_slots_fortune_snake/internal/service/game"
	gameResultService "game_server_slots_fortune_snake/internal/service/game_result"
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
	ResultFreeRealService() resultFreeService.Service
	ResultFreeDemoService() resultFreeService.Service
	ResultLoader() resultLoader.Service
	ResultFreeLoader() resultLoader.Service
	ReelsService() reelsService.Service
	ReelsFreeService() reelsService.Service
	SettleService() settleService.Service
	BetRecordService() betRecordService.Service
	GameResultService() gameResultService.Service
	SymbolService() symbolService.Service
}
