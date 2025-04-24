package player

import (
	"errors"
	"game_server_slots_fortune_snake/constants"
	"game_server_slots_fortune_snake/internal/model/player/get_player_session"
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

func (s *service) GetPlayerSession(input *get_player_session.Input) (output *get_player_session.Output) {
	data, err := s.playerRepo.FindPlayerSessionById(input.Ctx, input.PlayerId)
	if err != nil {
		if errors.Is(err, redis.Nil) {
			output.Code = constants.CodeUnauthorized
			output.Message = pkg.LocalizeInstance().LocalizeMessage("TokenExpired")
			return output
		}
		output.Code = constants.CodeBadRequest
		output.Message = pkg.LocalizeInstance().LocalizeMessage("RedisError", map[string]interface{}{"err": err})
	}
	output = &get_player_session.Output{}
	output.Code = constants.CodeSuccess
	output.Message = ""
	output.Data = data
	return output
}
