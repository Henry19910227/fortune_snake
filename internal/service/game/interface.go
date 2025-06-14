package game

import (
	"context"
	betModel "game_server_slots_fortune_snake/internal/model/controller/game/bet"
)

type Service interface {
	GameMode(gameMode string) Service

	SpinMode() int

	SaveGameResult(ctx context.Context, playerID uint64, item *betModel.GameResult) (err error)

	SaveBet(ctx context.Context, playerID uint64, bet int) (err error)

	SaveValue(ctx context.Context, playerID uint64, value int) (err error)

	GetGameResult(ctx context.Context, playerID uint64) (item *betModel.GameResult, err error)

	GetBet(ctx context.Context, playerID uint64) (bet int, err error)

	GetValue(ctx context.Context, playerID uint64) (value int, err error)
}
