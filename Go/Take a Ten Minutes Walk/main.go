package kata

import "fmt"

func IsValidWalk(walk []rune) bool {
	nsDist := 0
	ewDist := 0
	if len(walk) != 10 {
		return false
	}

	for i := 0; i < len(walk); i++ {
		fmt.Printf("%q\n", walk[i])
		switch walk[i] {
		case 'n':
			nsDist += 1
		case 's':
			nsDist -= 1
		case 'e':
			ewDist += 1
		case 'w':
			ewDist -= 1
		}
	}

	return (nsDist == 0 && ewDist == 0)
}
