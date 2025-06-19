package result_free

import (
	"context"
	"encoding/json"
	. "game_server_slots_fortune_snake/constants"
	freeOrderModel "game_server_slots_fortune_snake/internal/model/entity/player_free_order"
	"game_server_slots_fortune_snake/internal/model/service/result_free/save_items"
	freeOrder "game_server_slots_fortune_snake/internal/repository/player_free_order"
	resultFreeRepo "game_server_slots_fortune_snake/internal/repository/result_free"
	snowFlakeRepo "game_server_slots_fortune_snake/internal/repository/snow_flake"
	"time"
)

type serviceReal struct {
	resultFreeRepo resultFreeRepo.Repository
	freeOrderRepo  freeOrder.Repository
	snowFlakeRepo  snowFlakeRepo.Repository
}

func NewSvcInReal(resultFreeRepo resultFreeRepo.Repository,
	freeOrderRepo freeOrder.Repository,
	snowFlakeRepo snowFlakeRepo.Repository) Service {
	return &serviceReal{
		resultFreeRepo: resultFreeRepo,
		freeOrderRepo:  freeOrderRepo,
		snowFlakeRepo:  snowFlakeRepo}
}

func (s *serviceReal) SaveItems(param save_items.Param) error {
	if len(param.Items) == 0 {
		return nil
	}

	gameJson, err := json.Marshal(param.Items)
	if err != nil {
		return err
	}

	// 入庫
	table := freeOrderModel.New(param.Session)
	table.ID = s.snowFlakeRepo.GenerateID()
	table.TransactionId = param.TransactionID
	table.RoundId = param.RoundID
	table.GameFreeItems = int8(len(param.Items))
	table.GameFreeMultiplier = 1
	table.GameFreeUsedItems = 0
	table.Status = "settlement"
	table.FreeGameJson = gameJson
	table.FreeType = "play"
	table.IsSync = "no"
	table.Sync = "no"
	table.CreatedAt = time.Now().UnixMilli()
	table.UpdatedAt = time.Now().UnixMilli()

	// 轉碼
	list := make([]string, 0)
	for _, item := range param.Items {
		b, err := json.Marshal(item)
		if err != nil {
			continue
		}
		list = append(list, string(b))
	}

	// 存入緩存List
	if err := s.resultFreeRepo.SaveItems(param.Ctx, GameModeReal, param.Session.PlayerId, list); err != nil {
		return err
	}

	return nil
}

func (s *serviceReal) PopFirstItem(ctx context.Context, playerID uint64) ([][]int, error) {
	resultStr, err := s.resultFreeRepo.PopFirstItem(ctx, GameModeReal, playerID)
	if err != nil {
		return nil, err
	}
	result := make([][]int, 0)
	if err = json.Unmarshal([]byte(resultStr), &result); err != nil {
		return nil, err
	}
	return result, nil
}

func (s *serviceReal) Amount(ctx context.Context, playerID uint64) (int64, error) {
	return s.resultFreeRepo.Amount(ctx, GameModeReal, playerID)
}
