package system

import (
	model "game_server_slots_fortune_snake/internal/model/config/system"
)

type Config interface {
	Config() *model.Config
}
