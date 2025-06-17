package game

import (
	"context"
	"encoding/json"
	betModel "game_server_slots_fortune_snake/internal/model/controller/game/bet"
	saveRepoParam "game_server_slots_fortune_snake/internal/model/repository/game/save"
	saveSvcParam "game_server_slots_fortune_snake/internal/model/service/game/save"
	gameRepo "game_server_slots_fortune_snake/internal/repository/game"
	"game_server_slots_fortune_snake/util"
	"math/rand"
)

type service struct {
	gameRepo gameRepo.Repository
}

func NewService(gameRepo gameRepo.Repository) Service {
	return &service{gameRepo: gameRepo}
}

func (s *service) SpinMode() int {
	probability := rand.Float64()
	if probability <= 0.5 {
		return 1
	}
	return 0
}

func (s *service) Save(param saveSvcParam.Param) (err error) {
	p := saveRepoParam.Param{}
	p.Ctx = param.Ctx
	p.GameMode = param.GameMode
	p.PlayerID = param.PlayerID
	p.Bet = param.Bet
	p.Value = param.Value
	p.Bonus = param.Bonus
	if param.GameResult != nil {
		data, err := json.Marshal(param.GameResult)
		if err != nil {
			return err
		}
		p.GameResult = util.PointerString(string(data))
	}
	return s.gameRepo.Save(p)
}

func (s *service) SaveGameResult(ctx context.Context, gameMode string, playerID uint64, item *betModel.GameResult) (err error) {
	if item == nil {
		return nil
	}
	data, err := json.Marshal(item)
	if err != nil {
		return err
	}
	if err = s.gameRepo.SaveGameResult(ctx, gameMode, playerID, string(data)); err != nil {
		return err
	}
	return nil
}

func (s *service) SaveBet(ctx context.Context, gameMode string, playerID uint64, bet int) (err error) {
	if err = s.gameRepo.SaveBet(ctx, gameMode, playerID, bet); err != nil {
		return err
	}
	return nil
}

func (s *service) SaveValue(ctx context.Context, gameMode string, playerID uint64, value int) (err error) {
	if err = s.gameRepo.SaveValue(ctx, gameMode, playerID, value); err != nil {
		return err
	}
	return nil
}

func (s *service) SaveBonus(ctx context.Context, gameMode string, playerID uint64, bonus bool) (err error) {
	if err = s.gameRepo.SaveBonus(ctx, gameMode, playerID, bonus); err != nil {
		return err
	}
	return nil
}

func (s *service) SaveFatherID(ctx context.Context, gameMode string, playerID uint64, fatherID uint64) (err error) {
	if err = s.gameRepo.SaveFatherID(ctx, gameMode, playerID, fatherID); err != nil {
		return err
	}
	return nil
}

func (s *service) GetGameResult(ctx context.Context, gameMode string, playerID uint64) (item *betModel.GameResult, err error) {
	data, err := s.gameRepo.GetGameResult(ctx, gameMode, playerID)
	if err != nil {
		return nil, err
	}
	if len(data) == 0 {
		return nil, nil
	}
	gameResult := &betModel.GameResult{}
	if err = json.Unmarshal([]byte(data), gameResult); err != nil {
		return nil, err
	}
	return gameResult, nil
}

func (s *service) GetBet(ctx context.Context, gameMode string, playerID uint64) (bet int, err error) {
	data, err := s.gameRepo.GetBet(ctx, gameMode, playerID)
	if err != nil {
		return 0, err
	}
	return data, nil
}

func (s *service) GetValue(ctx context.Context, gameMode string, playerID uint64) (value int, err error) {
	data, err := s.gameRepo.GetValue(ctx, gameMode, playerID)
	if err != nil {
		return 0, err
	}
	return data, nil
}

func (s *service) GetFatherID(ctx context.Context, gameMode string, playerID uint64) (fatherID uint64, err error) {
	data, err := s.gameRepo.GetFatherID(ctx, gameMode, playerID)
	if err != nil {
		return 0, err
	}
	return data, nil
}
