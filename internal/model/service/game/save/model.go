package save

import (
	"context"
	betModel "game_server_slots_fortune_snake/internal/model/controller/game/bet"
)

type Param struct {
	Ctx        context.Context
	GameMode   string
	PlayerID   uint64
	GameResult *betModel.GameResult
	Bet        *int
	Value      *int
	Bonus      *bool
}
