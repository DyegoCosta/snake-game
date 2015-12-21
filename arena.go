package main

import "math/rand"

type Arena struct {
	Food       *Food
	Snake      *Snake
	hasFood    func(*Arena, Coord) bool
	Height     int
	Width      int
	pointsChan chan (int)
}

func newArena(s *Snake, p chan (int), h, w int) *Arena {
	a := &Arena{
		Snake:      s,
		Height:     h,
		Width:      w,
		pointsChan: p,
		hasFood:    hasFood,
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

	if a.hasFood(a, a.Snake.head()) {
		go a.addPoints(a.Food.Points)
		a.Snake.Length++
		a.placeFood()
	}

	return nil
}

func (a *Arena) snakeLeftArena() bool {
	h := a.Snake.head()
	return h.X > a.Width || h.Y > a.Height || h.X < 0 || h.Y < 0
}

func (a *Arena) addPoints(p int) {
	a.pointsChan <- p
}

func (a *Arena) placeFood() {
	var x, y int

	for {
		x = rand.Intn(a.Width)
		y = rand.Intn(a.Height)

		if !a.isOccupied(Coord{X: x, Y: y}) {
			break
		}
	}

	a.Food = NewFood(x, y)
}

func hasFood(a *Arena, c Coord) bool {
	return c.X == a.Food.X && c.Y == a.Food.Y
}

func (a *Arena) isOccupied(c Coord) bool {
	return a.Snake.isOnPosition(c)
}
