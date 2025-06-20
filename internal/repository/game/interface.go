package game

import (
	"context"
	gameModel "game_server_slots_fortune_snake/internal/model/entity/game"
	"game_server_slots_fortune_snake/internal/model/repository/game/save"
)

type Repository interface {
	// Info 獲取遊戲配置數據
	Info() (info *gameModel.Info, err error)

	Save(param save.Param) (err error)

	SaveGameResult(ctx context.Context, gameMode string, playerID uint64, data string) (err error)

	SaveBet(ctx context.Context, gameMode string, playerID uint64, bet int) (err error)

	SaveValue(ctx context.Context, gameMode string, playerID uint64, value int) (err error)

	SaveBonus(ctx context.Context, gameMode string, playerID uint64, bonus bool) (err error)

	SaveFatherID(ctx context.Context, gameMode string, playerID uint64, fatherID uint64) (err error)

	SaveRoundID(ctx context.Context, gameMode string, playerID uint64, roundID uint64) (err error)

	GetParam(ctx context.Context, gameMode string, playerID uint64) (results []interface{}, err error)

	GetGameResult(ctx context.Context, gameMode string, playerID uint64) (data string, err error)

	GetBet(ctx context.Context, gameMode string, playerID uint64) (bet int, err error)

	GetValue(ctx context.Context, gameMode string, playerID uint64) (value int, err error)

	GetFatherID(ctx context.Context, gameMode string, playerID uint64) (fatherID uint64, err error)

	GetRoundID(ctx context.Context, gameMode string, playerID uint64) (roundID uint64, err error)
}
