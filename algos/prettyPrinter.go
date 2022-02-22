package algos

import (
	"fmt"
	"strings"
)

// Func will output something like this
// accgaagcttacacggcgacgacgtcg
//             acg
//                   acg
//                      acg
func PrintMatches(text string, pattern string, matches []int) {
	fmt.Println(text)
	for _, match := range matches {
		prettyStringMatch := createEmptyString(len(text))
		prettyStringMatch = substituteStringAtIndex(prettyStringMatch, pattern, match)
		fmt.Println(prettyStringMatch)

	}
}

func createEmptyString(length int) string {
	result := strings.Repeat(" ", length)
	return result
}

func substituteStringAtIndex(text string, sub string, index int) string {
	return text[:index] + sub + text[index+1-len(sub):]
}
