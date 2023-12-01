package pi

import (
	"math"
)

func NumbersInPi(pi string, numbers []string) int {

	m := make(map[string]bool)

	for _, v := range numbers {
		m[v] = true
	}

	minSpaces := minSpace(pi, m, 0)

	if minSpaces == math.MaxInt32 {
		return -1
	}
	return minSpaces
}

func minSpace(pi string, m map[string]bool, i int) int {
	if i == len(pi) {
		return -1
	}
	minSpaces := math.MaxInt32
	for j := i; j < len(pi); j++ {
		piPrefix := pi[i : j+1]
		if _, ok := m[piPrefix]; ok {

			minSpacePrefix := minSpace(pi, m, j+1)
			if minSpacePrefix+1 < minSpaces {
				minSpaces = minSpacePrefix + 1
			}
		}
	}
	return minSpaces
}
