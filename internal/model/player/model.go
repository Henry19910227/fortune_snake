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
	Value    int
	GameMode int // 游戏模式(試玩:demo/真玩:real)
}
