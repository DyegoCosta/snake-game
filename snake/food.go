package snake

import "math/rand"

type Food struct {
	Emoji        rune
	Points, X, Y int
}

func NewFood(x, y int) *Food {
	return &Food{
		Points: 10,
		Emoji:  randomFoodEmoji(),
		X:      x,
		Y:      y,
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

	return f[rand.Intn(len(f))]
}
