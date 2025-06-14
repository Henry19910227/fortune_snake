package game

import (
	"context"
	"encoding/json"
	. "game_server_slots_fortune_snake/constants"
	betModel "game_server_slots_fortune_snake/internal/model/controller/game/bet"
	gameRepo "game_server_slots_fortune_snake/internal/repository/game"
	"math/rand"
)

type service struct {
	gameRepo gameRepo.Repository
	gameMode string
}

func NewService(gameRepo gameRepo.Repository) Service {
	return &service{gameRepo: gameRepo, gameMode: GameModeDemo}
}

func (s *service) GameMode(gameMode string) Service {
	s.gameMode = gameMode
	return s
}

func (s *service) SpinMode() int {
	probability := rand.Float64()
	if probability <= 0.5 {
		return 1
	}
	return 0
}

func (s *service) SaveGameResult(ctx context.Context, playerID uint64, item *betModel.GameResult) (err error) {
	if item == nil {
		return nil
	}
	data, err := json.Marshal(item)
	if err != nil {
		return err
	}
	if err = s.gameRepo.Mode(s.gameMode).SaveGameResult(ctx, playerID, string(data)); err != nil {
		return err
	}
	return nil
}

func (s *service) SaveBet(ctx context.Context, playerID uint64, bet int) (err error) {
	if err = s.gameRepo.Mode(s.gameMode).SaveBet(ctx, playerID, bet); err != nil {
		return err
	}
	return nil
}

func (s *service) SaveValue(ctx context.Context, playerID uint64, value int) (err error) {
	if err = s.gameRepo.Mode(s.gameMode).SaveValue(ctx, playerID, value); err != nil {
		return err
	}
	return nil
}

func (s *service) GetGameResult(ctx context.Context, playerID uint64) (item *betModel.GameResult, err error) {
	data, err := s.gameRepo.Mode(s.gameMode).GetGameResult(ctx, playerID)
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

func (s *service) GetBet(ctx context.Context, playerID uint64) (bet int, err error) {
	data, err := s.gameRepo.Mode(s.gameMode).GetBet(ctx, playerID)
	if err != nil {
		return 0, err
	}
	return data, nil
}

func (s *service) GetValue(ctx context.Context, playerID uint64) (value int, err error) {
	data, err := s.gameRepo.Mode(s.gameMode).GetValue(ctx, playerID)
	if err != nil {
		return 0, err
	}
	return data, nil
}
