package game

type repository struct {
}

func New() Repository {
	return &repository{}
}
