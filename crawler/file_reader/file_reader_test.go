package file_reader

import (
	"reflect"
	"testing"
)

func TestReadFile(t *testing.T) {
	got, err := ReadFile("test_data/example.txt")
	if err != nil {
		t.Error(err)
	}
	want := []string{
		"www.google.com",
		"www.nasa.gov",
		"www.google.com",
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("\nGot.: \n%v \nWant: \n%v", got, want)
	}
}
