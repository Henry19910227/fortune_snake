package result

import (
	"encoding/json"
	"game_server_slots_fortune_snake/constants"
	errMsg "game_server_slots_fortune_snake/internal/model/err"
	"game_server_slots_fortune_snake/internal/model/service/result/migrate"
	"game_server_slots_fortune_snake/internal/model/service/result/save_to_bucket"
	bucketRepo "game_server_slots_fortune_snake/internal/repository/bucket"
	resultRepo "game_server_slots_fortune_snake/internal/repository/result"
)

type service struct {
	resultRepo resultRepo.Repository
	bucketRepo bucketRepo.Repository
}

func New(resultRepo resultRepo.Repository, bucketRepo bucketRepo.Repository) Service {
	return &service{resultRepo: resultRepo, bucketRepo: bucketRepo}
}

func (s *service) SaveToBucket(input *save_to_bucket.Input) (output *save_to_bucket.Output) {
	result := input.Param.Result
	s.bucketRepo.Save(result)
	output = save_to_bucket.NewOutput(s.bucketRepo.Quota())
	return output
}

func (s *service) Migrate(input *migrate.Input) (err error) {
	// 判斷暫存區數據是否裝滿
	if s.bucketRepo.Quota() > 0 {
		return errMsg.New(constants.CodeInternalError, "盤面結果數據尚未齊全", nil)
	}
	// 獲取暫存區所有產生的盤面數據
	output := s.bucketRepo.AllItems()
	// 將盤面數據存入db
	err = s.resultRepo.CreateItems(output.Items())
	if err != nil {
		return err
	}
	return nil
}

func (s *service) LoadData() (err error) {
	return s.resultRepo.LoadData()
}

func (s *service) Random(rate float64) ([][]int, error) {
	item, err := s.resultRepo.Random(rate)
	if err != nil {
		return nil, err
	}
	var symbols [][]int
	err = json.Unmarshal([]byte(item.Symbols), &symbols)
	if err != nil {
		return nil, errMsg.New(constants.CodeInternalError, err.Error(), err)
	}
	return symbols, nil
}
