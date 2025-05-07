package result

import (
	"game_server_slots_fortune_snake/internal/model/service/result/save"
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

func (s *service) Save(input *save.Input) error {
	return nil
}

func (s *service) SaveToBucket(input *save_to_bucket.Input) (output *save_to_bucket.Output) {
	result := input.Param.Result
	s.bucketRepo.Save(result)
	output = save_to_bucket.NewOutput(s.bucketRepo.Quota())
	return output
}
