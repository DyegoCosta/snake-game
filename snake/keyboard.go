package snake

import "github.com/nsf/termbox-go"

func (g *Game) startKeyboardArrowsListener() {
	termbox.SetInputMode(termbox.InputAlt)
	s := g.Arena.Snake

	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			switch ev.Key {
			case termbox.KeyArrowLeft:
				s.changeDirection(LEFT)
			case termbox.KeyArrowDown:
				s.changeDirection(DOWN)
			case termbox.KeyArrowRight:
				s.changeDirection(RIGHT)
			case termbox.KeyArrowUp:
				s.changeDirection(UP)
			}
		case termbox.EventError:
			panic(ev.Err)
		}
	}
}
