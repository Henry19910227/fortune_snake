package enter_game

import (
	"game_server_slots_fortune_snake/internal/model"
	playerModel "game_server_slots_fortune_snake/internal/model/player"
)

type Input struct {
	model.BaseInput
	Session *playerModel.Session
}

type Output struct {
	model.BaseOutput
	Data *Data
}

type Data struct {
	Bets               []int     // 可投注金额选项
	Values             []float64 // 可投注价值选项
	Multipler          int       // 投注线
	IsNumeric          bool      // 是否数值模式
	GameMode           string    // 游戏模式 试玩 demo, 真实用户 real
	ScoreTry           int       // 玩家的试玩余额
	MultipleScoreLimit int       // 奖金上限倍数
	Bet                int       // 当前下注值
	Value              float64   // 当前下注金额
	LastGameResult     GameResult
}

type GameResult struct {
	SpinResults interface{}
	TotalScore  int     //当前中奖总积分
	IsCutShort  bool    //是否限制了最高赔付
	WinType     int     //输赢的类型 0.Lose  1.Win  2.Big Win  3.Mage Win  4.Super Win
	Revenue     int     //返水
	SpinMode    int     //旋转模式 0普通 1福牛
	WinRate     float64 //赔率
}
