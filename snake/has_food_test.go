package snake

import "testing"

func TestHasFood(t *testing.T) {
	a := newDoubleArena(20, 20)

	if !hasFood(a, []int{a.Food.X, a.Food.Y}) {
		t.Fatal("Food expected to be found")
	}
}

func TestHasNotFood(t *testing.T) {
	a := newDoubleArena(20, 20)

	if hasFood(a, []int{a.Food.X - 1, a.Food.Y}) {
		t.Fatal("No food expected to be found")
	}
}
