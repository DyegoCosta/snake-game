package snake

import "testing"

func NewDoubleSnake(d int) Snake {
	return Snake{
		Direction: d,
		Alive:     true,
		Body:      [][]int{{1, 2}, {1, 3}, {1, 4}},
	}
}

func TestSnakeBodyMove(t *testing.T) {
	snake := NewDoubleSnake(RIGHT)
	snake.Move()

	if snake.Body[0][0] != 1 || snake.Body[0][1] != 3 {
		t.Fatalf("Invalid body position %x", snake.Body)
	}

	if snake.Body[1][0] != 1 || snake.Body[1][1] != 4 {
		t.Fatalf("Invalid body position %x", snake.Body)
	}

	if snake.Body[2][0] != 2 || snake.Body[2][1] != 4 {
		t.Fatalf("Invalid body position %x", snake.Body)
	}
}

func TestSnakeHeadMoveRight(t *testing.T) {
	snake := NewDoubleSnake(RIGHT)
	snake.Move()

	if snake.Head()[0] != 2 || snake.Head()[1] != 4 {
		t.Fatalf("Expected head to have moved to position [2 4], got %x", snake.Head())
	}
}

func TestSnakeHeadMoveUp(t *testing.T) {
	snake := NewDoubleSnake(UP)
	snake.Move()

	if snake.Head()[0] != 1 || snake.Head()[1] != 5 {
		t.Fatalf("Expected head to have moved to position [1 5], got %x", snake.Head())
	}
}

func TestSnakeHeadMoveDown(t *testing.T) {
	snake := NewDoubleSnake(DOWN)
	snake.Move()

	if snake.Head()[0] != 1 || snake.Head()[1] != 3 {
		t.Fatalf("Expected head to have moved to position [2 3], got %x", snake.Head())
	}
}

func TestSnakeHeadMoveLeft(t *testing.T) {
	snake := NewDoubleSnake(LEFT)
	snake.Move()

	if snake.Head()[0] != 0 || snake.Head()[1] != 4 {
		t.Fatalf("Expected head to have moved to position [0 4], got %x", snake.Head())
	}
}

func TestChangeDirectionToNotOposity(t *testing.T) {
	snake := NewDoubleSnake(DOWN)
	snake.ChangeDirection(RIGHT)
	if snake.Direction != RIGHT {
		t.Fatal("Expected to change Snake Direction to DOWN")
	}
}

func TestChangeDirectionToOposity(t *testing.T) {
	snake := NewDoubleSnake(RIGHT)
	snake.ChangeDirection(LEFT)
	if snake.Direction == LEFT {
		t.Fatal("Expected not to have changed Snake Direction to LEFT")
	}
}

func TestSnakeDie(t *testing.T) {
	snake := NewDoubleSnake(RIGHT)
	snake.Die()
	if snake.Alive != false {
		t.Fatal("Expected Snake not to be alive")
	}
}
