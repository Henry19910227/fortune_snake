package bucket

type Item struct {
	LowerLimit  float64 //赔率下限(不包含)
	UpperLimit  float64 //赔率上限(包含)
	MaxCapacity int     //最大容量（存储上限）
}
