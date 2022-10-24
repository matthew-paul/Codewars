package kata

import (
	"strconv"
)

func LastDigit(n1, n2 string) int {
	max := func(a, b int) int {
		if a > b {
			return a
		} else {
			return b
		}
	}
	pow := func(n, e int) int {
		res := 1
		for i := 0; i < e; i++ {
			res *= n
		}
		return res
	}
	a, _ := strconv.Atoi(n1[len(n1)-1:])
	b, _ := strconv.Atoi(n2[max(0, len(n2)-2):])
	return pow(a, (b%4)+4) % 10

}
