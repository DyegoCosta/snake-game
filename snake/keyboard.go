package snake

import (
	"fmt"

	"github.com/nsf/termbox-go"
)

func startKeyboardListener(k chan string) {
	termbox.SetInputMode(termbox.InputAlt)
	data := make([]byte, 0, 64)

	for {
		beg := len(data)
		d := data[beg : beg+32]
		switch ev := termbox.PollRawEvent(d); ev.Type {
		case termbox.EventRaw:
			k <- fmt.Sprintf("%X", data[:beg+ev.N])
		case termbox.EventError:
			panic(ev.Err)
		}
	}
}

func keyToDirection(k string) int {
	switch k {
	case "1B4F41":
		return UP
	case "1B4F42":
		return DOWN
	case "1B4F43":
		return RIGHT
	case "1B4F44":
		return LEFT
	default:
		return 0
	}
}
