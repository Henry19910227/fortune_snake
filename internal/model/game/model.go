package game

// Info 遊戲數據
type Info struct {
	GameID    string
	Bets      []int     //可投注金额选项
	Values    []float64 //可投注价值选项
	Multipler int       //投注线
}
