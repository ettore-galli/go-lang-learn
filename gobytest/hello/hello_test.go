package main

import (
	"bytes"
	"reflect"
	"testing"
	"time"
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

type SpyDelayer struct {
	calls int
}

func (delayer *SpyDelayer) Sleep(d time.Duration) {
	delayer.calls++
}

func TestCountdown(t *testing.T) {
	buffer := bytes.Buffer{}
	spyDelayer := SpyDelayer{}
	Countdown(&buffer, &spyDelayer)
	want := "3\n2\n1\nGo!"
	got := buffer.String()
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
	wantCalls := 3
	spyCalls := spyDelayer.calls
	if spyDelayer.calls != wantCalls {
		t.Errorf("Inconsistent number of calls: got %q, want %q", spyCalls, wantCalls)
	}
}

type SpyInspector struct {
	operations []string
}

func (delayer *SpyInspector) Sleep(d time.Duration) {
	delayer.operations = append(delayer.operations, "sleep")
}

func (delayer *SpyInspector) Write(p []byte) (n int, err error) {
	delayer.operations = append(delayer.operations, "print")
	return 0, nil
}

func TestCountdownLogic(t *testing.T) {
	spy := SpyInspector{}
	Countdown(&spy, &spy)
	wantCalls := []string{"print", "sleep", "print", "sleep", "print", "sleep", "print"}
	spyCalls := spy.operations
	if !reflect.DeepEqual(spyCalls, wantCalls) {
		t.Errorf("got %q, want %q", spyCalls, wantCalls)
	}
}
