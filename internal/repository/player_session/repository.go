package player_session

import (
	"context"
	"encoding/json"
	"fmt"
	. "game_server_slots_fortune_snake/constants"
	playerModel "game_server_slots_fortune_snake/internal/model/entity/player"
	"github.com/redis/go-redis/v9"
)

type repository struct {
	rdb *redis.Client
}

func New(rdb *redis.Client) Repository {
	return &repository{rdb: rdb}
}

func (r *repository) FindPlayerSessionById(ctx context.Context, playerId uint64) (*playerModel.Session, error) {
	key := fmt.Sprintf(CacheNamePlayerSession, playerId)
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

func (r *repository) UpdateSessionById(ctx context.Context, playerId uint64, session *playerModel.Session) error {
	key := fmt.Sprintf(CacheNamePlayerSession, playerId)

	data, err := json.Marshal(session)
	if err != nil {
		return fmt.Errorf("failed to marshal session: %w", err)
	}

	if err := r.rdb.Set(ctx, key, data, CacheExpiredPlayerSession).Err(); err != nil {
		return err
	}

	return nil
}

func (r *repository) Balance() (balance int, err error) {
	return 0, err
}
