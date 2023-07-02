package functools

import (
	"testing"
)

func TestFunctionMap(t *testing.T) {

	got := fmap[int, int]([]int{1, 2, 3}, func(n int) int { return n + 10 })
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
