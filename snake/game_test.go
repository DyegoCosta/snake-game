package snake

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
