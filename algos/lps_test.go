package algos

import (
	"fmt"
	"testing"
)

func Test_LPS_abcdabca(t *testing.T) {
	pattern := "abcdabca"
	got := Lps(pattern)
	expected := []int{0, 0, 0, 0, 1, 2, 3, 1}
	if !eqSlices(got, expected) {
		errorString := fmt.Sprint("Failed test ", getCurrentFuncName(), ": ", "got: ", got, " - Expected: ", expected)
		t.Errorf(errorString)
	}
}

func Test_LPS_aca(t *testing.T) {
	pattern := "aca"
	got := Lps(pattern)
	expected := []int{0, 0, 1}
	if !eqSlices(got, expected) {
		t.Errorf("Failed Test_no_match_1 func")
	}
}

func Test_LPS_aaaaaa(t *testing.T) {
	pattern := "aaaaaaa"
	got := Lps(pattern)
	expected := []int{0, 1, 2, 3, 4, 5, 6}
	if !eqSlices(got, expected) {
		t.Errorf("Failed Test_no_match_1 func")
	}
}
