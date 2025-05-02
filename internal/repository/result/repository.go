package result

import (
	"game_server_slots_fortune_snake/internal/model"
	resultModel "game_server_slots_fortune_snake/internal/model/result"
	"sync"
)

type repository struct {
	config []*model.ResultConfig
	bucket [][]*resultModel.TableItem // 桶儲存器
	quota  []int                      // 配額
	mu     sync.Mutex
}

func New(config []*model.ResultConfig) Repository {
	bucket := make([][]*resultModel.TableItem, 0, len(config))
	quota := make([]int, 0, len(config))
	for _, configItem := range config {
		bucket = append(bucket, []*resultModel.TableItem{})
		quota = append(quota, configItem.MaxCapacity)
	}
	return &repository{config: config, bucket: bucket, quota: quota}
}

func (r *repository) Save(item *resultModel.TableItem) {
	r.mu.Lock()
	defer r.mu.Unlock()
	// 以賠率查找存放的桶
	index := r.findIndex(item.Rate)
	if index == -1 {
		return
	}
	// 查詢該桶剩餘配額
	if r.quota[index] == 0 {
		return
	}
	// 將數據存放到桶內
	r.bucket[index] = append(r.bucket[index], item)
	// 扣除該桶配額
	r.quota[index]--
}

func (r *repository) Quota() int {
	r.mu.Lock()
	defer r.mu.Unlock()
	var total int
	for _, v := range r.quota {
		total += v
	}
	return total
}

func (r *repository) findIndex(rate float64) int {
	if rate == 0 {
		return 0
	}
	for i, cfg := range r.config {
		if rate > cfg.LowerLimit && rate <= cfg.UpperLimit {
			return i
		}
	}
	return -1
}
