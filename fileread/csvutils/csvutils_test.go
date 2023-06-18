package csvutils

import (
	"bytes"
	"reflect"
	"testing"
)

const TestString = "Maps differing: \nExpected: %v \nActual..: %v"

func TestHappyPathFirstOk(t *testing.T) {
	got, err := ReadCsvMap("./data/test_happy_path.csv", true)
	if err != nil {
		t.Error(err)

	}

	want := OutMap{
		"1": "AAA",
		"2": "BBB",
		"3": "CCC",
		"4": "DDD",
	}

	if reflect.DeepEqual(got, want) == false {
		t.Errorf(TestString, want, got)
	}

}

func TestHappyPathLastOk(t *testing.T) {
	got, err := ReadCsvMap("./data/test_happy_path.csv", false)
	if err != nil {
		t.Error(err)

	}

	want := OutMap{
		"1": "AAA",
		"2": "BBB",
		"3": "C33",
		"4": "DDD",
	}

	if reflect.DeepEqual(got, want) == false {
		t.Errorf(TestString, want, got)
	}

}

func TestPerformReader(t *testing.T) {
	buffer := &bytes.Buffer{}

	buffer.Write([]byte("1,AAA\n2,BBB\n3,CCC\n4,DDD\n3,C33"))

	got, err := PerformReadCsvMap(buffer, false)
	if err != nil {
		t.Error(err)

	}
	want := OutMap{
		"1": "AAA",
		"2": "BBB",
		"3": "C33",
		"4": "DDD",
	}

	if reflect.DeepEqual(got, want) == false {
		t.Errorf(TestString, want, got)
	}

}
