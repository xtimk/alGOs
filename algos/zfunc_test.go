package algos

import (
	"fmt"
	"testing"
)

func Test_Zfunc_abcabc(t *testing.T) {
	text := "abcabc"
	got := Zfunc(text)
	expected := []int{0, 0, 0, 3, 0, 0}
	if !eqSlices(got, expected) {
		errorString := fmt.Sprint("Failed test ", getCurrentFuncName(), ": ", "got: ", got, " - Expected: ", expected)
		t.Errorf(errorString)
	}
}
func Test_aux_findZ(t *testing.T) {
	text := "abcabc"
	got := aux_findZ(text, 3, 0)
	expected := 3
	if got != expected {
		errorString := fmt.Sprint("Failed test ", getCurrentFuncName(), ": ", "got: ", got, " - Expected: ", expected)
		t.Errorf(errorString)
	}
}

func Test_aux_findZ2(t *testing.T) {
	text := "abcabcdf"
	got := aux_findZ(text, 3, 0)
	expected := 3
	if got != expected {
		errorString := fmt.Sprint("Failed test ", getCurrentFuncName(), ": ", "got: ", got, " - Expected: ", expected)
		t.Errorf(errorString)
	}
}
