package snake

import (
	"fmt"

	"github.com/mattn/go-runewidth"
	"github.com/nsf/termbox-go"
)

const (
	defaultColor = termbox.ColorDefault
	bgColor      = termbox.ColorDefault
	snakeColor   = termbox.ColorGreen
	foodColor    = termbox.ColorRed
)

func (g *Game) render() {
	termbox.Clear(defaultColor, defaultColor)

	var (
		w, h   = termbox.Size()
		midY   = h / 2
		midX   = (w - g.Arena.Width) / 2
		left   = midX - 1
		top    = midY - (g.Arena.Height / 2)
		bottom = midY + (g.Arena.Height / 2)
	)

	renderArena(g.Arena, top, bottom, midX)
	renderSnake(left, bottom, g.Arena.Snake)
	renderFood(left, bottom, g.Arena.Food)
	renderScore(midX, bottom, g.Score)

	termbox.Flush()
}

func renderSnake(left, bottom int, s *Snake) {
	for _, b := range s.Body {
		termbox.SetCell(left+b[0], bottom-b[1], '▇', snakeColor, bgColor)
	}
}

func renderFood(left, bottom int, f *Food) {
	termbox.SetCell(left+f.X, bottom-f.Y, '⌘', foodColor, bgColor)
}

func renderArena(a *Arena, top, bottom, midX int) {
	for i := top; i < bottom; i++ {
		termbox.SetCell(midX-1, i, '│', defaultColor, bgColor)
		termbox.SetCell(midX+a.Width, i, '│', defaultColor, bgColor)
	}

	termbox.SetCell(midX-1, top, '┌', defaultColor, bgColor)
	termbox.SetCell(midX-1, bottom, '└', defaultColor, bgColor)
	termbox.SetCell(midX+a.Width, top, '┐', defaultColor, bgColor)
	termbox.SetCell(midX+a.Width, bottom, '┘', defaultColor, bgColor)

	fill(midX, top, a.Width, 1, termbox.Cell{Ch: '─'})
	fill(midX, bottom, a.Width, 1, termbox.Cell{Ch: '─'})
}

func renderScore(midX, bottom, s int) {
	score := fmt.Sprintf("Score: %v", s)
	tbprint(midX, bottom+1, defaultColor, defaultColor, score)
}

func fill(x, y, w, h int, cell termbox.Cell) {
	for ly := 0; ly < h; ly++ {
		for lx := 0; lx < w; lx++ {
			termbox.SetCell(x+lx, y+ly, cell.Ch, cell.Fg, cell.Bg)
		}
	}
}

func tbprint(x, y int, fg, bg termbox.Attribute, msg string) {
	for _, c := range msg {
		termbox.SetCell(x, y, c, fg, bg)
		x += runewidth.RuneWidth(c)
	}
}
