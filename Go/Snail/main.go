package kata

import "fmt"

func Snail(snailMap [][]int) []int {
	result := []int{}
	if len(snailMap) == 0 {
		return result
	}
	xAmount := len(snailMap[0])
	yAmount := len(snailMap) - 1

	fmt.Printf("%v %v\n", xAmount, yAmount)

	x, y := 0, 0

	for xAmount > 0 || yAmount > 0 {
		for i := xAmount; i > 0; i-- {
			result = append(result, snailMap[y][x])
			x++
		}
		xAmount--
		y++
		x--
		for i := yAmount; i > 0; i-- {
			result = append(result, snailMap[y][x])
			y++
		}
		yAmount--
		x--
		y--
		for i := xAmount; i > 0; i-- {
			result = append(result, snailMap[y][x])
			x--
		}
		xAmount--
		y--
		x++
		for i := yAmount; i > 0; i-- {
			result = append(result, snailMap[y][x])
			y--
		}
		yAmount--
		x++
		y++
	}

	return result
}
