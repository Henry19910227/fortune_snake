package bucket

import (
	"fmt"
	model "game_server_slots_fortune_snake/internal/model/config/bucket"
	resultModel "game_server_slots_fortune_snake/internal/model/entity/result_free"
	"game_server_slots_fortune_snake/internal/model/repository/bucket_free/all_items"
	"sync"
)

type key struct {
	lowerLimit float64
	upperLimit float64
}

type repository struct {
	config    []*model.Item               // 容量配置檔
	bucketMap map[key][]*resultModel.Item // 桶儲存器
	quota     map[key]int                 // 配額
	mu        sync.Mutex
}

func New(config []*model.Item) Repository {
	bucketMap := make(map[key][]*resultModel.Item)
	quota := make(map[key]int)
	for _, configItem := range config {
		k := key{lowerLimit: configItem.LowerLimit, upperLimit: configItem.UpperLimit}
		bucketMap[k] = []*resultModel.Item{}
		quota[k] = configItem.MaxCapacity
	}
	return &repository{config: config, bucketMap: bucketMap, quota: quota}
}

func (r *repository) Save(item *resultModel.Item) {
	// 以賠率查找存放的桶
	k := r.findKey(item.Rate)
	if k == nil {
		return
	}
	// 上鎖
	r.mu.Lock()
	defer r.mu.Unlock()
	// 查詢該桶剩餘配額，如無配額則 return
	amount, ok := r.quota[*k]
	if amount == 0 || !ok {
		return
	}
	// 將數據存放到桶內
	if bucket, ok := r.bucketMap[*k]; ok {
		bucket = append(bucket, item)
		r.bucketMap[*k] = bucket
	}
	// 扣除該桶配額
	if amount, ok := r.quota[*k]; ok {
		r.quota[*k] = amount - 1
	}
}

func (r *repository) Quota() int {
	r.mu.Lock()
	defer r.mu.Unlock()
	var total int
	for k, v := range r.quota {
		total += v
		if v > 0 {
			fmt.Printf("%f : %d \n", k, v)
		}
	}
	return total
}

func (r *repository) Items(lowerLimit float64, upperLimit float64) []*resultModel.Item {
	r.mu.Lock()
	defer r.mu.Unlock()
	k := key{lowerLimit: lowerLimit, upperLimit: upperLimit}
	bucket, ok := r.bucketMap[k]
	if !ok {
		return []*resultModel.Item{}
	}
	return bucket
}

func (r *repository) AllItems() (output *all_items.Output) {
	items := make([]*resultModel.Item, 0)
	for _, bucket := range r.bucketMap {
		for _, item := range bucket {
			items = append(items, item)
		}
	}
	output = all_items.NewOutput(items)
	return output
}

func (r *repository) findKey(rate float64) *key {
	if rate == 0 {
		return &key{lowerLimit: 0, upperLimit: 0}
	}
	for _, cfg := range r.config {
		if rate > cfg.LowerLimit && rate <= cfg.UpperLimit {
			return &key{lowerLimit: cfg.LowerLimit, upperLimit: cfg.UpperLimit}
		}
	}
	return nil
}
