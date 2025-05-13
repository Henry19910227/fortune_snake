package weight

type Service interface {
	Load()
	GetRandomRate(rtp int)
}
