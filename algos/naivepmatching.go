package algos

func NaivePatternMatching(text string, pattern string) []int {
	textIndex := 0
	patternIndex := 0

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

		textIndex = possibleMatch + 1
		patternIndex = 0
	}
	return matches
}

func AnotherNaiveApproach(Text string, Pattern string) []int {

	var matches []int
	Text_pointer := 0
	// Pattern_pointer := 0

	Text_len := len(Text)
	Pattern_len := len(Pattern)

	Ipotetic_match := 0

	for (Text_len - Ipotetic_match) >= Pattern_len {
		match_found := true
		Text_pointer = Ipotetic_match
		for x := range Pattern {
			if Text[Text_pointer+x] != Pattern[x] {
				match_found = false
				break
			}
		}
		if match_found {
			// fmt.Println("Found a Match")
			matches = append(matches, Ipotetic_match)
		} // else {
		// fmt.Println("NOT a Match")
		// }
		Ipotetic_match = Ipotetic_match + 1
	}

	// fmt.Println(string(Text[Text_pointer]))
	// for t := range Text {
	// 	for p := range Pattern {

	// 	}
	// }

	return matches
}
