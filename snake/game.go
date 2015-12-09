package snake

type Game struct {
	Snake Snake
	Arena Arena
	Score int
}

func NewGame() Game {
	s := Snake{
		Direction: RIGHT,
		Body: [][]int{
			{1, 1},
			{1, 2},
			{1, 3},
			{1, 4},
		},
	}

	a := Arena{
		Height: 20,
		Width:  20,
	}

	return Game{
		Arena: a,
		Snake: s,
		Score: 0,
	}
}

func (g *Game) Start() {
	for {
		g.Snake.Move()
	}
}
