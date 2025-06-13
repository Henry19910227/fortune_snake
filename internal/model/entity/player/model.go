package player

type Session struct {
	PlayerId         uint64 `json:"player_id"`          // 玩家Id(主键)
	PlayerUsername   string `json:"player_username"`    // 第三方玩家账号
	GameAt           int64  `json:"game_at"`            // 首次游戏时间（毫秒级时间戳）
	GameIp           []byte `json:"game_ip"`            // 首次游戏Ip（支持Ipv4/Ipv6）
	GuestMode        bool   `json:"guest_mode"`         // 访客模式
	Balance          int64  `json:"balance"`            // 账户余额
	WalletMode       string `json:"wallet_mode"`        // 钱包模式
	Auth             string `json:"auth"`               // 授权
	CurrencyId       int    `json:"currency_id"`        // 货币
	CurrencyCode     string `json:"currency_code"`      // 货币代码
	CurrencySymbol   string `json:"currency_symbol"`    // 货币符号
	CurrencyExchange int    `json:"currency_exchange" ` // 币种兑换比例
	Mode             string `json:"mode"`               // 游戏模式
	Language         string `json:"language"`           // 语言
	MerchantId       uint64 `json:"merchant_id"`        // 商户Id
	MerchantUsername string `json:"merchant_username"`  // 商户账号
	MerchantName     string `json:"merchant_name"`      // 商户名
	GameId           uint64 `json:"game_id"`            // 游戏Id
	GameCode         string `json:"game_code"`          // 游戏代码
	GameRtp          int8   `json:"game_rtp"`           // 游戏rtp值
	GameteBetType    string `json:"game_bet_type"`      // 游戏投注额验证类型；match 一致 range 区间
	GameDatetime     int64  `json:"game_datetime"`      // 游戏启动时间
}

type Table struct {
	Id               uint64 `json:"id" db:"id"`                               // 玩家Id（TiDB分布式Id）
	Username         string `json:"username" db:"username"`                   // 商户的用户名/商户的用户Id
	MerchantId       uint64 `json:"merchant_id" db:"merchant_id"`             // 商户id
	MerchantUsername string `json:"merchant_username" db:"merchant_username"` // 商户账号
	MerchantName     string `json:"merchant_name" db:"merchant_name"`         // 商户名
	CurrencyId       int    `json:"currency_id" db:"currency_id"`             // 币种id
	CurrencyCode     string `json:"currency_code" db:"currency_code"`         // 币种名称
	CurrencySymbol   string `json:"currency_symbol" db:"currency_symbol"`     // 币种符号
	CurrencyExchange int    `json:"currency_exchange" db:"currency_exchange"` // 币种兑换比例
	WalletMode       string `json:"wallet_mode" db:"wallet_mode"`             // 钱包模式（transfer: 转账, synchronous: 同步）
	Balance          int64  `json:"balance" db:"balance"`                     // 主账户余额
	BetTotal         int64  `json:"bet_total" db:"bet_total"`                 // 总投注金额
	WinTotal         int64  `json:"win_total" db:"win_total"`                 // 总输赢
	Status           string `json:"status" db:"status"`                       // 状态（yes: 启用, no: 禁用）
	GameAt           int64  `json:"game_at" db:"game_at"`                     // 首次游戏时间（毫秒级时间戳）
	GameIP           []byte `json:"game_ip" db:"game_ip"`                     // 首次游戏IP（支持IPv4/IPv6）
	CreatedAt        int64  `json:"created_at" db:"created_at"`               // 创建时间（毫秒级时间戳）
	CreatedIP        []byte `json:"created_ip" db:"created_ip"`               // 创建IP（支持IPv4/IPv6）
	UpdatedAt        int64  `json:"updated_at" db:"updated_at"`               // 更新时间（毫秒级时间戳）
	UpdatedIP        []byte `json:"updated_ip" db:"updated_ip"`               // 更新IP（支持IPv4/IPv6）
}

func (*Table) TableName() string {
	return "player"
}

// User 用戶數據
type User struct {
	Id       int
	NickName string
	Sex      int
	Currency string
	RTP      int
}

// GameData 用戶遊戲中緩存數據
type GameData struct {
	UserId   int // 用戶 id
	Bet      int
	Value    float64
	GameMode int // 游戏模式(試玩:demo/真玩:real)
}
