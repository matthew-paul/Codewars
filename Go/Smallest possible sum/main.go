package kata

import "math"

func GCD(x int, y int) int {
	for y != 0 {
		x, y = y, x%y
	}

	return x
}

func Solution(ar []int) int {
	X := make([]int, len(ar))
	copy(X, ar)
	min := X[0]
	for i := 1; i < len(X); i++ {
		min = int(math.Min(float64(GCD(min, X[i])), float64(min)))
		X[i] = min
	}

	return min * len(ar)
}
