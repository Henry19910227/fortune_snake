package weight

type Stat struct {
	Rate          float64 `json:"rate"`            //赔率
	FromWeights   []int64 `json:"from_weights"`    //权重开始（包含）[]int64为切片的序列化，为RTP96、92、88等扩展
	WeightAccount []int64 `json:"weights_account"` //权重值 []int64 为切片的序列化，为RTP96、92、88等扩展
	ToWeights     []int64 `json:"to_weights"`      //权重结束（不包含）[]int64 为切片的序列化，为RTP96、92、88等扩展
}
