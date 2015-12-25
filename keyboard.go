package main

import "github.com/nsf/termbox-go"

type KeyboardEventType int

const (
	MOVE KeyboardEventType = 1 + iota
	RETRY
	END
)

type KeyboardEvent struct {
	Type KeyboardEventType
	Key  termbox.Key
}

func keyToDirection(k termbox.Key) Direction {
	switch k {
	case termbox.KeyArrowLeft:
		return LEFT
	case termbox.KeyArrowDown:
		return DOWN
	case termbox.KeyArrowRight:
		return RIGHT
	case termbox.KeyArrowUp:
		return UP
	default:
		return 0
	}
}

func listenToKeyboard(evChan chan KeyboardEvent) {
	termbox.SetInputMode(termbox.InputEsc)

	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			switch ev.Key {
			case termbox.KeyArrowLeft:
				evChan <- KeyboardEvent{Type: MOVE, Key: ev.Key}
			case termbox.KeyArrowDown:
				evChan <- KeyboardEvent{Type: MOVE, Key: ev.Key}
			case termbox.KeyArrowRight:
				evChan <- KeyboardEvent{Type: MOVE, Key: ev.Key}
			case termbox.KeyArrowUp:
				evChan <- KeyboardEvent{Type: MOVE, Key: ev.Key}
			case termbox.KeyEsc:
				evChan <- KeyboardEvent{Type: END, Key: ev.Key}
			default:
				if ev.Ch == 'r' {
					evChan <- KeyboardEvent{Type: RETRY, Key: ev.Key}
				}
			}
		case termbox.EventError:
			panic(ev.Err)
		}
	}
}
