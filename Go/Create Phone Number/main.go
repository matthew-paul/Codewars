package kata

import "fmt"

func CreatePhoneNumber(numbers [10]uint) string {
	numbersInterface := make([]interface{}, 10)
	for i, val := range numbers {
		numbersInterface[i] = val
	}
	return fmt.Sprintf("(%v%v%v) %v%v%v-%v%v%v%v", numbersInterface...)
}
