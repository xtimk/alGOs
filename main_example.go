package main

// import from github to re-use code
// this file is an example of how to use algos package inside your own project

import (
	"fmt"

	"github.com/xtimk/alGOs/algos"
)

func main() {
	fmt.Println("** Naive Pattern Matching **")
	fmt.Println()
	text := "accgaagcttacacggcgacgacgtcg"
	pattern := "acg"
	matches := algos.NaivePatternMatching(text, pattern)
	fmt.Println("Text: ", text)
	fmt.Println("Pattern: ", pattern)
	fmt.Println("Pretty Print of Matches:")
	algos.PrettyPrintMatches(text, pattern, matches)
	fmt.Println("-----------------")
	fmt.Println()

	fmt.Println("** LPS: Longest Prefix that is also a Suffix ** ")
	pattern = "abcdabca"
	lps := algos.Lps(pattern)
	fmt.Println("Text: ", pattern)
	fmt.Println("LPS: ", lps)
	fmt.Println("-----------------")
	fmt.Println()

	fmt.Println("** KMP: Knuth Morris Pratt ** ")
	text = "accgaagcttacacggcgacgacgtcg"
	pattern = "acg"
	matches = algos.Kmp(text, pattern)
	fmt.Println("Text: ", text)
	fmt.Println("Pattern: ", pattern)
	// fmt.Println("Matches: ", matches)
	fmt.Println("Pretty Print of Matches:")
	algos.PrettyPrintMatches(text, pattern, matches)
	fmt.Println("-----------------")
	fmt.Println()

	fmt.Println("** ZFunc: Z array computation ** ")
	text = "accgaagcttacacggcgacgacgtcg"
	matches = algos.Zfunc(text)
	fmt.Println("Text: ", text)
	fmt.Println("ZArray: ", matches)
	fmt.Println("Pretty Print of ZArray:")
	algos.PrettyPrintZArray(text, matches)
	fmt.Println("-----------------")
	fmt.Println()
}

// Pretty print the matches
