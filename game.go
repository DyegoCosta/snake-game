package main

import (
	"time"

	"github.com/nsf/termbox-go"
)

var (
	pointsChan         = make(chan int)
	keyboardEventsChan = make(chan keyboardEvent)
)

type game struct {
	arena  *arena
	score  int
	isOver bool
	chance int
}

func initialSnake(snakeSize int) *snake {
	cr := []coord{}
	for i := 0; i < snakeSize; i++ {
		cr = append(cr, coord{x: 1, y: i + 1})
	}
	return newSnake(RIGHT, cr)
}

func initialScore() int {
	return 0
}

func initialArena() *arena {
	return newArena(initialSnake(4), pointsChan, 20, 50)
}

func initialChance() int {
	return 3
}

func (g *game) end() {
	if g.chance > 0 {
		g.chance--
		if g.chance == 0 {
			g.isOver = true
		} else {
			g.newChance()
		}
	} else {
		g.isOver = true
	}
}

func (g *game) moveInterval() time.Duration {
	ms := 100 - (g.score / 10)
	return time.Duration(ms) * time.Millisecond
}

func (g *game) newChance() {
	g.arena = newArena(initialSnake(g.arena.snake.length), pointsChan, 20, 50)
}

func (g *game) retry() {
	g.arena = initialArena()
	g.score = initialScore()
	g.chance = initialChance()
	g.isOver = false
}

func (g *game) addPoints(p int) {
	g.score += p
}

func NewGame() *game {
	return &game{arena: initialArena(), score: initialScore(), chance: initialChance()}
}

func (g *game) Start() {
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

			time.Sleep(g.moveInterval())
		}
	}
}
