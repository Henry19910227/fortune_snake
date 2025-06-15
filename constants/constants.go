package constants

import "time"

// 定义业务错误码常量
const (
	CodeSuccess          = 200 // 成功
	CodeBadRequest       = 400 // 参数错误
	CodeUnauthorized     = 401 // 未授权
	CodeMethodNotAllowed = 405 // 错误的请求方式
	CodeForbidden        = 403 // 禁止访问
	CodeTooManyRequests  = 429 // API访问频繁
	CodeNotFound         = 404 // 资源不存在
	CodeInternalError    = 500 // 服务器内部错误
)

const (
	GameModeDemo = "demo"
	GameModeReal = "real"
)

const (
	SpinModeBase = 0
	SpinModeFree = 1
)

const (
	BaseModeInReal      = "BaseModeInReal"
	StartFreeModeInReal = "StartFreeModeInReal"
	FreeModeInReal      = "FreeModeInReal"
	FinalFreeModeInReal = "FinalFreeModeInReal"

	BaseModeInDemo      = "BaseModeInDemo"
	StartFreeModeInDemo = "StartFreeModeInDemo"
	FreeModeInDemo      = "FreeModeInDemo"
	FinalFreeModeInDemo = "FinalFreeModeInDemo"
)

const (
	SYMBOL_0 int = 0 + iota // 百搭 0  ====》200
	SYMBOL_1                // 元宝 1  ====》100
	SYMBOL_2                // 福箱 2  ====》50
	SYMBOL_3                // 福袋 3  ====》20
	SYMBOL_4                // 红包 4  ====》10
	SYMBOL_5                // 橘子 5  ====》5
	SYMBOL_6                // 鞭炮 6  ====》3
)

// 缓存命名

// CacheNamePlayerSession 定义玩家会话 key 名称
const CacheNamePlayerSession = "Player:Session:%d"

// CacheNameChannelSettlement 游戏结算的发布订阅
const CacheNameChannelSettlement = "Channel_Settlement_Game_%d"

// CacheNamePlayerFreeResults 該用戶剩餘免費盤面 Player:FreeResults:GameMode:PlayerID
const CacheNamePlayerFreeResults = "Player:FreeResults:%v:%v"

// CacheNamePlayerGameInfo 該用戶續存數據 Player:GameInfo:GameMode:PlayerID
const CacheNamePlayerGameInfo = "Player:GameInfo:%v:%v"

// CacheExpiredPlayerFreeResults 剩餘免費盤面過期時間
const CacheExpiredPlayerFreeResults = 20 * 24 * time.Hour

// CacheExpiredPlayerSession 玩家Session過期時間
const CacheExpiredPlayerSession = 60 * time.Minute

// CacheExpiredPlayerGameInfo 玩家續存數據過期時間
const CacheExpiredPlayerGameInfo = 20 * 24 * time.Hour

// SpecialProbability 进入特殊模式的概率
const SpecialProbability = 0.008 // 进入特殊模式的概率
