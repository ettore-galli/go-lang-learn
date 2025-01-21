package calc

import (
	"fmt"
	"testing"
)

func TestSum(t *testing.T) {
	want := 5
	got := Sum(2, 3)
	if got != want {
		t.Errorf("Sum failed; want: %d, got; %d", want, got)
	}
}

func Example_calcSum() {
	a := 3
	b := 7
	s := Sum(a, b)
	fmt.Printf("%d + %d = %d\n", a, b, s)
	// Output: 3 + 7 = 10
}

func TestSumList(t *testing.T) {
	var numbers [ListSize]int = [ListSize]int{1, 2, 3, 4, 5, 6, 7}

	want := 28
	got := SumList(numbers)
	if got != want {
		t.Errorf("SumList failed; given: %v; want: %d, got; %d", numbers, want, got)
	}
}
