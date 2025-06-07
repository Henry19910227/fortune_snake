package constants

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
	SpinModeNormal = 0
	SpinModeFree   = 1
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

// CacheNameRocketCrashBetsData 小火箭投注记录(区分不同的RTP)
const CacheNameRocketCrashBetsData = "Bets:RocketCrash:%d"

// CacheNameChannelSettlement 游戏结算的发布订阅
const CacheNameChannelSettlement = "Channel_Settlement_Game_%d"

// CacheNameSpinMode 該用戶當前旋轉模式
const CacheNameSpinMode = "Player:%v:FortuneSnake:SpinMode"

// CacheNameCurrentFreeTimes 該用戶剩餘免費旋轉次數
const CacheNameCurrentFreeTimes = "Player:%v:FortuneSnake:CurrentFreeTimes"

// CacheNameFreeResults 該用戶剩餘免費盤面
const CacheNameFreeResults = "Player:%v:FortuneSnake:%v:FreeResults"

// SpecialProbability 进入特殊模式的概率
const SpecialProbability = 0.008 // 进入特殊模式的概率
