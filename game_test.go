package main

import (
	"testing"
	"time"
)

func TestDefaultGameScore(t *testing.T) {
	g := NewGame()

	if g.score != 0 {
		t.Fatalf("Initial Game Score expected to be 0 but it was %d", g.score)
	}
}

func TestGameMoveInterval(t *testing.T) {
	e := time.Duration(85) * time.Millisecond
	g := NewGame()
	g.score = 150

	if d := g.moveInterval(); d != e {
		t.Fatalf("Expected move interval to be %d but got %d", e, d)
	}
}

func TestAddPoints(t *testing.T) {
	g := NewGame()
	s := g.score
	g.addPoints(10)

	if s != 0 || g.score != 10 {
		t.Fatal("Expected 10 points to have been added to Game Score")
	}
}

func TestRetryGoBackToGameInitialState(t *testing.T) {
	g := NewGame()
	initScore := g.score
	initSnake := g.arena.snake

	g.arena.snake.changeDirection(UP)
	g.arena.moveSnake()
	g.addPoints(10)
	g.end()

	g.retry()

	if g.score != initScore {
		t.Fatal("Expected Score to have been reset")
	}

	for i, c := range g.arena.snake.body {
		if initSnake.body[i].x == c.x && initSnake.body[i].y == c.y {
			t.Fatal("Expected Snake body to have been reset")
		}
	}

	if g.arena.snake.direction == initSnake.direction {
		t.Fatal("Expected Snake direction to have been reset")
	}
}

func TestGetMoreChance(t *testing.T) {
	g := NewGame()
	initScore := g.score
	initChance := g.chance
	initSnakeLength := g.arena.snake.length

	g.arena.snake.changeDirection(UP)

	for i := 1; i <= initChance; i++ {
		g.arena.moveSnake()
		g.arena.snake.length++
		g.addPoints(10)
		g.end()

		if i < initChance {
			if g.isOver != false {
				t.Fatalf("Expected player will get more chance : got %v", g.isOver)
			}

			if g.score == initScore {
				t.Fatalf("Expected score will not reset when get more chance : got %v", g.score)
			}

			if initSnakeLength == g.arena.snake.length {
				t.Fatalf("Expected snake length not reset when get more chance : got %v", g.arena.snake.length)
			}
		} else {
			if g.isOver == false {
				t.Fatalf("Expected player won't get more chance : got %v", g.isOver)
			}
		}
	}
}
