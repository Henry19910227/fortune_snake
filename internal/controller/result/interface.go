package result

type Controller interface {
	// Generate 生成盤面並存到Bucket(跑庫專用)
	Generate()
	GenerateFree()
}
