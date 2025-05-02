package config

import (
	"game_server_slots_fortune_snake/internal/model"
)

type Config interface {
	Config() *model.Config
	BaseResultConfig() []*model.ResultConfig
}
