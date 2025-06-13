package player

type Repository interface {
	UpdateBalance(playerId uint64, newBalance int64) error
}
