package player

import (
	"context"
	"errors"
	"game_server_slots_fortune_snake/constants"
	playerModel "game_server_slots_fortune_snake/internal/model/entity/player"
	errMsg "game_server_slots_fortune_snake/internal/model/err"
	playerRepo "game_server_slots_fortune_snake/internal/repository/player"
	"game_server_slots_fortune_snake/internal/server"
	"game_server_slots_fortune_snake/pkg"
	"github.com/redis/go-redis/v9"
)

type service struct {
	playerRepo playerRepo.Repository
}

func NewService(playerRepo playerRepo.Repository) Service {
	return &service{playerRepo: playerRepo}
}

func (s *service) GetPlayerSession(ctx *server.Context, playerId uint64) (output *playerModel.Session, err error) {
	grpcCtx := ctx.MustGet("ctx").(context.Context)
	data, err := s.playerRepo.FindPlayerSessionById(grpcCtx, playerId)
	if err != nil {
		var e *errMsg.Error
		if errors.Is(err, redis.Nil) {
			e = errMsg.New(constants.CodeUnauthorized, pkg.LocalizeInstance().LocalizeMessage("TokenExpired"), err)
			return nil, e
		}
		e = errMsg.New(constants.CodeBadRequest, pkg.LocalizeInstance().LocalizeMessage("RedisError", map[string]interface{}{"err": err}), err)
		return nil, e
	}
	return data, nil
}
