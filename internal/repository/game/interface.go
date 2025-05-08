package game

import (
	"game_server_slots_fortune_snake/internal/model/repository/game/info"
)

type Repository interface {
	// Info 獲取遊戲配置數據
	Info(input *info.Input) (output *info.Output, err error)
	// IsSpecialMode 是否是特殊模式
	IsSpecialMode() bool
}
