package player

import (
	model "game_server_slots_fortune_snake/internal/model/entity/player"
	"gorm.io/gorm"
	"time"
)

type repository struct {
	db *gorm.DB
}

func New(db *gorm.DB) Repository {
	return &repository{db: db}
}

func (r *repository) UpdateBalance(playerId uint64, newBalance int64) error {
	if err := r.db.Model(&model.Table{}).
		Where("id = ?", playerId).
		Updates(map[string]interface{}{"balance": newBalance, "updated_at": time.Now().UnixMilli()}).Error; err != nil {
		return err
	}
	return nil
}
