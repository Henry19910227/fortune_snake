package settle

type repository struct {
	hitLines [][]int
}

func New() Repository {
	hitLines := [][]int{
		{0, 0, 0}, // Line 1
		{0, 1, 0}, // Line 2
		{0, 1, 1}, // Line 3
		{1, 1, 0}, // Line 4
		{1, 1, 1}, // Line 5
		{1, 2, 1}, // Line 6
		{1, 2, 2}, // Line 7
		{2, 2, 1}, // Line 8
		{2, 2, 2}, // Line 9
		{2, 3, 2}, // Line 10
	}
	return &repository{hitLines: hitLines}
}

func (r *repository) HitLines() [][]int {
	return r.hitLines
}
