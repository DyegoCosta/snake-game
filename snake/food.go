package snake

type Food struct {
	Points int
	X      int
	Y      int
}

func NewFood(x, y int) *Food {
	return &Food{
		Points: 10,
		X:      x,
		Y:      y,
	}
}
