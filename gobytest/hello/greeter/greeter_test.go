package greeter

import "testing"

func TestMakeHelloGreeting(t *testing.T) {
	got := MakeHelloGreeting()
	want := "Hello, world!"

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
