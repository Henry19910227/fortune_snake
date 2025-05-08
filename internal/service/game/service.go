package game

import (
	"game_server_slots_fortune_snake/constants"
	errMsg "game_server_slots_fortune_snake/internal/model/err"
	"game_server_slots_fortune_snake/internal/model/repository/game/info"
	"game_server_slots_fortune_snake/internal/model/service/game/bet"
	"game_server_slots_fortune_snake/internal/model/service/game/enter_game"
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

func (s *service) EnterGame(input *enter_game.Input) (output *enter_game.Output, err error) {
	// 獲取遊戲配置
	infoOutput, err := s.gameRepo.Info(info.NewInput())
	if err != nil {
		return nil, errMsg.New(constants.CodeBadRequest, err.Error(), err)
	}
	// 獲取玩家遊戲緩存數據
	playerGameData, err := s.playerRepo.GameData()
	if err != nil {
		return nil, errMsg.New(constants.CodeBadRequest, err.Error(), err)
	}
	// 處理回傳
	output = &enter_game.Output{}
	output.Data = &enter_game.Data{
		Bets:               infoOutput.GetInfo().Bets,
		Values:             infoOutput.GetInfo().Values,
		Bet:                playerGameData.Bet,
		Value:              playerGameData.Value,
		GameMode:           input.Session.Mode,
		Multipler:          infoOutput.GetInfo().Multipler,
		MultipleScoreLimit: infoOutput.GetInfo().MultipleScoreLimit,
		ScoreTry:           0, // 真實遊玩回傳 0
	}
	return output, nil
}

func (s *service) Bet(input *bet.Input) (output *bet.Output, err error) {
	output = &bet.Output{}
	output.Data = &bet.Data{}
	return output, nil
}
