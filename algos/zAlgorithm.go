package algos

import (
	"errors"
	"strings"
)

func Z(text string, pattern string) ([]int, error) {
	var matches []int
	if strings.Contains(text, "$") || strings.Contains(pattern, "$") {
		return matches, errors.New("Z: text/pattern cant contain the $ char.\n")
	}

	concatString := pattern + "$" + text
	zarray := Zfunc(concatString)

	for i, z := range zarray {
		if z == len(pattern) {
			matches = append(matches, (i - len(pattern) - 1))
		}
	}

	return matches, nil
}
