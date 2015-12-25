package main

import (
	"testing"
	"time"
)

func TestDefaultGameScore(t *testing.T) {
	g := NewGame()

	if g.Score != 0 {
		t.Fatalf("Initial Game Score expected to be 0 but it was %d", g.Score)
	}
}

func TestGameMoveInterval(t *testing.T) {
	e := time.Duration(85) * time.Millisecond
	g := NewGame()
	g.Score = 150

	if d := g.moveInterval(); d != e {
		t.Fatalf("Expected move interval to be %d but got %d", e, d)
	}
}

func TestAddPoints(t *testing.T) {
	g := NewGame()
	s := g.Score
	g.addPoints(10)

	if s != 0 || g.Score != 10 {
		t.Fatal("Expected 10 points to have been added to Game Score")
	}
}

func TestRetryGoBackToGameInitialState(t *testing.T) {
	g := NewGame()
	initScore := g.Score
	initSnake := g.Arena.Snake

	g.Arena.Snake.changeDirection(UP)
	g.Arena.moveSnake()
	g.addPoints(10)
	g.end()

	g.retry()

	if g.Score != initScore {
		t.Fatal("Expected Score to have been reset")
	}

	for i, c := range g.Arena.Snake.Body {
		if initSnake.Body[i].X == c.X && initSnake.Body[i].Y == c.Y {
			t.Fatal("Expected Snake body to have been reset")
		}
	}

	if g.Arena.Snake.Direction == initSnake.Direction {
		t.Fatal("Expected Snake direction to have been reset")
	}
}
