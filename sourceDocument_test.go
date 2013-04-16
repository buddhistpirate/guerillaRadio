package guerillaRadio

import "testing"

func TestReadFile(t *testing.T) {
	source := SourceDocument{ FileName: "fixtures/oneline.txt"}
	source.ReadFile()

	length := len(source.Lines)
	expected_length := 1

	if expected_length != length {
		t.Errorf("Expected length to be %v, was %v", expected_length, length)
	}
	expected := "first line\n"
	actual := source.Lines[0]

	if expected != string(actual) {
		t.Errorf("expected '%v' , read '%v'", expected, actual)
	}
}
