package calc

import (
	"fmt"
	"reflect"
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
			t.Errorf("given: %v; want: %d, got; %d", numbers, want, got)
		}
	})
}

func TestSumSlice(t *testing.T) {
	t.Run("Test di base", func(t *testing.T) {
		var numbersSlice []int = []int{1, 2, 3, 4, 5, 6, 7}

		want := 28
		got := SumSlice(numbersSlice)
		if got != want {
			t.Errorf("given: %v; want: %d, got; %d", numbersSlice, want, got)
		}
	})
}

func TestSumSlices(t *testing.T) {
	t.Run("Test di base", func(t *testing.T) {
		var slice1 []int = []int{1, 2, 3}
		var slice2 []int = []int{10, 20, 30}
		var slice3 []int = []int{100, 200, 300}

		want := []int{6, 60, 600}
		got := SumSlices(slice1, slice2, slice3)

		if !reflect.DeepEqual(got, want) {
			t.Errorf("want: %d, got; %d", want, got)
		}
	})
}

func TestSumTails(t *testing.T) {
	checkResult := func(t testing.TB, got []int, want []int) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("want: %d, got; %d", want, got)
		}

	}
	t.Run("Test di base", func(t *testing.T) {
		var slice1 []int = []int{1, 2, 3}
		var slice2 []int = []int{10, 20, 30}
		var slice3 []int = []int{100, 200, 300}

		want := []int{5, 50, 500}
		got := SumTails(slice1, slice2, slice3)

		checkResult(t, got, want)

	})

	t.Run("Test vuoti", func(t *testing.T) {
		var slice1 []int = []int{}
		var slice2 []int = []int{}

		want := []int{0, 0}
		got := SumTails(slice1, slice2)

		checkResult(t, got, want)

	})
}

func TestPerimeter(t *testing.T) {
	t.Run("Test di base", func(t *testing.T) {
		want := 50.0
		got := Perimeter(Rectangle{width: 10.0, height: 15.0})
		if got != want {
			t.Errorf("want: %f, got; %f", want, got)
		}
	})
}

func TestAreaStandalone(t *testing.T) {
	t.Run("Rettangolo", func(t *testing.T) {
		want := 150.0
		got := Area(Rectangle{width: 10.0, height: 15.0})
		if got != want {
			t.Errorf("want: %f, got; %f", want, got)
		}
	})

}

func TestAreaMethods(t *testing.T) {
	checkArea := func(t testing.TB, shape Shape, want float64) {
		t.Helper()

		got := shape.Area()
		if got != want {
			t.Errorf("want: %f, got; %f", want, got)
		}
	}

	t.Run("Rettangolo", func(t *testing.T) {
		rectangle := Rectangle{width: 10.0, height: 15.0}
		want := 150.0
		checkArea(t, rectangle, want)

	})

	t.Run("Cerchio", func(t *testing.T) {
		circle := Circle{radius: 10.0}
		want := 628.0
		checkArea(t, circle, want)
	})

}

func TestAreaMulti(t *testing.T) {
	testCases := []struct {
		shape Shape
		want  float64
	}{
		{shape: Rectangle{2, 3}, want: 6.0},
		{shape: Circle{5}, want: 157.0},
		{shape: Triangle{base: 5, height: 3}, want: 7.5},
	}

	for _, testCase := range testCases {
		t.Run(fmt.Sprintf("%T", testCase.shape), func(t *testing.T) {
			got := testCase.shape.Area()
			if got != testCase.want {
				t.Errorf("case: %+v : want: %f, got; %f", testCase.shape, testCase.want, got)
			}
		})
	}

}
