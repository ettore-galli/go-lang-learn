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

func TestSumArray(t *testing.T) {
	t.Run("Test di base", func(t *testing.T) {
		var numbers [ListSize]int = [ListSize]int{1, 2, 3, 4, 5, 6, 7}

		want := 28
		got := SumArray(numbers)
		if got != want {
			t.Errorf("SumList failed; given: %v; want: %d, got; %d", numbers, want, got)
		}
	})
}

func TestSumSlice(t *testing.T) {
	t.Run("Test di base", func(t *testing.T) {
		var numbersSlice []int = []int{1, 2, 3, 4, 5, 6, 7}

		want := 28
		got := SumSlice(numbersSlice)
		if got != want {
			t.Errorf("SumList failed; given: %v; want: %d, got; %d", numbersSlice, want, got)
		}
	})
}
