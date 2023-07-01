package response_writer

import (
	"os"
	"path/filepath"
	"testing"
)

func TestResponseWriter(t *testing.T) {

	dir, _ := os.MkdirTemp("", "example")

	defer os.RemoveAll(dir)

	tempFileName := "TEST_OUTPUT.txt"

	rw := ResponseWriter{fileNameBuilder: func(c string) string { return tempFileName }}
	rw.WriteResponse(dir, "<Hello><World></World></Hello>")

	tempFileFqn := filepath.Join(dir, tempFileName)
	got_bytes, _ := os.ReadFile(tempFileFqn)

	got := string(got_bytes)
	want := "<Hello><World></World></Hello>"

	if got != want {
		t.Errorf("Got.: \n%v \nWant: %v", got, want)
	}

}
