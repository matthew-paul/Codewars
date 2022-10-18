package kata

func Beeramid(bonus int, price float64) int {
	var availableCans int = int(float64(bonus) / price)

	level := 1
	for ; level*level <= availableCans; level++ {
		availableCans -= level * level
	}

	return level - 1
}
