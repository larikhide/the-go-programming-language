package ch1

import (
	"testing"
)

func TestDup1(t *testing.T) {
	input :=
		`line1
line2
dup
dup
line3`

	want := "dup\t2"

	got := dup1(input)

	if want != got {
		t.Errorf("Expected: %s\nGot: %s", want, got)
	}

}
