package snake

import "math/rand"

type Arena struct {
	Food   *Food
	Snake  *Snake
	Height int
	Width  int
	points chan (int)
}

func newArena(s *Snake, h, w int) *Arena {
	a := &Arena{
		Snake:  s,
		Height: h,
		Width:  w,
	}

	a.placeFood()

	return a
}

func (a *Arena) moveSnake() error {
	if err := a.Snake.move(); err != nil {
		return err
	}

	if a.snakeLeftArena() {
		return a.Snake.die()
	}

	if a.hasFood(a.Snake.head()) {
		a.addPoints(a.Food.Points)
		a.Snake.Length++
		a.placeFood()
	}

	return nil
}

func (a *Arena) snakeLeftArena() bool {
	h := a.Snake.head()
	return h[0] > a.Width || h[1] > a.Height || h[0] < 0 || h[1] < 0
}

func (a *Arena) addPoints(p int) {
	a.points <- p
}

func (a *Arena) placeFood() {
	var x, y int

	for {
		x = rand.Intn(a.Width)
		y = rand.Intn(a.Height)

		if !a.isOccupied([]int{x, y}) {
			break
		}
	}

	a.Food = NewFood(x, y)
}

func (a *Arena) hasFood(p []int) bool {
	return p[0] == a.Food.X && p[1] == a.Food.Y
}

func (a *Arena) isOccupied(p []int) bool {
	return a.Snake.isOnPosition(p)
}
