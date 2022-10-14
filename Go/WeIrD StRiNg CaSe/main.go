package kata

import "strings"

func toWeirdCase(str string) string {
	result := ""

	wordIndex := 0
	for i := 0; i < len(str); i++ {
		if wordIndex%2 == 0 {
			result = result + strings.ToUpper(string(str[i]))
		} else {
			result = result + strings.ToLower(string(str[i]))
		}
		wordIndex++
		if str[i] == ' ' {
			wordIndex = 0
		}
	}

	return result

}
