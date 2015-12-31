package main

import (
	"testing"
)

func TestFoodDefaultPoints(t *testing.T) {
	f := newFood(10, 10, '?')

	if f.points != 10 {
		t.Fatalf("Expected Food default points to be 10 but got %v", f.points)
	}
}

func TestCuteFoodEmoji(t *testing.T) {
	f := newCuteFood(10, 10)

	if string(f.emoji) == "" {
		t.Fatal("Food emoji not expected to be blank")
	}
}

func TestBoringCuteFoodEmoji(t *testing.T) {
	f := newBoringFood(10, 10)

	if string(f.emoji) != "@" {
		t.Fatal("Boring food emoji differnet from expected")
	}
}
