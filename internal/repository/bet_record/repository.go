package bet_record

import (
	model "game_server_slots_fortune_snake/internal/model/entity/bet_record"
	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func New(db *gorm.DB) Repository {
	return &repository{db: db}
}

func (r *repository) Create(table *model.Table) (id uint64, err error) {
	err = r.db.Model(&model.Table{}).Create(&table).Error
	if err != nil {
		return 0, err
	}
	return table.ID, err
}
