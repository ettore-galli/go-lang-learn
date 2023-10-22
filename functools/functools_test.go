package functools

import (
	"testing"
)

func TestFmap(t *testing.T) {

	got := Fmap[int, int]([]int{1, 2, 3}, func(n int) int { return n + 10 })
	want := []int{11, 12, 13}

	if len(got) != len(want) {
		t.Errorf("Size of 'got'.: \n%v \nOf 'want': %v", len(got), len(want))
	}

	for idx, g := range got {
		w := want[idx]
		if g != w {
			t.Errorf("Got.: \n%v \nWant: %v", g, w)
		}
	}

}

func TestFfilter(t *testing.T) {

	got := Ffilter[int]([]int{1, 2, 3, 4, 5, 6, 7, 8}, func(n int) bool { return n%2 == 0 })
	want := []int{2, 4, 6, 8}

	if len(got) != len(want) {
		t.Errorf("Size of 'got'.: \n%v \nOf 'want': %v", len(got), len(want))
	}

	for idx, g := range got {
		w := want[idx]
		if g != w {
			t.Errorf("Got.: \n%v \nWant: %v", g, w)
		}
	}

}

func TestFreduce(t *testing.T) {

	got := Freduce[int]([]int{1, 2, 3, 4, 5, 6, 7, 8}, func(acc int, val int, _ int) int { return acc + val }, 0)
	want := 36

	if got != want {
		t.Errorf("Got: %v; want: %v", got, want)
	}

}
