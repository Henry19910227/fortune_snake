package result

type Item struct {
	ID         int64   `gorm:"column:id"`          // id
	Rate       float64 `gorm:"column:rate"`        // 赔率
	Symbols    string  `gorm:"column:symbols"`     // 盤面數據
	AllSymbols string  `gorm:"column:all_symbols"` // 免費模式盤面數據
	CreatedAt  int64   `gorm:"column:created_at"`
	UpdatedAt  int64   `gorm:"column:updated_at"`
}

func (Item) TableName() string {
	return "fortune_snake_results"
}
