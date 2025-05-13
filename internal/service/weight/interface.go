package weight

type Service interface {
	Load()
	RandomBaseWeightRate(rtp float64) (float64, error)
	RandomBaseWeightHRate(rtp float64) (float64, error)
	RandomFreeWeightRate(rtp float64) (float64, error)
	RandomFreeWeightHRate(rtp float64) (float64, error)
}
