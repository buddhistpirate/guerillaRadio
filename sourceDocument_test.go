package guerillaRadio

import "testing"

func TestReadFile(t *testing.T) {
	expected := "first line"
	source := SourceDocument{ FileName: "fixtures/oneline.txt"}
	source.ReadFile()
	actual := source.Lines[0]
	if expected != string(actual) {
		t.Errorf("expected '%v' , read '%v'", expected, actual)
	}
}
