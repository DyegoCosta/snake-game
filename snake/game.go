package snake

type Game struct {
	Snake *Snake
	Arena *Arena
	Score int
}

func NewGame() *Game {
	s := newSnake(RIGHT, [][]int{{1, 1}, {1, 2}, {1, 3}, {1, 4}})
	a := newArena(s, 20, 20)
	return &Game{Arena: a, Score: 0}
}

func (g *Game) end() {
}

func (g *Game) Start() {
	g.Score += <-g.Arena.points

	for {
		if err := g.Arena.moveSnake(); err != nil {
			g.end()
		}
	}
}
