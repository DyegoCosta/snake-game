package snake

import (
	"fmt"
	"time"

	"github.com/nsf/termbox-go"
)

var pointsChan = make(chan int)

type Game struct {
	Arena *Arena
	Score int
}

func NewGame() *Game {
	s := newSnake(RIGHT, [][]int{{1, 1}, {1, 2}, {1, 3}, {1, 4}})
	a := newArena(s, pointsChan, 20, 50)
	return &Game{Arena: a, Score: 0}
}

func (g *Game) end() {
	fmt.Println("Game Over")
	fmt.Printf("Score: %v", g.Score)
}

func (g *Game) handlePointsReceived() {
	for {
		if p, ok := <-pointsChan; ok {
			g.Score += p
		}
	}
}

func (g *Game) moveInterval() time.Duration {
	ms := 100 - (g.Score / 10)
	return time.Duration(ms) * time.Millisecond
}

func initTermbox() {
	if termbox.IsInit {
		return
	} else if err := termbox.Init(); err != nil {
		panic(err)
	}
}

func (g *Game) Start() {
	initTermbox()
	defer termbox.Close()

	go g.handlePointsReceived()
	defer close(pointsChan)

	go g.startKeyboardArrowsListener()

	g.render()

mainloop:
	for {
		if err := g.Arena.moveSnake(); err != nil {
			g.end()
			break mainloop
		}

		if err := g.render(); err != nil {
			panic(err)
		}

		time.Sleep(g.moveInterval())
	}
}
