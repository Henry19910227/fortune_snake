package game

import (
	model "game_server_slots_fortune_snake/internal/model/config/bucket"
	"game_server_slots_fortune_snake/internal/model/config/symbol"
)

type Config interface {
	BaseBucketConfig() []*model.Item
	FreeBucketConfig() []*model.Item
	BaseSymbolConfig() []*symbol.Item
	FreeSymbolConfig() []*symbol.Item
}
