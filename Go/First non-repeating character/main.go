package kata

import (
	"fmt"
	"unicode"
)

func FirstNonRepeating(str string) string {
	counts := make([]int, 256)
	for i := 0; i < len(str); i++ {
		ch := unicode.ToLower(rune(str[i]))
		fmt.Println("%i", int(ch))
		counts[int(ch)] += 1
	}
	for i := 0; i < len(str); i++ {
		ch := unicode.ToLower(rune(str[i]))
		fmt.Println("%i", int(ch))
		if counts[int(ch)] == 1 {
			return string(str[i])
		}
	}
	return ""
}
