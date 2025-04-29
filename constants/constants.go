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

// 缓存命名

// CacheNamePlayerSession 定义玩家会话 key 名称
const CacheNamePlayerSession = "Player:Session:%d"

// CacheNameRocketCrashBetsData 小火箭投注记录(区分不同的RTP)
const CacheNameRocketCrashBetsData = "Bets:RocketCrash:%d"

// CacheNameChannelSettlement 游戏结算的发布订阅
const CacheNameChannelSettlement = "Channel_Settlement_Game_%d"

// SpecialProbability 进入特殊模式的概率
const SpecialProbability = 0.008 // 进入特殊模式的概率
