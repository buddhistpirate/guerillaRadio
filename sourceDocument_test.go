package guerillaradio

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

func TestReadFileWithTwoLines(t *testing.T) {
	source := SourceDocument{FileName: "fixtures/twolines.txt"}
	err := source.ReadFile()

	if err != nil {
		t.Errorf("ReadFile had an unexpcted err %v", err)
	}

	length := len(source.Lines)
	expected_length := 2

	if expected_length != length {
		t.Errorf("Expected length to be %v, was %v", expected_length, length)
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

func TestRetrievesOnlyOneLine(t *testing.T) {

	source := SourceDocument{FileName: "fixtures/oneline.txt"}
	err := source.ReadFile()

	if err != nil {
		t.Errorf("Test could not read file unexpectedly: %v", err)
	}

	number_of_lines_requested := 2
	lines, lines_returned := source.RetrieveLines(number_of_lines_requested)

	if len(lines) != lines_returned {
		t.Errorf("Number of lines returned does not match reported number of lines. %v != %v", len(lines), lines_returned)
	}

	if lines_returned != 1 {
		t.Errorf("%v Lines returned and we expected only a single one since the file only has one")
	}

}

func TestRetrieveMultipleLines(t *testing.T) {

	source := SourceDocument{FileName: "fixtures/sixlines.txt"}
	err := source.ReadFile()

	if err != nil {
		t.Errorf("Test could not read file unexpectedly: %v", err)
	}

	number_of_lines_requested := 5
	lines, lines_returned := source.RetrieveLines(number_of_lines_requested)

	if len(lines) != lines_returned {
		t.Errorf("Number of lines returned does not match reported number of lines. %v != %v", len(lines), lines_returned)
	}

	expected_lines := 1
	if lines_returned != expected_lines {
		t.Errorf("%v Lines returned and we expected %v with the rand seed value", lines_returned, expected_lines)
	}
}
