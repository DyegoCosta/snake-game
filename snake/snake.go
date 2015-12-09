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

	s.Body = append(s.Body[1:], h)
}
