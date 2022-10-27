package kata

import "math/big"

func reverseNumber(num int) int {
	res := 0
	for num > 0 {
		remainder := num % 10
		res = (res * 10) + remainder
		num /= 10
	}
	return res
}

func SqCubRevPrime(n int) int {
	lastMatch := -1
	for i := 0; i < n; i++ {
		for j := lastMatch + 1; ; j++ {
			x := reverseNumber(j * j)
			y := reverseNumber(j * j * j)
			if big.NewInt(int64(x)).ProbablyPrime(20) && big.NewInt(int64(y)).ProbablyPrime(20) {
				lastMatch = j
				break
			}
		}
	}
	return lastMatch
}
