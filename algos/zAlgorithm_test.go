package algos

import (
	"fmt"
	"testing"
)

func Test_Z_1(t *testing.T) {
	got, err := Z("MyText", "MyP")
	var expected []int
	if !eqSlices(got, expected) || err != nil {
		errorString := fmt.Sprint("Failed test ", getCurrentFuncName(), ": ", "got: ", got, " - Expected: ", expected)
		t.Errorf(errorString)
	}
}

func Test_Z_no_match_pattern_longer_than_text(t *testing.T) {
	got, err := Z("MyText", "MyLongPattern")
	var expected []int
	if !eqSlices(got, expected) || err != nil {
		errorString := fmt.Sprint("Failed test ", getCurrentFuncName(), ": ", "got: ", got, " - Expected: ", expected)
		t.Errorf(errorString)
	}
}

func Test_Z_match_1(t *testing.T) {
	got, err := Z("accgaagcttacacggcgacgacgtcg", "acg")
	expected := []int{12, 18, 21}
	if !eqSlices(got, expected) || err != nil {
		errorString := fmt.Sprint("Failed test ", getCurrentFuncName(), ": ", "got: ", got, " - Expected: ", expected)
		t.Errorf(errorString)
	}
}

func Test_Z_text_containing_dollar_char(t *testing.T) {
	got, err := Z("accgaagcttaca$cggcgacgacgtcg", "acg")
	expected := []int{12, 18, 21}
	if err == nil {
		errorString := fmt.Sprint("Failed test ", getCurrentFuncName(), ": ", "got: ", got, " - Expected: ", expected)
		t.Errorf(errorString)
	}
}
