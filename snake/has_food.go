package snake

func hasFood(a *Arena, p []int) bool {
	return p[0] == a.Food.X && p[1] == a.Food.Y
}
