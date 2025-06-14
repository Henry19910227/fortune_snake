package bet_record

import (
	model "game_server_slots_fortune_snake/internal/model/entity/bet_record"
	playerModel "game_server_slots_fortune_snake/internal/model/entity/player"
	betRecordRepo "game_server_slots_fortune_snake/internal/repository/bet_record"
	snowFlakeRepo "game_server_slots_fortune_snake/internal/repository/snow_flake"
	"time"
)

type service struct {
	betRecordRepo betRecordRepo.Repository
	snowFlakeRepo snowFlakeRepo.Repository
}

func New(betRecordRepo betRecordRepo.Repository, snowFlakeRepo snowFlakeRepo.Repository) Service {
	return &service{betRecordRepo: betRecordRepo, snowFlakeRepo: snowFlakeRepo}
}

func (s *service) Create(session *playerModel.Session) (item *model.Table, err error) {
	TransactionId := s.snowFlakeRepo.GenerateID()
	item = &model.Table{}
	item.TransactionId = TransactionId
	item.TransactionSubId = TransactionId
	item.RoundId = s.snowFlakeRepo.GenerateID()
	item.PlayerId = session.PlayerId
	item.PlayerUsername = session.PlayerUsername
	item.MerchantId = session.MerchantId
	item.MerchantUsername = session.MerchantUsername
	item.MerchantName = session.MerchantName
	item.GameId = session.GameId
	item.GameCode = session.GameCode
	item.RTP = session.GameRtp
	item.CurrencyId = session.CurrencyId
	item.CurrencyCode = session.CurrencyCode
	item.CurrencySymbol = session.CurrencySymbol
	item.CurrencyExchange = session.CurrencyExchange
	item.Balance = session.Balance
	item.Mode = session.Mode
	item.Free = "no"
	item.Special = "no"
	item.Jackpot = "no"
	item.Expand = "no"
	item.Status = "bet"
	item.Sync = "no"
	item.CreatedAt = time.Now().UnixMilli()
	item.UpdatedAt = time.Now().UnixMilli()
	_, err = s.betRecordRepo.Create(item)
	if err != nil {
		return nil, err
	}
	return item, nil
}

func (s *service) Update(item *model.Table) (err error) {
	item.UpdatedAt = time.Now().UnixMilli()
	return s.betRecordRepo.Update(item)
}
