package guerillaradio

import "testing"

func TestCanJsonifyLibraryRequest(t *testing.T) {

	library := Library{}
	err := library.AddSourceDocument("fixtures/oneline.txt")

	if err != nil {
		t.Errorf("Error while loading library: %v", err)
	}

	json, err := ServiceRequest(&library, 1)

	if err != nil {
		t.Errorf("Error while converting json: %v", err)
	}

	expected_json := "[\"first line\\n\"]"

	if expected_json != string(json) {
		t.Errorf("Converted Document: %v did not match %v", string(json), expected_json)
	}
}

func TestCanConvertFromRequestToInt(t *testing.T) {
	num_lines, err := convertRequestToInt("5\n")
	if err != nil {
		t.Errorf("Could not convert string to int: %v",err)
	}

	if 5 != num_lines {
		t.Errorf("Could not convert string to 5")
	}
}
