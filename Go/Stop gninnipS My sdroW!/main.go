package kata

import "strings"

func reverse(str string) string {
	result := []rune(str)
	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}
	return string(result)
}

func SpinWords(str string) string {
	result := ""
	words := strings.Split(str, " ")
	for _, word := range words {
		if len(word) >= 5 {
			result += reverse(word) + " "
		} else {
			result += word + " "
		}
	}

	return strings.Trim(result, " ")
}
