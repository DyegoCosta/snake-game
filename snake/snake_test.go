package snake

import "testing"

func NewDoubleSnake(d int) *Snake {
	return newSnake(d, [][]int{
		{1, 0},
		{1, 1},
		{1, 2},
		{1, 3},
		{1, 4},
	})
}

func TestSnakeBodyMove(t *testing.T) {
	snake := NewDoubleSnake(RIGHT)
	snake.move()

	if snake.Body[0][0] != 1 || snake.Body[0][1] != 1 {
		t.Fatalf("Invalid body position %x", snake.Body[0])
	}

	if snake.Body[1][0] != 1 || snake.Body[1][1] != 2 {
		t.Fatalf("Invalid body position %x", snake.Body[1])
	}

	if snake.Body[2][0] != 1 || snake.Body[2][1] != 3 {
		t.Fatalf("Invalid body position %x", snake.Body[2])
	}

	if snake.Body[3][0] != 1 || snake.Body[3][1] != 4 {
		t.Fatalf("Invalid body position %x", snake.Body[3])
	}

	if snake.Body[4][0] != 2 || snake.Body[4][1] != 4 {
		t.Fatalf("Invalid body position %x", snake.Body[4])
	}
}

func TestSnakeHeadMoveRight(t *testing.T) {
	snake := NewDoubleSnake(RIGHT)
	snake.move()

	if snake.head()[0] != 2 || snake.head()[1] != 4 {
		t.Fatalf("Expected head to have moved to position [2 4], got %x", snake.head())
	}
}

func TestSnakeHeadMoveUp(t *testing.T) {
	snake := NewDoubleSnake(UP)
	snake.move()

	if snake.head()[0] != 1 || snake.head()[1] != 5 {
		t.Fatalf("Expected head to have moved to position [1 5], got %x", snake.head())
	}
}

func TestSnakeHeadMoveDown(t *testing.T) {
	snake := NewDoubleSnake(RIGHT)
	snake.move()

	snake.changeDirection(DOWN)
	snake.move()

	if snake.head()[0] != 2 || snake.head()[1] != 3 {
		t.Fatalf("Expected head to have moved to position [2 3], got %x", snake.head())
	}
}

func TestSnakeHeadMoveLeft(t *testing.T) {
	snake := NewDoubleSnake(LEFT)
	snake.move()

	if snake.head()[0] != 0 || snake.head()[1] != 4 {
		t.Fatalf("Expected head to have moved to position [0 4], got %x", snake.head())
	}
}

func TestChangeDirectionToNotOposity(t *testing.T) {
	snake := NewDoubleSnake(DOWN)
	snake.changeDirection(RIGHT)
	if snake.Direction != RIGHT {
		t.Fatal("Expected to change Snake Direction to DOWN")
	}
}

func TestChangeDirectionToOposity(t *testing.T) {
	snake := NewDoubleSnake(RIGHT)
	snake.changeDirection(LEFT)
	if snake.Direction == LEFT {
		t.Fatal("Expected not to have changed Snake Direction to LEFT")
	}
}

func TestChangeDirectionToInvalidDirection(t *testing.T) {
	snake := NewDoubleSnake(RIGHT)
	snake.changeDirection(5)
	if snake.Direction != RIGHT {
		t.Fatal("Expected not to have changed Snake Direction")
	}
}

func TestSnakeDie(t *testing.T) {
	snake := NewDoubleSnake(RIGHT)

	if err := snake.die(); err.Error() != "Died" {
		t.Fatal("Expected Snake die() to return error")
	}
}

func TestSnakeDieWhenMoveOnTopOfItself(t *testing.T) {
	snake := NewDoubleSnake(RIGHT)
	snake.move()

	snake.changeDirection(DOWN)
	snake.move()

	snake.changeDirection(LEFT)

	if err := snake.die(); err.Error() != "Died" {
		t.Fatal("Expected Snake to die when moved on top of itself")
	}
}
