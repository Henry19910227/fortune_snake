package snow_flake

type Repository interface {
	GenerateID() uint64
}
