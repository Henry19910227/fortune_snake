package game_result

import (
	betModel "game_server_slots_fortune_snake/internal/model/controller/game/bet"
	model "game_server_slots_fortune_snake/internal/model/entity/game_result"
	playerModel "game_server_slots_fortune_snake/internal/model/entity/player"
	gameResultRepo "game_server_slots_fortune_snake/internal/repository/game_result"
	snowFlakeRepo "game_server_slots_fortune_snake/internal/repository/snow_flake"
	"time"
)

type service struct {
	gameResultRepo gameResultRepo.Repository
	snowFlakeRepo  snowFlakeRepo.Repository
}

func New(gameResultRepo gameResultRepo.Repository, snowFlakeRepo snowFlakeRepo.Repository) Service {
	return &service{gameResultRepo: gameResultRepo, snowFlakeRepo: snowFlakeRepo}
}

func (s *service) Create(session *playerModel.Session, result *betModel.GameResult) (id uint64, err error) {
	item := &model.Table{}
	item.ID = s.snowFlakeRepo.GenerateID()
	item.RoundId = s.snowFlakeRepo.GenerateID()
	item.GameCode = session.GameCode
	item.RTP = session.GameRtp
	item.Result = result.Encode()
	item.Status = "yes"
	item.CreatedAt = time.Now().UnixMilli()
	item.UpdatedAt = time.Now().UnixMilli()
	return s.gameResultRepo.Create(item)
}
