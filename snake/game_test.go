package snake

import "testing"

func TestDefaultGameScore(t *testing.T) {
	g := NewGame()

	if g.Score != 0 {
		t.Fatalf("Initial Game Score expected to be 0 but it was %d", g.Score)
	}
}
