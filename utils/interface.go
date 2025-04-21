package utils

type SnowFlake interface {
	GenerateID() (uint64, error)
}
