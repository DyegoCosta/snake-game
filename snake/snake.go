package snake

import "errors"

const (
	RIGHT Direction = 1 + iota
	LEFT
	UP
	DOWN
)

type Direction int

type Snake struct {
	Body      []Coord
	Direction Direction
	Length    int
}

func newSnake(d Direction, b []Coord) *Snake {
	return &Snake{
		Length:    len(b),
		Body:      b,
		Direction: d,
	}
}

func (s *Snake) changeDirection(d Direction) {
	opposites := map[Direction]Direction{
		RIGHT: LEFT,
		LEFT:  RIGHT,
		UP:    DOWN,
		DOWN:  UP,
	}

	if o := opposites[d]; o != 0 && o != s.Direction {
		s.Direction = d
	}
}

func (s *Snake) head() Coord {
	return s.Body[len(s.Body)-1]
}

func (s *Snake) die() error {
	return errors.New("Died")
}

func (s *Snake) move() error {
	h := s.head()
	c := Coord{X: h.X, Y: h.Y}

	switch s.Direction {
	case RIGHT:
		c.X += 1
	case LEFT:
		c.X -= 1
	case UP:
		c.Y += 1
	case DOWN:
		c.Y -= 1
	}

	if s.isOnPosition(c) {
		return s.die()
	}

	if s.Length > len(s.Body) {
		s.Body = append(s.Body, c)
	} else {
		s.Body = append(s.Body[1:], c)
	}

	return nil
}

func (s *Snake) isOnPosition(c Coord) bool {
	for _, b := range s.Body {
		if b.X == c.X && b.Y == c.Y {
			return true
		}
	}

	return false
}
