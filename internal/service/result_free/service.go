package result_free

import (
	"context"
	"encoding/json"
	"game_server_slots_fortune_snake/constants"
	resultFreeRepo "game_server_slots_fortune_snake/internal/repository/result_free"
)

type service struct {
	resultFreeRepo resultFreeRepo.Repository
	gameMode       string
}

func New(resultFreeRepo resultFreeRepo.Repository) Service {
	return &service{resultFreeRepo: resultFreeRepo, gameMode: constants.GameModeDemo}
}

func (s *service) GameMode(gameMode string) Service {
	s.gameMode = gameMode
	return s
}

func (s *service) SaveItems(ctx context.Context, playerId int, items [][][]int) error {
	if len(items) == 0 {
		return nil
	}
	list := make([]string, 0)
	for _, item := range items {
		b, err := json.Marshal(item)
		if err != nil {
			continue
		}
		list = append(list, string(b))
	}
	return s.resultFreeRepo.GameMode(s.gameMode).SaveItems(ctx, playerId, list)
}

func (s *service) PopFirstItem(ctx context.Context, playerId int) ([][]int, error) {
	resultStr, err := s.resultFreeRepo.GameMode(s.gameMode).PopFirstItem(ctx, playerId)
	if err != nil {
		return nil, err
	}
	result := make([][]int, 0)
	if err = json.Unmarshal([]byte(resultStr), &result); err != nil {
		return nil, err
	}
	return result, nil
}

func (s *service) Amount(ctx context.Context, playerId int) (int64, error) {
	return s.resultFreeRepo.GameMode(s.gameMode).Amount(ctx, playerId)
}
