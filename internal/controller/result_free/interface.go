package result_free

type Controller interface {
	// Generate 生成盤面並存到Bucket(跑庫專用)
	Generate()
}
