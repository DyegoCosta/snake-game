package snake

import "errors"

const (
	RIGHT = iota
	LEFT
	UP
	DOWN
)

type Snake struct {
	Body      [][]int
	Direction int
	Length    int
}

func newSnake(d int, b [][]int) *Snake {
	return &Snake{
		Length:    len(b),
		Body:      b,
		Direction: d,
	}
}

func (s *Snake) changeDirection(d int) {
	oposity := map[int]int{RIGHT: LEFT, LEFT: RIGHT, UP: DOWN, DOWN: UP}

	if oposity[s.Direction] != d {
		s.Direction = d
	}
}

func (s *Snake) head() []int {
	return s.Body[len(s.Body)-1]
}

func (s *Snake) die() error {
	return errors.New("Died")
}

func (s *Snake) move() error {
	h := make([]int, 2)
	copy(h, s.head())

	switch s.Direction {
	case RIGHT:
		h[0]++
	case LEFT:
		h[0]--
	case UP:
		h[1]++
	case DOWN:
		h[1]--
	}

	if s.isOnPosition(h) {
		return s.die()
	}

	s.Body = append(s.Body[1:], h)

	return nil
}

func (s *Snake) isOnPosition(h []int) bool {
	for _, p := range s.Body {
		if p[0] == h[0] && p[1] == h[1] {
			return true
		}
	}

	return false
}
