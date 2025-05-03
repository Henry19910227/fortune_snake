package result

type Item struct {
	ID        int64   `gorm:"column:id"`         // id
	Rate      float64 `gorm:"column:rate"`       //赔率
	RateIndex int     `gorm:"column:rate_index"` //排序
	Symbols   string  `gorm:"column:symbols"`    //中奖结果
}

func (Item) TableName() string {
	return "result"
}
