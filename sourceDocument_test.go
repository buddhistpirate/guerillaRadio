package guerillaRadio

import "testing"

func TestReadFile(t *testing.T) {
	source := SourceDocument{FileName: "fixtures/oneline.txt"}
	err := source.ReadFile()

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
	
	if err != nil {
		t.Errorf("ReadFile had an unexpcted err %v", err)
	}
}

func TestNoFileToRead(t *testing.T) {
	source := SourceDocument{FileName: "/noway/this/works"}
	err := source.ReadFile()

	if err == nil {
		t.Errorf("Error not returned for fake File")
	}

	expected_size := 0
	if expected_size != len(source.Lines) {
		t.Errorf("Expected to Slice to be empty")
	}
}
