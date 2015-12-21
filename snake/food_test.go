package snake

import (
	"testing"
)

func TestFoodDefaultPoints(t *testing.T) {
	f := NewFood(10, 10)

	if f.Points != 10 {
		t.Fatalf("Expected Food default points to be 10 but got %v", f.Points)
	}
}

func TestFoodEmoji(t *testing.T) {
	f := NewFood(10, 10)

	if string(f.Emoji) == "" {
		t.Fatal("Food emoji not expected to be blank")
	}
}
