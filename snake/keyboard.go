package snake

import "github.com/nsf/termbox-go"

func listenToKeyboard(moves chan Direction, end chan bool) {
	termbox.SetInputMode(termbox.InputEsc)

	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			switch ev.Key {
			case termbox.KeyArrowLeft:
				moves <- LEFT
			case termbox.KeyArrowDown:
				moves <- DOWN
			case termbox.KeyArrowRight:
				moves <- RIGHT
			case termbox.KeyArrowUp:
				moves <- UP
			case termbox.KeyEsc:
				end <- true
			}
		case termbox.EventError:
			panic(ev.Err)
		}
	}
}
