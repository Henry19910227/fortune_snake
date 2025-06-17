package save

import "context"

type Param struct {
	Ctx        context.Context
	GameMode   string
	PlayerID   uint64
	GameResult *string
	Bet        *int
	Value      *int
	Bonus      *bool
}
