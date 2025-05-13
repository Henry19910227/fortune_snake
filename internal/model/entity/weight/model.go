package weight

// Stat 盤面賠率權重表
type Stat struct {
	Rate          float64 //赔率
	FromWeights   []int64 //权重开始（包含）[]int64为切片的序列化，为RTP96、92、88等扩展
	WeightAccount []int64 //权重值 []int64 为切片的序列化，为RTP96、92、88等扩展
	ToWeights     []int64 //权重结束（不包含）[]int64 为切片的序列化，为RTP96、92、88等扩展
}
