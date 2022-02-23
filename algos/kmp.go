package algos

func Kmp(text string, pattern string) []int {
	textIndex := 0
	patternIndex := 0
	lps := Lps(pattern)

	var matches []int

	for textIndex < len(text) {
		possibleMatch := textIndex
		for patternIndex < len(pattern) && text[textIndex] == pattern[patternIndex] {
			patternIndex++
			textIndex++
		}

		if patternIndex == len(pattern) {
			matches = append(matches, possibleMatch)
		}

		if patternIndex != 0 {
			patternIndex = lps[patternIndex-1]
		} else {
			textIndex++
		}
	}
	return matches
}
