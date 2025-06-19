package save_items

import (
	"context"
	"game_server_slots_fortune_snake/internal/model/entity/player"
)

type Param struct {
	Ctx           context.Context
	GameMode      string
	Session       *player.Session
	TransactionID uint64
	RoundID       uint64
	Items         [][][]int
}
