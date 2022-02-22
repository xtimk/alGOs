package algos

// Returns the LPS array: For each substring the longest prefix that is also suffix.
// solved using a recursive approach
func Lps(text string) []int {
	lps := make([]int, len(text))
	lps[0] = 0
	i := 0
	j := 1

	for j < len(text) {
		if text[i] == text[j] {
			lps[j] = i + 1
			i++
			j++
		} else {
			if i == 0 {
				lps[j] = 0
				j++
			} else {
				i = lps[i-1]
			}

		}
	}

	return lps
}
