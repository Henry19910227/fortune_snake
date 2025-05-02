package result

type TableItem struct {
	ID        int64   `gorm:"column:id"`  // id
	Rate      float64 `gorm:"rate"`       //赔率
	RateIndex int     `gorm:"rate_index"` //排序
	Symbols   string  `gorm:"symbols"`    //中奖结果
}

func (TableItem) TableName() string {
	return "result"
}
