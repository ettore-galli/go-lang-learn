package iter

import (
	"fmt"
	"testing"
)

func TestRepeatChar(t *testing.T) {
	want := "xxxxx"
	got := RepeatChar("x", 5)

	if got != want {
		t.Errorf("Want [%s], got [%s]", want, got)
	}
}

func BenchmarkRepeatChar(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RepeatChar("x", 5)
	}

}

func ExampleRepeatChar() {
	fmt.Print(RepeatChar("w", 3))
	// Output: www
}
