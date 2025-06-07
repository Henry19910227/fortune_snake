package result_free

import (
	"context"
	"errors"
	"fmt"
	"game_server_slots_fortune_snake/constants"
	"github.com/redis/go-redis/v9"
	"time"
)

type repository struct {
	rdb      *redis.Client
	gameMode string
}

func New(rdb *redis.Client) Repository {
	return &repository{rdb: rdb, gameMode: constants.GameModeDemo}
}

func (r *repository) GameMode(gameMode string) Repository {
	r.gameMode = gameMode
	return r
}

func (r *repository) SaveItems(ctx context.Context, playerId int, items []string) error {
	key := fmt.Sprintf(constants.CacheNameFreeResults, playerId, r.gameMode)
	if len(items) == 0 {
		return nil
	}
	pipe := r.rdb.TxPipeline()
	pipe.Del(ctx, key)
	pipe.RPush(ctx, key, items)
	pipe.Expire(ctx, key, 20*24*time.Hour)
	_, err := pipe.Exec(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (r *repository) PopFirstItem(ctx context.Context, playerId int) (string, error) {
	key := fmt.Sprintf(constants.CacheNameFreeResults, playerId, r.gameMode)
	result, err := r.rdb.LPop(ctx, key).Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return "", nil
		}
		return "", err
	}
	return result, nil
}
