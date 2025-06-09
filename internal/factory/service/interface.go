package service

import (
	gameService "game_server_slots_fortune_snake/internal/service/game"
	playerService "game_server_slots_fortune_snake/internal/service/player"
	reelsService "game_server_slots_fortune_snake/internal/service/reels"
	reelsFreeService "game_server_slots_fortune_snake/internal/service/reels_free"
	resultFreeService "game_server_slots_fortune_snake/internal/service/result_free"
	resultLoader "game_server_slots_fortune_snake/internal/service/result_loader"
	settleService "game_server_slots_fortune_snake/internal/service/settle"
	weightService "game_server_slots_fortune_snake/internal/service/weight"
)

type Factory interface {
	GameService() gameService.Service
	GameDemoService() gameService.Service
	PlayerService() playerService.Service
	WeightService() weightService.Service
	ResultFreeService() resultFreeService.Service
	ResultLoader() resultLoader.Service
	ResultFreeLoader() resultLoader.Service
	ReelsService() reelsService.Service
	ReelsFreeService() reelsFreeService.Service
	SettleService() settleService.Service
}
