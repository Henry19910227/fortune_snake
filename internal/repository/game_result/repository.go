package game_result

import (
	model "game_server_slots_fortune_snake/internal/model/entity/game_result"
	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func New(db *gorm.DB) Repository {
	return &repository{db: db}
}

func (r repository) Create(item *model.Table) (id uint64, err error) {
	err = r.db.Model(&model.Table{}).Create(item).Error
	if err != nil {
		return 0, err
	}
	return item.ID, err
}
