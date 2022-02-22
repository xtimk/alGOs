package main

// import from github to re-use code
import (
	"fmt"

	"github.com/xtimk/alGOs/algos"
)

func main() {
	fmt.Println("Naive Pattern Matching")
	text := "accgaagcttacacggcgacgacgtcg"
	pattern := "acg"
	matches := algos.NaivePatternMatching(text, pattern)
	algos.PrintMatches(text, pattern, matches)
	fmt.Println("-----------------")

	fmt.Println("\nLPS: Longest Prefix that is also a Suffix")
	pattern = "abcdabca"
	lps := algos.Lps(pattern)
	fmt.Println(lps)
}

// Pretty print the matches
