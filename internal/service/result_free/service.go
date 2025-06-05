package result

//type service struct {
//	resultRepo resultRepo.Repository
//	bucketRepo bucketRepo.Repository
//}
//
//func New(resultRepo resultRepo.Repository, bucketRepo bucketRepo.Repository) Service {
//	return &service{resultRepo: resultRepo, bucketRepo: bucketRepo}
//}
//
//func (s *service) SaveToBucket(input *save_to_bucket.Input) (output *save_to_bucket.Output) {
//	result := input.Param.Result
//	s.bucketRepo.Save(result)
//	output = save_to_bucket.NewOutput(s.bucketRepo.Quota())
//	return output
//}
//
//func (s *service) Migrate() (err error) {
//	// 判斷暫存區數據是否裝滿
//	if s.bucketRepo.Quota() > 0 {
//		return errMsg.New(constants.CodeInternalError, "盤面結果數據尚未齊全", nil)
//	}
//	// 獲取暫存區所有產生的盤面數據
//	items := s.bucketRepo.AllItems()
//	// 將盤面數據存入db
//	err = s.resultRepo.CreateItems(items)
//	if err != nil {
//		return err
//	}
//	return nil
//}
//
//func (s *service) LoadData() (err error) {
//	return s.resultRepo.LoadData()
//}
//
//func (s *service) Random(rate float64) (output *model.Item, err error) {
//	outputData, err := s.resultRepo.Random(rate)
//	if err != nil {
//		return nil, err
//	}
//	return outputData, nil
//}
