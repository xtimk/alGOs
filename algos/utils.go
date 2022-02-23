package algos

import (
	"fmt"
	"runtime"
	"strings"
)

// utility functions:
// func whos name is starting with a capital letter are exported and usable by another user
// the other func are used for internal purposes and should not be available for the user

// PrettyPrintMatches: a simple pretty printer for the matches of a pattern in a text
// Func will output something like this
// accgaagcttacacggcgacgacgtcg
//             acg
//                   acg
//                      acg
func PrettyPrintMatches(text string, pattern string, matches []int) {
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

func getCurrentFuncName() string {
	pc := make([]uintptr, 10) // at least 1 entry needed
	runtime.Callers(2, pc)
	f := runtime.FuncForPC(pc[0])
	// file, line := f.FileLine(pc[0])
	// fmt.Printf("%s:%d %s\n", file, line, f.Name())
	return f.Name()
}

func eqSlices(a, b []int) bool {
	// If one is nil, the other must also be nil.
	if (a == nil) != (b == nil) {
		return false
	}

	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}
