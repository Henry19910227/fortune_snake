package game

import (
	"game_server_slots_fortune_snake/constants"
	"game_server_slots_fortune_snake/internal/model/game/enter_game"
	gameRepo "game_server_slots_fortune_snake/internal/repository/game"
	playerRepo "game_server_slots_fortune_snake/internal/repository/player"
)

// 真錢模式的 game service
type service struct {
	gameRepo   gameRepo.Repository
	playerRepo playerRepo.Repository
}

func NewService(gameRepo gameRepo.Repository, playerRepo playerRepo.Repository) Service {
	return &service{gameRepo: gameRepo, playerRepo: playerRepo}
}

func (s *service) EnterGame(input *enter_game.Input) (output *enter_game.Output) {
	// 獲取遊戲配置
	gameInfo, err := s.gameRepo.Info()
	if err != nil {
		output.Code = constants.CodeBadRequest
		output.Message = err.Error()
		return output
	}
	// 獲取玩家遊戲緩存數據
	playerGameData, err := s.playerRepo.GameData()
	if err != nil {
		output.Code = constants.CodeBadRequest
		output.Message = err.Error()
		return output
	}
	// 處理回傳
	output = &enter_game.Output{}
	output.Data = &enter_game.Data{
		Bets:               gameInfo.Bets,
		Values:             gameInfo.Values,
		Bet:                playerGameData.Bet,
		Value:              playerGameData.Value,
		GameMode:           input.Session.Mode,
		Multipler:          gameInfo.Multipler,
		MultipleScoreLimit: gameInfo.MultipleScoreLimit,
		ScoreTry:           0, // 真實遊玩回傳 0
	}
	return output
}
