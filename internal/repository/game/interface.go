package game

import (
	"context"
	gameModel "game_server_slots_fortune_snake/internal/model/entity/game"
)

type Repository interface {
	Mode(gameMode string) Repository
	// Info 獲取遊戲配置數據
	Info() (info *gameModel.Info, err error)

	SaveGameResult(ctx context.Context, playerID uint64, data string) (err error)

	SaveBet(ctx context.Context, playerID uint64, bet int) (err error)

	SaveValue(ctx context.Context, playerID uint64, value int) (err error)

	SaveFatherID(ctx context.Context, playerID uint64, fatherID uint64) (err error)

	GetGameResult(ctx context.Context, playerID uint64) (data string, err error)

	GetBet(ctx context.Context, playerID uint64) (bet int, err error)

	GetValue(ctx context.Context, playerID uint64) (value int, err error)

	GetFatherID(ctx context.Context, playerID uint64) (fatherID uint64, err error)
}
