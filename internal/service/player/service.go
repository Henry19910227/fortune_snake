package player

import (
	"context"
	"errors"
	"game_server_slots_fortune_snake/constants"
	playerModel "game_server_slots_fortune_snake/internal/model/entity/player"
	errMsg "game_server_slots_fortune_snake/internal/model/err"
	playerRepo "game_server_slots_fortune_snake/internal/repository/player"
	sessionRepo "game_server_slots_fortune_snake/internal/repository/player_session"
	"game_server_slots_fortune_snake/internal/server"
	"game_server_slots_fortune_snake/pkg"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type service struct {
	playerRepo  playerRepo.Repository
	sessionRepo sessionRepo.Repository
}

func New(playerRepo playerRepo.Repository, sessionRepo sessionRepo.Repository) Service {
	return &service{playerRepo: playerRepo, sessionRepo: sessionRepo}
}

func (s *service) Tx(tx *gorm.DB) Service {
	return s.txService(tx)
}

func (s *service) GetPlayerSession(ctx *server.Context, playerId uint64) (output *playerModel.Session, err error) {
	grpcCtx := ctx.MustGet("ctx").(context.Context)
	data, err := s.sessionRepo.FindPlayerSessionById(grpcCtx, playerId)
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

func (s *service) UpdateBalance(ctx *server.Context, playerId uint64, value int64) error {
	grpcCtx := ctx.MustGet("ctx").(context.Context)
	session, err := s.GetPlayerSession(ctx, playerId)
	if err != nil {
		return err
	}
	session.Balance = value

	// 更新 session
	if err := s.sessionRepo.UpdateSessionById(grpcCtx, playerId, session); err != nil {
		return err
	}
	if session.Mode == "demo" {
		return nil
	}
	// 轉帳模式改table
	if session.WalletMode == "transfer" {
		if err := s.playerRepo.UpdateBalance(playerId, value); err != nil {
			return err
		}
	}
	// 無縫模式走redis發布訂閱

	return nil
}

func (s *service) txService(tx *gorm.DB) Service {
	return &service{playerRepo: playerRepo.New(tx), sessionRepo: s.sessionRepo}
}
