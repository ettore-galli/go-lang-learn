package calc

import (
	"fmt"
	"testing"
)

func TestSum(t *testing.T) {
	want := 5
	got := sum(2, 3)
	if got != want {
		t.Errorf("Sum failed; want: %d, got; %d", want, got)
	}
}

func Example_calcSum() {
	a := 3
	b := 7
	s := sum(a, b)
	fmt.Printf("%d + %d = %d\n", a, b, s)
	// Output: 3 + 7 = 10
}
