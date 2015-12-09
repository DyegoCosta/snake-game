package snake

const (
	RIGHT = iota
	LEFT
	UP
	DOWN
)

type Snake struct {
	Body      [][]int
	Direction int
	Alive     bool
}

func NewSnake(d int, b [][]int) Snake {
	return Snake{
		Alive:     true,
		Body:      b,
		Direction: d,
	}
}

func (s *Snake) ChangeDirection(d int) {
	oposity := map[int]int{RIGHT: LEFT, LEFT: RIGHT, UP: DOWN, DOWN: UP}

	if oposity[s.Direction] != d {
		s.Direction = d
	}
}

func (s *Snake) Head() []int {
	return s.Body[len(s.Body)-1]
}

func (s *Snake) Die() {
	s.Alive = false
}

func (s *Snake) Move() {
	h := make([]int, 2)
	copy(h, s.Head())

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

	if s.onTopOfItself(h) {
		s.Die()
	}

	s.Body = append(s.Body[1:], h)
}

func (s *Snake) onTopOfItself(h []int) bool {
	for _, p := range s.Body {
		if p[0] == h[0] && p[1] == h[1] {
			return true
		}
	}

	return false
}
