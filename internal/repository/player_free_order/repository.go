package player_free_order

import (
	"game_server_slots_fortune_snake/internal/model/entity/player"
	model "game_server_slots_fortune_snake/internal/model/entity/player_free_order"
	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func New(db *gorm.DB) Repository {
	return &repository{db: db}
}

func (r *repository) Create(item *model.Table) (id uint64, err error) {
	err = r.db.Model(&model.Table{}).Create(item).Error
	if err != nil {
		return 0, err
	}
	return item.ID, err
}

func (r *repository) Find(session *player.Session) (item *model.Table, err error) {
	item = &model.Table{}
	err = r.db.
		Model(model.Table{}).
		Where("game_id = ? AND player_id = ? AND status = ?", session.GameId, session.PlayerId, "settlement").
		Order("created_at DESC").
		Limit(1).
		Take(item).Error
	if err != nil {
		return nil, err
	}
	return item, nil
}

func (r *repository) Update(item *model.Table) (err error) {
	err = r.db.Model(&model.Table{}).Where("id = ?", item.ID).Save(item).Error
	return err
}
