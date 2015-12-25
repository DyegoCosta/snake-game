package main

import (
	"time"

	"github.com/nsf/termbox-go"
)

var (
	pointsChan         = make(chan int)
	keyboardEventsChan = make(chan KeyboardEvent)
)

type Game struct {
	Arena  *Arena
	Score  int
	IsOver bool
}

func initialSnake() *Snake {
	return newSnake(RIGHT, []Coord{
		Coord{X: 1, Y: 1},
		Coord{X: 1, Y: 2},
		Coord{X: 1, Y: 3},
		Coord{X: 1, Y: 4},
	})
}

func initialScore() int {
	return 0
}

func initialArena() *Arena {
	return newArena(initialSnake(), pointsChan, 20, 50)
}

func (g *Game) end() {
	g.IsOver = true
}

func (g *Game) moveInterval() time.Duration {
	ms := 100 - (g.Score / 10)
	return time.Duration(ms) * time.Millisecond
}

func (g *Game) retry() {
	g.Arena = initialArena()
	g.Score = initialScore()
	g.IsOver = false
}

func (g *Game) addPoints(p int) {
	g.Score += p
}

func NewGame() *Game {
	return &Game{Arena: initialArena(), Score: initialScore()}
}

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
			switch e.Type {
			case MOVE:
				d := keyToDirection(e.Key)
				g.Arena.Snake.changeDirection(d)
			case RETRY:
				g.retry()
			case END:
				break mainloop
			}
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
