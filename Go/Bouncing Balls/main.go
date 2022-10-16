package kata

func BouncingBall(h, bounce, window float64) int {
	if h < 0 || bounce <= 0 || bounce >= 1 || window >= h {
		return -1
	}
	count := -1
	for ; h > window; h *= bounce {
		count += 2
	}
	return count
}
