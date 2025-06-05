package result

import (
	model "game_server_slots_fortune_snake/internal/model/entity/result"
	"gorm.io/gorm"
)

type repositoryFree struct {
	db         *gorm.DB
	resultsMap map[float64][]*model.Item
}

func NewFree(db *gorm.DB) Repository {
	return &repositoryFree{db: db, resultsMap: make(map[float64][]*model.Item)}
}

func (r *repositoryFree) CreateItems(items []*model.Item) (err error) {
	//TODO implement me
	panic("implement me")
}

func (r *repositoryFree) LoadData() (err error) {
	//TODO implement me
	panic("implement me")
}

func (r *repositoryFree) Random(rate float64) (*model.Item, error) {
	//TODO implement me
	panic("implement me")
}
