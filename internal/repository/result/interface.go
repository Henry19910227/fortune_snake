package result

import (
	resultModel "game_server_slots_fortune_snake/internal/model/result"
)

type Repository interface {
	Save(item *resultModel.TableItem)
	// Quota 獲取總配額
	Quota() int
}
