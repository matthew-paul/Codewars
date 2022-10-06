package kata

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func HighAndLow(in string) string {
	max := math.MinInt32
	min := math.MaxInt32
	numbers := strings.Split(in, " ")
	for _, value := range numbers {
		number, _ := strconv.Atoi(value)
		if number > max {
			max = number
		}
		if number < min {
			min = number
		}
	}

	result := fmt.Sprintf("%v %v", max, min)
	return result
}
