package symbol

type Response struct {
	Symbols []*Item `json:"symbols"`
}

type Item struct {
	ID   int    `json:"id"`   // 編號
	Name string `json:"name"` // 圖案名
	Pow  int    `json:"pow"`  // 倍率
}
