package kata

import (
	"fmt"
	"sort"
)

func PrimeFactors(n int) string {
	factors := map[int]int{}
	for i := 2; n > 1; i++ {
		if n%i == 0 {
			n /= i
			factors[i]++
			i--
		}
	}

	keys := make([]int, 0)
	for k, _ := range factors {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	result := ""
	for _, k := range keys {
		if factors[k] == 1 {
			result += fmt.Sprintf("(%d)", k)
		} else {
			result += fmt.Sprintf("(%d**%d)", k, factors[k])
		}
	}

	return result
}
