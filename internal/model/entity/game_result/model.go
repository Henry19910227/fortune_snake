package game_result

import "encoding/json"

type Table struct {
	ID        uint64          `gorm:"column:id" json:"id"`                 // 唯一 Id（TiDB 分布式 Id）
	RoundId   uint64          `gorm:"column:round_id" json:"round_id"`     // 游戏局号
	GameId    uint64          `gorm:"column:game_id" json:"game_id"`       // 游戏 Id
	GameCode  string          `gorm:"column:game_code" json:"game_code"`   // 游戏代码
	RTP       int8            `gorm:"column:rtp" json:"rtp"`               // RTP（理论返还率）
	Result    json.RawMessage `gorm:"column:result" json:"result"`         // 开奖结果（JSON 格式）
	Status    string          `gorm:"column:status" json:"status"`         // 状态（yes, no）
	CreatedAt int64           `gorm:"column:created_at" json:"created_at"` // 创建时间（毫秒级时间戳）
	UpdatedAt int64           `gorm:"column:updated_at" json:"updated_at"` // 更新时间（毫秒级时间戳）
}

func (Table) TableName() string {
	return "game_result"
}
