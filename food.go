package main

import "math/rand"

const defpoints = 10

type food struct {
	emoji        rune
	points, x, y int
}

func newFood(x int, y int, emoji rune) *food {
	return &food{
		points: defpoints,
		emoji:  emoji,
		x:      x,
		y:      y,
	}
}

func newCuteFood(x, y int) *food {
	return newFood(x, y, randomFoodEmoji())
}

func newBoringFood(x, y int) *food {
	return newFood(x, y, '@')
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
