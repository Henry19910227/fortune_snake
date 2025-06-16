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

func (s *service) SaveItems(ctx context.Context, gameMode string, playerID uint64, items [][][]int) error {
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
	return s.resultFreeRepo.SaveItems(ctx, gameMode, playerID, list)
}

func (s *service) PopFirstItem(ctx context.Context, gameMode string, playerID uint64) ([][]int, error) {
	resultStr, err := s.resultFreeRepo.PopFirstItem(ctx, gameMode, playerID)
	if err != nil {
		return nil, err
	}
	result := make([][]int, 0)
	if err = json.Unmarshal([]byte(resultStr), &result); err != nil {
		return nil, err
	}
	return result, nil
}

func (s *service) Amount(ctx context.Context, gameMode string, playerID uint64) (int64, error) {
	return s.resultFreeRepo.Amount(ctx, gameMode, playerID)
}
