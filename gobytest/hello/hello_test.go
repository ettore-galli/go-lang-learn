package main

import (
	"bytes"
	"testing"
)

func TestHello(t *testing.T) {
	got := Hello()
	want := " ___       _  _       ___    ___       _  _       ___    ___       _  _       ___                      _     _  ___\n|__ \\ ___ | || | ___ |__ \\  |__ \\ ___ | || | ___ |__ \\  |__ \\ ___ | || | ___ |__ \\  __ __ __ ___  _ _ | | __| ||__ \\\n  /_// -_)| || |/ _ \\  /_/    /_// -_)| || |/ _ \\  /_/    /_// -_)| || |/ _ \\  /_/  \\ V  V // _ \\| '_|| |/ _` |  /_/\n (_) \\___||_||_|\\___/ (_)    (_) \\___||_||_|\\___/ (_)    (_) \\___||_||_|\\___/ (_)    \\_/\\_/ \\___/|_|  |_|\\__,_| (_)"

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Ettore")
	want := "Hello, Ettore!"
	got := buffer.String()
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
