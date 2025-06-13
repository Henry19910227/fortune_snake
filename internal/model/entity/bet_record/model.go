package bet_record

import "encoding/json"

type Table struct {
	ID               uint64          `gorm:"column:id" json:"id"`                                 // 唯一 Id（TiDB 分布式 Id）
	TransactionId    uint64          `gorm:"column:transaction_id" json:"transaction_id"`         // 主交易单号
	TransactionSubId uint64          `gorm:"column:transaction_sub_id" json:"transaction_sub_id"` // 子交易单号
	RoundId          uint64          `gorm:"column:round_id" json:"round_id"`                     // 游戏局号
	PlayerId         uint64          `gorm:"column:player_id" json:"player_id"`                   // 玩家 Id
	PlayerUsername   string          `gorm:"column:player_username" json:"player_username"`       // 玩家用户名
	MerchantId       uint64          `gorm:"column:merchant_id" json:"merchant_id"`               // 商户 Id
	MerchantUsername string          `gorm:"column:merchant_username" json:"merchant_username"`   // 商户账号
	MerchantName     string          `gorm:"column:merchant_name" json:"merchant_name"`           // 商户名
	GameId           uint64          `gorm:"column:game_id" json:"game_id"`                       // 游戏 Id
	GameCode         string          `gorm:"column:game_code" json:"game_code"`                   // 游戏代码
	RTP              int8            `gorm:"column:rtp" json:"rtp"`                               // RTP（理论返还率）
	CurrencyId       int             `gorm:"column:currency_id" json:"currency_id"`               // 币种 Id
	CurrencyCode     string          `gorm:"column:currency_code" json:"currency_code"`           // 币种名称
	CurrencySymbol   string          `gorm:"column:currency_symbol" json:"currency_symbol"`       // 币种符号
	CurrencyExchange int             `gorm:"column:currency_exchange" json:"currency_exchange"`   // 币种兑换比例
	Amount           int64           `gorm:"column:amount" json:"amount"`                         // 投注金额
	AmountWin        int64           `gorm:"column:amount_win" json:"amount_win"`                 // 中奖金额
	Balance          int64           `gorm:"column:balance" json:"balance"`                       // 投注前余额
	Bet              json.RawMessage `gorm:"column:bet" json:"bet"`                               // 投注内容（JSON 格式）
	Result           json.RawMessage `gorm:"column:result" json:"result"`                         // 开奖结果（JSON 格式）
	Mode             string          `gorm:"column:mode" json:"mode"`                             // 投注模式（real/demo）
	Free             string          `gorm:"column:free" json:"free"`                             // 是否触发免费游戏（yes/no）
	Special          string          `gorm:"column:special" json:"special"`                       // 是否触发特殊奖励（yes/no）
	Jackpot          string          `gorm:"column:jackpot" json:"jackpot"`                       // 是否触发 Jackpot（yes/no）
	Expand           string          `gorm:"column:expand" json:"expand"`                         // 是否有扩展数据（yes/no）
	Status           string          `gorm:"column:status" json:"status"`                         // 状态（bet, cancel, hand, settlement, finish, error）
	Sync             string          `gorm:"column:sync" json:"sync"`                             // 是否同步三方（yes/no）
	CreatedAt        int64           `gorm:"column:created_at" json:"created_at"`                 // 创建时间（毫秒级时间戳）
	UpdatedAt        int64           `gorm:"column:updated_at" json:"updated_at"`                 // 更新时间（毫秒级时间戳）
}

func (Table) TableName() string {
	return "bet_record"
}
