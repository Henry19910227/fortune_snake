package game

import gameModel "game_server_slots_fortune_snake/internal/model/game"

type Repository interface {
	// Info 獲取遊戲配置數據
	Info() (output *gameModel.Info, err error)
}
