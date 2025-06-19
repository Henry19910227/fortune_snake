package player_free_order

import (
	"encoding/json"
	"game_server_slots_fortune_snake/internal/model/entity/player"
)

type Table struct {
	ID                 uint64          `gorm:"column:id" json:"id"`                                     // 唯一 ID（TiDB 分布式自增 ID）
	TransactionId      uint64          `gorm:"column:transaction_id" json:"transaction_id"`             // 交易单号
	RoundId            uint64          `gorm:"column:round_id" json:"round_id"`                         // 局号
	PlayerId           uint64          `gorm:"column:player_id" json:"player_id"`                       // 玩家 ID
	PlayerUsername     string          `gorm:"column:player_username" json:"player_username"`           // 玩家用户名
	MerchantId         uint64          `gorm:"column:merchant_id" json:"merchant_id"`                   // 商户 ID
	MerchantUsername   string          `gorm:"column:merchant_username" json:"merchant_username"`       // 商户账号
	MerchantName       string          `gorm:"column:merchant_name" json:"merchant_name"`               // 商户名
	GameId             uint64          `gorm:"column:game_id" json:"game_id"`                           // 游戏 ID
	GameCode           string          `gorm:"column:game_code" json:"game_code"`                       // 游戏代码
	GameFreeId         uint64          `gorm:"column:game_free_id" json:"game_free_id"`                 // 免费游戏 ID
	GameFreeItems      int8            `gorm:"column:game_free_items" json:"game_free_items"`           // 免费次数
	GameFreeMultiplier int8            `gorm:"column:game_free_multiplier" json:"game_free_multiplier"` // 免费倍数
	GameFreeUsedItems  int8            `gorm:"column:game_free_used_items" json:"game_free_used_items"` // 已使用免费次数
	CurrencyId         int             `gorm:"column:currency_id" json:"currency_id"`                   // 币种 ID
	CurrencyCode       string          `gorm:"column:currency_code" json:"currency_code"`               // 币种名称
	CurrencySymbol     string          `gorm:"column:currency_symbol" json:"currency_symbol"`           // 币种符号
	CurrencyExchange   int             `gorm:"column:currency_exchange" json:"currency_exchange"`       // 币种兑换比例
	Amount             int64           `gorm:"column:amount" json:"amount"`                             // 金额
	AmountTotal        int64           `gorm:"column:amount_total" json:"amount_total"`                 // 总金额
	Status             string          `gorm:"column:status" json:"status"`                             // 状态
	FreeGameJson       json.RawMessage `gorm:"column:free_game_json" json:"free_game_json"`             // 免费游戏 JSON 数据
	FreeType           string          `gorm:"column:free_type" json:"free_type"`                       // 免费类型
	IsSync             string          `gorm:"column:is_sync" json:"is_sync"`                           // 是否需要同步
	Sync               string          `gorm:"column:sync" json:"sync"`                                 // 同步状态
	CreatedAt          int64           `gorm:"column:created_at" json:"created_at"`                     // 创建时间（毫秒级时间戳）
	UpdatedAt          int64           `gorm:"column:updated_at" json:"updated_at"`                     // 更新时间（毫秒级时间戳）
}

func New(session *player.Session) *Table {
	table := &Table{}
	table.PlayerId = session.PlayerId
	table.PlayerUsername = session.PlayerUsername
	table.MerchantId = session.MerchantId
	table.MerchantUsername = session.MerchantUsername
	table.MerchantName = session.MerchantName
	table.GameId = session.GameId
	table.GameCode = session.GameCode
	table.GameFreeId = session.GameId
	table.CurrencyId = session.CurrencyId
	table.CurrencyCode = session.CurrencyCode
	table.CurrencySymbol = session.CurrencySymbol
	table.CurrencyExchange = session.CurrencyExchange
	return table
}

func (Table) TableName() string {
	return "player_free_order"
}
