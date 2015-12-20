package snake

import (
	"fmt"
	"time"

	"github.com/nsf/termbox-go"
)

var pointsChan = make(chan int)
var keyboardChan = make(chan string)

type Game struct {
	Arena *Arena
	Score int
}

func NewGame() *Game {
	s := newSnake(RIGHT, [][]int{{1, 1}, {1, 2}, {1, 3}, {1, 4}})
	a := newArena(s, pointsChan, 20, 20)
	return &Game{Arena: a, Score: 0}
}

func (g *Game) end() {
	fmt.Println("Game Over")
	fmt.Printf("Score: %v", g.Score)
}

func (g *Game) handlePointsReceived() {
	for {
		select {
		case p, _ := <-pointsChan:
			g.Score += p
		}
	}
}

func (g *Game) handleKeyPress() {
	for {
		select {
		case k, _ := <-keyboardChan:
			s := g.Arena.Snake
			d := keyToDirection(k)
			s.changeDirection(d)
		}
	}
}

func (g *Game) dificultyMoveInterval() time.Duration {
	return 150 * time.Millisecond
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

	go startKeyboardListener(keyboardChan)
	go g.handleKeyPress()
	defer close(keyboardChan)

mainloop:
	for {
		if err := g.Arena.moveSnake(); err != nil {
			g.end()
			break mainloop
		}

		time.Sleep(g.dificultyMoveInterval())
	}
}
