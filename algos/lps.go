package algos

// Returns the LPS array: For each substring the longest prefix that is also suffix.
// solved using a recursive approach
// Complexity: O(|text|)
func Lps(text string) []int {
	lps := make([]int, len(text))
	lps[0] = 0
	i := 0
	j := 1

	// 'i' will be the current length of the LPS of the substring text[0,j]

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
				// tricky part: here we consider the LPS of the current LPS that we are checking
				// note that this LPS can also be 0, and in this case we end up in simply comparing text[0] with text[j]
			}
		}
	}
	return lps
}
