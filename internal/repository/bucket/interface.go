package result

import (
	resultModel "game_server_slots_fortune_snake/internal/model/entity/result"
	"game_server_slots_fortune_snake/internal/model/repository/bucket/all_items"
)

type Repository interface {
	Save(item *resultModel.Item)
	// Quota 獲取剩餘總配額
	Quota() int
	// Items 以一個範圍區間獲取 items
	Items(lowerLimit float64, upperLimit float64) []*resultModel.Item
	// AllItems 獲取所有 item
	AllItems() (output *all_items.Output)
}
