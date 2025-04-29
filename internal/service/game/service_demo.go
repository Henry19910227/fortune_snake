package game

import (
	"game_server_slots_fortune_snake/constants"
	"game_server_slots_fortune_snake/internal/model/game/bet"
	"game_server_slots_fortune_snake/internal/model/game/enter_game"
	gameRepo "game_server_slots_fortune_snake/internal/repository/game"
	playerRepo "game_server_slots_fortune_snake/internal/repository/player"
)

// 試玩模式的 game service
type serviceDemo struct {
	gameRepo   gameRepo.Repository
	playerRepo playerRepo.Repository
}

func NewServiceDemo(gameRepo gameRepo.Repository, playerRepo playerRepo.Repository) Service {
	return &serviceDemo{gameRepo: gameRepo, playerRepo: playerRepo}
}

func (s *serviceDemo) EnterGame(input *enter_game.Input) (output *enter_game.Output) {
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
		ScoreTry:           20000 * 10000, // 試玩金額
	}
	return output
}

func (s *serviceDemo) Bet(input *bet.Input) (output *bet.Output) {
	output = &bet.Output{}
	output.Code = constants.CodeSuccess
	output.Message = "success"
	output.Data = &bet.Data{}
	return output
}
