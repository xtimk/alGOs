package algos

import (
	"fmt"
	"testing"
)

func Test_kmp_no_match_1(t *testing.T) {
	got := Kmp("MyText", "MyP")
	var expected []int
	if !eqSlices(got, expected) {
		errorString := fmt.Sprint("Failed test ", getCurrentFuncName(), ": ", "got: ", got, " - Expected: ", expected)
		t.Errorf(errorString)
	}
}

func Test_kmp_no_match_pattern_longer_than_text(t *testing.T) {
	got := Kmp("MyText", "MyLongPattern")
	var expected []int
	if !eqSlices(got, expected) {
		errorString := fmt.Sprint("Failed test ", getCurrentFuncName(), ": ", "got: ", got, " - Expected: ", expected)
		t.Errorf(errorString)
	}
}

func Test_kmp_match_1(t *testing.T) {
	got := Kmp("accgaagcttacacggcgacgacgtcg", "acg")
	expected := []int{12, 18, 21}
	if !eqSlices(got, expected) {
		errorString := fmt.Sprint("Failed test ", getCurrentFuncName(), ": ", "got: ", got, " - Expected: ", expected)
		t.Errorf(errorString)
	}
}
