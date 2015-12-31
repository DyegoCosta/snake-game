package main

import "math/rand"

type food struct {
	emoji        rune
	points, x, y int
}

func newFood(x, y int) *food {
	return &food{
		points: 10,
		emoji:  newFoodEmoji(),
		x:      x,
		y:      y,
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

func newFoodEmoji() rune {
	if *cutefood {
		return randomFoodEmoji()
	}
	return '@'
}
