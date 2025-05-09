package player

import (
	"context"
	"errors"
	"game_server_slots_fortune_snake/constants"
	errMsg "game_server_slots_fortune_snake/internal/model/err"
	"game_server_slots_fortune_snake/internal/model/service/player/get_player_session"
	playerRepo "game_server_slots_fortune_snake/internal/repository/player"
	"game_server_slots_fortune_snake/pkg"
	"github.com/redis/go-redis/v9"
)

type service struct {
	playerRepo playerRepo.Repository
}

func NewService(playerRepo playerRepo.Repository) Service {
	return &service{playerRepo: playerRepo}
}

func (s *service) GetPlayerSession(input *get_player_session.Input) (output *get_player_session.Output, err error) {
	grpcCtx := input.Ctx.MustGet("ctx").(context.Context)
	data, err := s.playerRepo.FindPlayerSessionById(grpcCtx, input.PlayerId)
	if err != nil {
		var e *errMsg.Error
		if errors.Is(err, redis.Nil) {
			e = errMsg.New(constants.CodeUnauthorized, pkg.LocalizeInstance().LocalizeMessage("TokenExpired"), err)
			return nil, e
		}
		e = errMsg.New(constants.CodeBadRequest, pkg.LocalizeInstance().LocalizeMessage("RedisError", map[string]interface{}{"err": err}), err)
		return nil, e
	}
	output = &get_player_session.Output{}
	output.Session = data
	return output, nil
}
