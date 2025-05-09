package game

import model "game_server_slots_fortune_snake/internal/model/config/bucket"

type Config interface {
	BaseBucketConfig() []*model.Item
	FreeBucketConfig() []*model.Item
}
