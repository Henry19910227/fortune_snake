package config

import (
	"game_server_slots_fortune_snake/model"
)

type Config interface {
	Config() *model.Config
}
