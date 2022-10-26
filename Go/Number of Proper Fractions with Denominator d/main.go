package kata

func ProperFractions(n int) int {

	if n == 1 {
		return 0
	}

	// using eulers totient function
	result := n
	distinct_primes := make(map[int]int)
	for i := 2; i*i <= n; {
		if n%i == 0 {
			distinct_primes[i]++
			n /= i
		} else {
			i++
		}
	}
	if n > 1 {
		distinct_primes[n]++
	}
	if len(distinct_primes) == 0 {
		result = n - 1
	} else {
		for k, _ := range distinct_primes {
			result = (result * (k - 1)) / k
		}
	}

	return result
}
