package main

import "testing"

var pointsDouble = make(chan int)

func newDoubleArenaWithFoodFinder(h, w int, f func(*Arena, Coord) bool) *Arena {
	a := newDoubleArena(h, w)
	a.hasFood = f
	return a
}

func newDoubleArena(h, w int) *Arena {
	s := newSnake(RIGHT, []Coord{
		Coord{X: 1, Y: 0},
		Coord{X: 1, Y: 1},
		Coord{X: 1, Y: 2},
		Coord{X: 1, Y: 3},
		Coord{X: 1, Y: 4},
	})

	return newArena(s, pointsDouble, h, w)
}

func TestArenaHaveFoodPlaced(t *testing.T) {
	if a := newDoubleArena(20, 20); a.Food == nil {
		t.Fatal("Arena expected to have food placed")
	}
}

func TestMoveSnakeOutOfArenaHeightLimit(t *testing.T) {
	a := newDoubleArena(4, 10)
	a.Snake.changeDirection(UP)

	if err := a.moveSnake(); err == nil || err.Error() != "Died" {
		t.Fatal("Expected Snake to die when moving outside the Arena height limits")
	}
}

func TestMoveSnakeOutOfArenaWidthLimit(t *testing.T) {
	a := newDoubleArena(10, 1)
	a.Snake.changeDirection(LEFT)

	if err := a.moveSnake(); err == nil || err.Error() != "Died" {
		t.Fatal("Expected Snake to die when moving outside the Arena height limits")
	}
}

func TestPlaceNewFoodWhenEatFood(t *testing.T) {
	a := newDoubleArenaWithFoodFinder(10, 10, func(*Arena, Coord) bool {
		return true
	})

	f := a.Food

	a.moveSnake()

	if a.Food.X == f.X && a.Food.Y == f.Y {
		t.Fatal("Expected new food to have been placed on Arena")
	}
}

func TestIncreaseSnakeLengthWhenEatFood(t *testing.T) {
	a := newDoubleArenaWithFoodFinder(10, 10, func(*Arena, Coord) bool {
		return true
	})

	l := a.Snake.Length

	a.moveSnake()

	if a.Snake.Length != l+1 {
		t.Fatal("Expected Snake to have grown")
	}
}

func TestAddPointsWhenEatFood(t *testing.T) {
	a := newDoubleArenaWithFoodFinder(10, 10, func(*Arena, Coord) bool {
		return true
	})

	if p, ok := <-pointsDouble; ok && p != a.Food.Points {
		t.Fatalf("Value %d was expected but got %d", a.Food.Points, p)
	}

	a.moveSnake()
}

func TestDoesNotAddPointsWhenFoodNotFound(t *testing.T) {
	a := newDoubleArenaWithFoodFinder(10, 10, func(*Arena, Coord) bool {
		return false
	})

	select {
	case p, _ := <-pointsChan:
		t.Fatalf("No point was expected to be received but received %d", p)
	default:
		close(pointsChan)
	}

	a.moveSnake()
}

func TestDoesNotPlaceNewFoodWhenFoodNotFound(t *testing.T) {
	a := newDoubleArenaWithFoodFinder(10, 10, func(*Arena, Coord) bool {
		return false
	})

	f := a.Food

	a.moveSnake()

	if a.Food.X != f.X || a.Food.Y != f.Y {
		t.Fatal("Food in Arena expected not to have changed")
	}
}

func TestDoesNotIncreaseSnakeLengthWhenFoodNotFound(t *testing.T) {
	a := newDoubleArenaWithFoodFinder(10, 10, func(*Arena, Coord) bool {
		return false
	})

	l := a.Snake.Length

	a.moveSnake()

	if a.Snake.Length != l {
		t.Fatal("Expected Snake not to have grown")
	}
}

func TestHasFood(t *testing.T) {
	a := newDoubleArena(20, 20)

	if !hasFood(a, Coord{X: a.Food.X, Y: a.Food.Y}) {
		t.Fatal("Food expected to be found")
	}
}

func TestHasNotFood(t *testing.T) {
	a := newDoubleArena(20, 20)

	if hasFood(a, Coord{X: a.Food.X - 1, Y: a.Food.Y}) {
		t.Fatal("No food expected to be found")
	}
}
