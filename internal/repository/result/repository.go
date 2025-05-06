package result

import (
	"game_server_slots_fortune_snake/internal/model/entity/result"
	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func New(db *gorm.DB) Repository {
	return &repository{db: db}
}

func (r repository) Create(item *result.Item) (id int64, err error) {
	return 0, nil
}
