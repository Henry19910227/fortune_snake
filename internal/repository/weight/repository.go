package weight

import (
	weightModel "game_server_slots_fortune_snake/internal/model/weight"
)

type repository struct {
}

func New() Repository {
	return &repository{}
}

func (r *repository) BaseWeight() []*weightModel.Stat {
	//TODO implement me
	panic("implement me")
}

func (r *repository) BaseWeightH() []*weightModel.Stat {
	//TODO implement me
	panic("implement me")
}

func (r *repository) FreeWeight() []*weightModel.Stat {
	//TODO implement me
	panic("implement me")
}

func (r *repository) FreeWeightH() []*weightModel.Stat {
	//TODO implement me
	panic("implement me")
}
