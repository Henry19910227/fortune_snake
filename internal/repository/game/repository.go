package game

import (
	"context"
	"fmt"
	"game_server_slots_fortune_snake/constants"
	gameModel "game_server_slots_fortune_snake/internal/model/entity/game"
	"github.com/redis/go-redis/v9"
)

type repository struct {
	rdb *redis.Client
}

func New(rdb *redis.Client) Repository {
	return &repository{rdb: rdb}
}

func (r *repository) Info() (info *gameModel.Info, err error) {
	info = &gameModel.Info{
		Bets:               []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		Values:             []float64{100, 1000, 4000, 20000},
		Multipler:          10,
		MultipleScoreLimit: 2000,
	}
	return info, nil
}

func (r *repository) IsSpecialMode(ctx context.Context, playerId uint64) (bool, error) {
	key := fmt.Sprintf(constants.CacheNamePlayerSession, playerId)
	result, err := r.rdb.Get(ctx, key).Result()
	if err != nil {
		return false, err
	}
	if result == "0" {
		return false, nil
	}
	return true, nil
}
