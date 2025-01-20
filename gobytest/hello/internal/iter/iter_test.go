package iter

import (
	"testing"
)

func TestRepeatChar(t *testing.T) {
	want := "xxxxx"
	got := RepeatChar("x", 5)

	if got != want {
		t.Errorf("Want [%s], got [%s]", want, got)
	}
}
