package game

import (
	"context"
	"errors"
	"fmt"
	. "game_server_slots_fortune_snake/constants"
	gameModel "game_server_slots_fortune_snake/internal/model/entity/game"
	"game_server_slots_fortune_snake/internal/model/repository/game/save"
	"github.com/redis/go-redis/v9"
	"strconv"
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

func (r *repository) Save(param save.Param) (err error) {
	key := fmt.Sprintf(CacheNamePlayerGameInfo, param.PlayerID, GameCode, param.GameMode)
	pipe := r.rdb.TxPipeline()

	fields := make(map[string]interface{})
	if param.GameResult != nil {
		fields["GameResult"] = *param.GameResult
	}
	if param.Bet != nil {
		fields["Bet"] = *param.Bet
	}
	if param.Value != nil {
		fields["Value"] = *param.Value
	}
	if param.Bonus != nil {
		fields["Bonus"] = *param.Bonus
	}
	if len(fields) > 0 {
		pipe.HSet(param.Ctx, key, fields)
	}
	pipe.Expire(param.Ctx, key, CacheExpiredPlayerGameInfo)
	_, err = pipe.Exec(param.Ctx)
	return err
}

func (r *repository) SaveGameResult(ctx context.Context, gameMode string, playerID uint64, data string) (err error) {
	key := fmt.Sprintf(CacheNamePlayerGameInfo, playerID, GameCode, gameMode)
	pipe := r.rdb.TxPipeline()
	pipe.HSet(ctx, key, "GameResults", data)
	pipe.Expire(ctx, key, CacheExpiredPlayerGameInfo)
	_, err = pipe.Exec(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (r *repository) SaveBet(ctx context.Context, gameMode string, playerID uint64, bet int) (err error) {
	key := fmt.Sprintf(CacheNamePlayerGameInfo, playerID, GameCode, gameMode)
	pipe := r.rdb.TxPipeline()
	pipe.HSet(ctx, key, "Bet", bet)
	pipe.Expire(ctx, key, CacheExpiredPlayerGameInfo)
	_, err = pipe.Exec(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (r *repository) SaveValue(ctx context.Context, gameMode string, playerID uint64, value int) (err error) {
	key := fmt.Sprintf(CacheNamePlayerGameInfo, playerID, GameCode, gameMode)
	pipe := r.rdb.TxPipeline()
	pipe.HSet(ctx, key, "Value", value)
	pipe.Expire(ctx, key, CacheExpiredPlayerGameInfo)
	_, err = pipe.Exec(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (r *repository) SaveBonus(ctx context.Context, gameMode string, playerID uint64, bonus bool) (err error) {
	key := fmt.Sprintf(CacheNamePlayerGameInfo, playerID, GameCode, gameMode)
	pipe := r.rdb.TxPipeline()
	pipe.HSet(ctx, key, "Bonus", bonus)
	pipe.Expire(ctx, key, CacheExpiredPlayerGameInfo)
	_, err = pipe.Exec(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (r *repository) SaveFatherID(ctx context.Context, gameMode string, playerID uint64, fatherID uint64) (err error) {
	key := fmt.Sprintf(CacheNamePlayerGameInfo, playerID, GameCode, gameMode)
	pipe := r.rdb.TxPipeline()
	pipe.HSet(ctx, key, "FatherID", fatherID)
	pipe.Expire(ctx, key, CacheExpiredPlayerGameInfo)
	_, err = pipe.Exec(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (r *repository) SaveRoundID(ctx context.Context, gameMode string, playerID uint64, roundID uint64) (err error) {
	key := fmt.Sprintf(CacheNamePlayerGameInfo, playerID, GameCode, gameMode)
	pipe := r.rdb.TxPipeline()
	pipe.HSet(ctx, key, "RoundID", roundID)
	pipe.Expire(ctx, key, CacheExpiredPlayerGameInfo)
	_, err = pipe.Exec(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (r *repository) GetParam(ctx context.Context, gameMode string, playerID uint64) (results []interface{}, err error) {
	key := fmt.Sprintf(CacheNamePlayerGameInfo, playerID, GameCode, gameMode)
	fields := []string{"Bet", "Value", "Bonus"}
	data, err := r.rdb.HMGet(ctx, key, fields...).Result()
	if errors.Is(err, redis.Nil) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (r *repository) GetGameResult(ctx context.Context, gameMode string, playerID uint64) (data string, err error) {
	key := fmt.Sprintf(CacheNamePlayerGameInfo, playerID, GameCode, gameMode)
	data, err = r.rdb.HGet(ctx, key, "GameResults").Result()
	if errors.Is(err, redis.Nil) {
		return "", nil
	}
	if err != nil {
		return "", err
	}
	return data, nil
}

func (r *repository) GetBet(ctx context.Context, gameMode string, playerID uint64) (bet int, err error) {
	key := fmt.Sprintf(CacheNamePlayerGameInfo, playerID, GameCode, gameMode)
	val, err := r.rdb.HGet(ctx, key, "Bet").Result()
	if errors.Is(err, redis.Nil) {
		return 0, nil
	}
	if err != nil {
		return 0, err
	}
	bet, err = strconv.Atoi(val)
	if err != nil {
		return 0, err
	}
	return bet, nil
}

func (r *repository) GetValue(ctx context.Context, gameMode string, playerID uint64) (value int, err error) {
	key := fmt.Sprintf(CacheNamePlayerGameInfo, playerID, GameCode, gameMode)
	val, err := r.rdb.HGet(ctx, key, "Value").Result()
	if errors.Is(err, redis.Nil) {
		return 0, nil
	}
	if err != nil {
		return 0, err
	}
	value, err = strconv.Atoi(val)
	if err != nil {
		return 0, err
	}
	return value, nil
}

func (r *repository) GetFatherID(ctx context.Context, gameMode string, playerID uint64) (fatherID uint64, err error) {
	key := fmt.Sprintf(CacheNamePlayerGameInfo, playerID, GameCode, gameMode)
	val, err := r.rdb.HGet(ctx, key, "FatherID").Result()
	if errors.Is(err, redis.Nil) {
		return 0, nil
	}
	if err != nil {
		return 0, err
	}
	value, err := strconv.ParseUint(val, 10, 64)
	if err != nil {
		return 0, err
	}
	return value, nil
}

func (r *repository) GetRoundID(ctx context.Context, gameMode string, playerID uint64) (roundID uint64, err error) {
	key := fmt.Sprintf(CacheNamePlayerGameInfo, playerID, GameCode, gameMode)
	val, err := r.rdb.HGet(ctx, key, "RoundID").Result()
	if errors.Is(err, redis.Nil) {
		return 0, nil
	}
	if err != nil {
		return 0, err
	}
	value, err := strconv.ParseUint(val, 10, 64)
	if err != nil {
		return 0, err
	}
	return value, nil
}
