package simplemath

import "testing"

func TestSqrt1(t *testing.T) {
	v := Sqrt(4)
	if v != 2 {
		t.Error("Sqrt(4) failed. Got %d, expected 2.", v)
	}
}
