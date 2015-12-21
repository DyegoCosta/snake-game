package snake

import "github.com/nsf/termbox-go"

func listenToKeyboard(k chan int) {
	termbox.SetInputMode(termbox.InputAlt)

	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			switch ev.Key {
			case termbox.KeyArrowLeft:
				k <- LEFT
			case termbox.KeyArrowDown:
				k <- DOWN
			case termbox.KeyArrowRight:
				k <- RIGHT
			case termbox.KeyArrowUp:
				k <- UP
			}
		case termbox.EventError:
			panic(ev.Err)
		}
	}
}
