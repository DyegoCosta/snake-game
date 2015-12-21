package snake

import (
	"time"

	"github.com/nsf/termbox-go"
)

var (
	pointsChan  = make(chan int)
	movesChan   = make(chan int)
	endGameChan = make(chan bool)
)

type Game struct {
	Arena  *Arena
	Score  int
	IsOver bool
}

func NewGame() *Game {
	s := newSnake(RIGHT, [][]int{{1, 1}, {1, 2}, {1, 3}, {1, 4}})
	a := newArena(s, pointsChan, 20, 50)
	return &Game{Arena: a, Score: 0}
}

func (g *Game) end() {
	g.IsOver = true
}

func (g *Game) moveInterval() time.Duration {
	ms := 100 - (g.Score / 10)
	return time.Duration(ms) * time.Millisecond
}

func initTermbox() {
	if err := termbox.Init(); err != nil {
		panic(err)
	}
}

func (g *Game) Start() {
	initTermbox()
	defer termbox.Close()

	go listenToKeyboard(movesChan, endGameChan)

	g.render()

mainloop:
	for {
		select {
		case d := <-movesChan:
			g.Arena.Snake.changeDirection(d)
		case p := <-pointsChan:
			g.Score += p
		case <-endGameChan:
			break mainloop
		default:
			if !g.IsOver {
				if err := g.Arena.moveSnake(); err != nil {
					g.end()
				}
			}

			if err := g.render(); err != nil {
				panic(err)
			}

			time.Sleep(g.moveInterval())
		}
	}
}
