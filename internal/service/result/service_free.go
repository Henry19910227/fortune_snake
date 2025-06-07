package result

import (
	"encoding/json"
	"game_server_slots_fortune_snake/constants"
	"game_server_slots_fortune_snake/internal/model/entity/result"
	errMsg "game_server_slots_fortune_snake/internal/model/err"
	bucketRepo "game_server_slots_fortune_snake/internal/repository/bucket"
	resultRepo "game_server_slots_fortune_snake/internal/repository/result"
)

type serviceFree struct {
	resultRepo resultRepo.Repository
	bucketRepo bucketRepo.Repository
}

func NewFree(resultRepo resultRepo.Repository, bucketRepo bucketRepo.Repository) Service {
	return &serviceFree{resultRepo: resultRepo, bucketRepo: bucketRepo}
}

func (s *serviceFree) SaveToBucket(result *result.Item) (quota int) {
	s.bucketRepo.Save(result)
	return s.bucketRepo.Quota()
}

func (s *serviceFree) Migrate() (err error) {
	// 判斷暫存區數據是否裝滿
	if s.bucketRepo.Quota() > 0 {
		return errMsg.New(constants.CodeInternalError, "盤面結果數據尚未齊全", nil)
	}
	// 獲取暫存區所有產生的盤面數據
	items := s.bucketRepo.AllItems()
	// 將盤面數據存入db
	err = s.resultRepo.CreateItems(items)
	if err != nil {
		return err
	}
	return nil
}

func (s *serviceFree) LoadData() (err error) {
	return s.resultRepo.LoadData()
}

func (s *serviceFree) Random(rate float64) ([][][]int, error) {
	item, err := s.resultRepo.Random(rate)
	if err != nil {
		return nil, err
	}
	var allSymbols [][][]int
	err = json.Unmarshal([]byte(item.AllSymbols), &allSymbols)
	if err != nil {
		return nil, errMsg.New(constants.CodeInternalError, err.Error(), err)
	}
	return allSymbols, nil
}
