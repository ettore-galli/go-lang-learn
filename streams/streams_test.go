package streams

import (
	"fmt"
	"testing"
)

func TestProduce(t *testing.T) {

	ch := produceStrings()

	var got []string

	for item := range *ch {
		got = append(got, item)
	}

	want := []string{"aaa", "bbb", "ccc"}

	s := StreamSource[string]{}

	fmt.Println("STREAMS")
	fmt.Println(s)

	for i, s := range got {
		if s != want[i] {
			t.Errorf("At %v: Got: %v; want: %v", i, s, want[i])
		}
	}

}

func TestStream(t *testing.T) {

	got := 3
	want := 3

	s := StreamSource[string]{}

	fmt.Println("STREAMS")
	fmt.Println(s)

	if got != want {
		t.Errorf("Got: %v; want: %v", got, want)
	}

}
