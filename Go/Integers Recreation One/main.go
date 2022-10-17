package kata

import (
	"math"
)

func ListSquared(m, n int) [][]int {
	result := [][]int{}
	for i := m; i <= n; i++ {
		sum := 0
		for j := 1; j <= n; j++ {
			if i%j == 0 {
				sum += j * j
			}
		}
		if math.Mod(math.Sqrt(float64(sum)), 1.) == 0 {
			slice := []int{i, sum}
			result = append(result, slice)
		}
	}
	return result
}
