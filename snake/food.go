package snake

type Food struct {
	Points int
	X      int
	Y      int
}

func NewFood(x, y int) *Food {
	return &Food{
		Points: 1,
		X:      x,
		Y:      y,
	}
}
