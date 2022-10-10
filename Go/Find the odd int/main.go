package kata

func FindOdd(seq []int) int {
	xor := 0

	for _, i := range seq {
		xor ^= i
	}

	return xor
}
