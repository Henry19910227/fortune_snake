package player

import (
	"context"
	"encoding/json"
	"fmt"
	"game_server_slots_fortune_snake/constants"
	playerModel "game_server_slots_fortune_snake/internal/model/player"
	"github.com/redis/go-redis/v9"
)

type repository struct {
	rdb *redis.Client
}

func New(rdb *redis.Client) Repository {
	return &repository{rdb: rdb}
}

func (r *repository) FindPlayerSessionById(ctx context.Context, playerId uint64) (*playerModel.Session, error) {
	key := fmt.Sprintf(constants.CacheNamePlayerSession, playerId)
	result, err := r.rdb.Get(ctx, key).Result()
	if err != nil {
		return nil, err
	}
	var session playerModel.Session
	if err := json.Unmarshal([]byte(result), &session); err != nil {
		return nil, err
	}
	return &session, nil
}

func (r *repository) Balance() (balance int, err error) {
	return 0, err
}

func (r *repository) GameData() (*playerModel.GameData, error) {
	return &playerModel.GameData{Bet: 1, Value: 100}, nil
}
