package game

import (
	"context"
	"errors"
	"fmt"
	. "game_server_slots_fortune_snake/constants"
	gameModel "game_server_slots_fortune_snake/internal/model/entity/game"
	"github.com/redis/go-redis/v9"
	"time"
)

type repository struct {
	rdb      *redis.Client
	gameMode string
}

func New(rdb *redis.Client) Repository {
	return &repository{rdb: rdb, gameMode: GameModeDemo}
}

func (r *repository) Mode(gameMode string) Repository {
	r.gameMode = gameMode
	return r
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

func (r *repository) SetSpecialMode(ctx context.Context, playerId uint64, specialMode bool) error {
	key := fmt.Sprintf(CacheNamePlayerSession, playerId)
	err := r.rdb.SetEx(ctx, key, specialMode, 20*24*time.Hour).Err()
	return err
}

func (r *repository) IsSpecialMode(ctx context.Context, playerId uint64) (bool, error) {
	key := fmt.Sprintf(CacheNamePlayerSession, playerId)
	result, err := r.rdb.Get(ctx, key).Result()
	if err != nil {
		return false, err
	}
	if result == "0" {
		return false, nil
	}
	return true, nil
}

func (r *repository) SaveFreeResults(ctx context.Context, playerId int, items []string) error {
	key := fmt.Sprintf(CacheNameFreeResults, playerId, r.gameMode)
	if len(items) == 0 {
		return nil
	}
	pipe := r.rdb.TxPipeline()
	pipe.RPush(ctx, key, items)
	pipe.Expire(ctx, key, 20*24*time.Hour)
	_, err := pipe.Exec(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (r *repository) PopFirstResult(ctx context.Context, playerId int) (string, error) {
	key := fmt.Sprintf(CacheNameFreeResults, playerId, r.gameMode)
	result, err := r.rdb.LPop(ctx, key).Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return "", nil
		}
		return "", err
	}
	return result, nil
}
