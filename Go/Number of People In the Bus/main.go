package kata

func Number(busStops [][2]int) int {
	total := 0
	for _, stops := range busStops {
		total += stops[0]
		total -= stops[1]
	}

	return total
}
