package utils

import "testing"

func TestSomma100(t *testing.T) {
	cases := []struct {
		a, b, want int
	}{
		{3, 2, 500},
		{1, 1, 200},
	}
	for _, c := range cases {
		got := Somma100(c.a, c.b)
		if got != c.want {
			t.Errorf("Somma100(%v, %v) == %v, want %v", c.a, c.b, got, c.want)
		}
	}
}

func BenchmarkSomma100(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Somma100(i, i)
	}
}

func ExampleStampaSomma100() {
	StampaSomma100(3, 4)
	// Output:
	// Somma100(3, 4) => 700
}
