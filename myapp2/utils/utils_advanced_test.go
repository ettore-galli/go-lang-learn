package utils

import (
	"testing"
)

func TestFormatSomma100(t *testing.T) {

	got := FormatSomma100(3, 4)
	want := "Somma100(3, 4) => 700"
	if got != want {
		t.Errorf("%v != %v", got, want)
	}

}
