package snake

import "testing"

func newDoubleArena(h, w int) *Arena {
	s := newSnake(RIGHT, [][]int{
		{1, 0},
		{1, 1},
		{1, 2},
		{1, 3},
		{1, 4},
	})

	return newArena(s, h, w)
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
	t.Fatal()
}

func TestIncreaseSnakeLengthWhenEatFood(t *testing.T) {
	t.Fatal()
}

func TestAddPointsWhenEatFood(t *testing.T) {
	t.Fatal()
}
