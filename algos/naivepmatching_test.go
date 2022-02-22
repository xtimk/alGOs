package algos

import (
	"testing"
)

func eqSlices(a, b []int) bool {
	// If one is nil, the other must also be nil.
	if (a == nil) != (b == nil) {
		return false
	}

	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}

func TestHi(t *testing.T) {
	got := Hi()
	if got != "Hi" {
		t.Errorf("Hi doesn't returned correct valulue")
	}
}

func Test_no_match_1(t *testing.T) {
	got := NaivePatternMatching("MyText", "MyP")
	var expected []int
	if !eqSlices(got, expected) {
		t.Errorf("Failed Test_no_match_1 func")
	}
}

func Test_no_match_pattern_longer_than_text(t *testing.T) {
	got := NaivePatternMatching("MyText", "MyLongPattern")
	var expected []int
	if !eqSlices(got, expected) {
		t.Errorf("Failed Test_no_match_pattern_longer_than_text func")
	}
}

func Test_a_match_1(t *testing.T) {
	got := NaivePatternMatching("accgaagcttacacggcgacgacgtcg", "acg")
	expected := []int{12, 18, 21}
	if !eqSlices(got, expected) {
		t.Errorf("Failed Test_a_match_1 func")
	}
}
