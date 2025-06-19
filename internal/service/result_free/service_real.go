package result_free

import (
	"context"
	"encoding/json"
	. "game_server_slots_fortune_snake/constants"
	"game_server_slots_fortune_snake/internal/model/entity/player"
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

	_, err = s.freeOrderRepo.Create(table)
	if err != nil {
		return err
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

	// 存入緩存List
	if err := s.resultFreeRepo.SaveItems(param.Ctx, GameModeReal, param.Session.PlayerId, list); err != nil {
		return err
	}

	return nil
}

func (s *serviceReal) PopFirstItem(ctx context.Context, session *player.Session) ([][]int, error) {
	// 取出一筆盤面
	resultStr, err := s.resultFreeRepo.PopFirstItem(ctx, GameModeReal, session.PlayerId)
	if err != nil {
		return nil, err
	}
	results := make([][]int, 0)
	if err = json.Unmarshal([]byte(resultStr), &results); err != nil {
		return nil, err
	}

	// 獲取剩餘盤面
	amount, err := s.resultFreeRepo.Amount(ctx, GameModeReal, session.PlayerId)
	if err != nil {
		return nil, err
	}

	// 查詢免費遊戲紀錄
	item, err := s.freeOrderRepo.Find(session)
	if err != nil {
		return nil, err
	}

	// 修改已使用免費次數
	item.GameFreeUsedItems = item.GameFreeItems - int8(amount)

	// 更新免費遊戲紀錄
	if err = s.freeOrderRepo.Update(item); err != nil {
		return nil, err
	}

	return results, nil
}

func (s *serviceReal) Amount(ctx context.Context, session *player.Session) (int64, error) {
	// 查看緩存是否有免費盤面
	amount, err := s.resultFreeRepo.Amount(ctx, GameModeReal, session.PlayerId)
	if err != nil {
		return 0, err
	}
	if amount > 0 {
		return amount, nil
	}

	// 找不到免費盤面則查看數據庫是否有盤面
	item, err := s.freeOrderRepo.Find(session)
	if err != nil {
		return 0, err
	}

	// 計算剩餘免費次數
	if item.GameFreeItems-item.GameFreeUsedItems <= 0 {
		return 0, nil
	}

	// 將字串轉為array
	var results [][][]int
	if err = json.Unmarshal(item.FreeGameJson, &results); err != nil {
		return 0, err
	}
	list := make([]string, 0)
	for _, result := range results {
		b, err := json.Marshal(result)
		if err != nil {
			continue
		}
		list = append(list, string(b))
	}

	// 存入緩存List
	if err = s.resultFreeRepo.SaveItems(ctx, GameModeReal, session.PlayerId, list[item.GameFreeUsedItems:]); err != nil {
		return 0, err
	}

	return int64(len(results)), nil
}
