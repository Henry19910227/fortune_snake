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

// 試玩模式的 game service
type serviceDemo struct {
	gameRepo   gameRepo.Repository
	playerRepo playerRepo.Repository
}

func NewServiceDemo(gameRepo gameRepo.Repository, playerRepo playerRepo.Repository) Service {
	return &serviceDemo{gameRepo: gameRepo, playerRepo: playerRepo}
}

func (s *serviceDemo) EnterGame(input *enter_game.Input) (output *enter_game.Output, err error) {
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
		ScoreTry:           20000 * 10000, // 試玩金額
	}
	return output, nil
}

func (s *serviceDemo) Bet(input *bet.Input) (output *bet.Output, err error) {
	output = &bet.Output{}
	output.Data = &bet.Data{}
	return output, nil
}
