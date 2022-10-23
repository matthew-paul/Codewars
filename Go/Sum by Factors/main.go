package kata

import (
	"fmt"
	"sort"
)

func getFactors(n int) []int {
	result := []int{}
	if n < 0 {
		n *= -1
	}
	for i := 2; i <= n; i++ {
		if n%i == 0 {
			result = append(result, i)
			for n%i == 0 {
				n /= i
			}
		}
	}
	return result
}

func SumOfDivided(lst []int) string {
	primeFactors := make(map[int]int)
	for _, n := range lst {
		fmt.Printf("n: %v\n", n)
		primes := getFactors(n)
		fmt.Printf("prime: ")
		for _, prime := range primes {
			fmt.Printf("%v ", prime)
			primeFactors[prime] += n
		}
		fmt.Println("")
	}

	keys := make([]int, 0)
	for k, _ := range primeFactors {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	result := ""
	for _, key := range keys {
		result += fmt.Sprintf("(%v %v)", key, primeFactors[key])
	}

	return result
}
