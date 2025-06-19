package result_free

import (
	"context"
	"encoding/json"
	. "game_server_slots_fortune_snake/constants"
	"game_server_slots_fortune_snake/internal/model/service/result_free/save_items"
	resultFreeRepo "game_server_slots_fortune_snake/internal/repository/result_free"
)

type serviceDemo struct {
	resultFreeRepo resultFreeRepo.Repository
}

func NewSvcInDemo(resultFreeRepo resultFreeRepo.Repository) Service {
	return &serviceDemo{resultFreeRepo: resultFreeRepo}
}

func (s serviceDemo) SaveItems(param save_items.Param) error {
	if len(param.Items) == 0 {
		return nil
	}
	// 轉碼
	list := make([]string, 0)
	for _, item := range param.Items {
		b, err := json.Marshal(item)
		if err != nil {
			continue
		}
		list = append(list, string(b))
	}

	return s.resultFreeRepo.SaveItems(param.Ctx, GameModeDemo, param.Session.PlayerId, list)
}

func (s serviceDemo) PopFirstItem(ctx context.Context, playerID uint64) ([][]int, error) {
	resultStr, err := s.resultFreeRepo.PopFirstItem(ctx, GameModeDemo, playerID)
	if err != nil {
		return nil, err
	}
	result := make([][]int, 0)
	if err = json.Unmarshal([]byte(resultStr), &result); err != nil {
		return nil, err
	}
	return result, nil
}

func (s serviceDemo) Amount(ctx context.Context, playerID uint64) (int64, error) {
	return s.resultFreeRepo.Amount(ctx, GameModeDemo, playerID)
}
