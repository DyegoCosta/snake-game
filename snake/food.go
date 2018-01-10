package snake

import (
	"os"
	"strings"
)

type food struct {
	emoji        rune
	points, x, y int
}

func newFood(x, y int) *food {
	return &food{
		points: 10,
		emoji:  getFoodEmoji(),
		x:      x,
		y:      y,
	}
}

func getFoodEmoji() rune {
	if hasUnicodeSupport() {
		return randomFoodEmoji()
	} else {
		return '@'
	}
}

func randomFoodEmoji() rune {
	f := []rune{
		'🍒',
		'🍍',
		'🍑',
		'🍇',
		'🍏',
		'🍌',
		'🍫',
		'🍭',
		'🍕',
		'🍩',
		'🍗',
		'🍖',
		'🍬',
		'🍤',
		'🍪',
	}

	return f[randomInt(len(f))]
}

func hasUnicodeSupport() bool {
	return strings.Contains(os.Getenv("LANG"), "UTF-8")
}
