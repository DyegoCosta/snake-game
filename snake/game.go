package snake

import (
	"time"

	"github.com/nsf/termbox-go"
)

var (
	pointsChan         = make(chan int)
	keyboardEventsChan = make(chan keyboardEvent)
)

// Game type
type Game struct {
	arena  *arena
	score  int
	isOver bool
}

func initialSnake() *snake {
	return newSnake(RIGHT, []coord{
		coord{x: 1, y: 1},
		coord{x: 1, y: 2},
		coord{x: 1, y: 3},
		coord{x: 1, y: 4},
	})
}

func initialScore() int {
	return 0
}

func initialArena() *arena {
	return newArena(initialSnake(), pointsChan, 20, 50)
}

func (g *Game) end() {
	g.isOver = true
}

func (g *Game) moveInterval(d direction) time.Duration {
	ms := 100 - (g.score / 10)
	dur := time.Duration(ms) * time.Millisecond

	if d == UP || d == DOWN {
		return dur * 2
	}
	return dur
}

func (g *Game) retry() {
	g.arena = initialArena()
	g.score = initialScore()
	g.isOver = false
}

func (g *Game) addPoints(p int) {
	g.score += p
}

// NewGame creates new Game object
func NewGame() *Game {
	return &Game{arena: initialArena(), score: initialScore()}
}

// Start starts the game
func (g *Game) Start() {
	if err := termbox.Init(); err != nil {
		panic(err)
	}
	defer termbox.Close()

	go listenToKeyboard(keyboardEventsChan)

	if err := g.render(); err != nil {
		panic(err)
	}

mainloop:
	for {
		select {
		case p := <-pointsChan:
			g.addPoints(p)
		case e := <-keyboardEventsChan:
			switch e.eventType {
			case MOVE:
				d := keyToDirection(e.key)
				g.arena.snake.changeDirection(d)
			case RETRY:
				g.retry()
			case END:
				break mainloop
			}
		default:
			if !g.isOver {
				if err := g.arena.moveSnake(); err != nil {
					g.end()
				}
			}

			if err := g.render(); err != nil {
				panic(err)
			}

			time.Sleep(g.moveInterval(g.arena.snake.direction))
		}
	}
}
