package result

import (
	"game_server_slots_fortune_snake/constants"
	errMsg "game_server_slots_fortune_snake/internal/model/err"
	"game_server_slots_fortune_snake/internal/model/repository/result/create_items"
	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func New(db *gorm.DB) Repository {
	return &repository{db: db}
}

func (r *repository) CreateItems(input *create_items.Input) (err error) {
	err = r.db.Create(input.GetItems()).Error
	if err != nil {
		return errMsg.New(constants.CodeInternalError, err.Error(), err)
	}
	return nil
}
