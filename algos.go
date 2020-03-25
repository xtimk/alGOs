package main

import (
	"fmt"
	"github.com/xtimk/alGOs/naivepmatching"
	)

func main() {
	fmt.Println("Starting..\n")
	matches := naivepmatching.Naive_pmatching("accgaagcttacacggcgacgacgtcg", "acg")
	fmt.Println(matches)
}