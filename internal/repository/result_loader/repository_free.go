package result

import (
	"game_server_slots_fortune_snake/constants"
	model "game_server_slots_fortune_snake/internal/model/entity/result"
	errMsg "game_server_slots_fortune_snake/internal/model/err"
	"gorm.io/gorm"
	"math/rand"
	"sync"
)

type repositoryFree struct {
	db         *gorm.DB
	mu         sync.RWMutex
	resultsMap map[float64][]*model.Item
}

func (r *repositoryFree) CreateItems(items []*model.Item) (err error) {
	err = r.db.Table("fortune_snake_results_free").Create(items).Error
	if err != nil {
		return errMsg.New(constants.CodeInternalError, err.Error(), err)
	}
	return nil
}

func (r *repositoryFree) LoadData() (err error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	// 從 db 讀取 model
	var items []*model.Item
	if err := r.db.Table("fortune_snake_results_free").Find(&items).Error; err != nil {
		return errMsg.New(constants.CodeInternalError, err.Error(), err)
	}
	// 將數據依照 rate 分類
	for _, item := range items {
		if _, ok := r.resultsMap[item.Rate]; !ok {
			r.resultsMap[item.Rate] = []*model.Item{}
		}
		r.resultsMap[item.Rate] = append(r.resultsMap[item.Rate], item)
	}
	return nil
}

func (r *repositoryFree) Random(rate float64) (*model.Item, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	items, ok := r.resultsMap[rate]
	if !ok {
		return nil, errMsg.New(constants.CodeInternalError, "找不到賠率", nil)
	}
	index := rand.Intn(len(items))
	item := items[index]
	return item, nil
}
