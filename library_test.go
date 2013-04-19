package guerillaradio

import "testing"

func TestAddFileToLibrary(t *testing.T) {
	library := Library{}
	err := library.AddSourceDocument("fixtures/oneline.txt")

	if err != nil {
		t.Errorf("Received an error adding source document %v", err)
	}

	number_of_documents := library.Size()
	expected_number := 1

	if expected_number != number_of_documents {
		t.Errorf("Expected number of documens to be %v, was %v", expected_number, number_of_documents)
	}

}

func TestAddFilesToLibrary(t *testing.T) {
	library := Library{}
	filenames := []string{"fixtures/oneline.txt", "fixtures/twolines.txt"}
	num_files, err := library.AddSourceDocuments(filenames)

	if err != nil {
		t.Errorf("Received an error adding documents to Library: %v", err)
	}

	expected_number_of_documents := 2

	if expected_number_of_documents != num_files {
		t.Errorf("Expected %v files to be added but library contains: %v", expected_number_of_documents, num_files)
	}

}

func TestCanRetrieveRequestedLines(t *testing.T) {
	library := Library{}
	filenames := []string{"fixtures/oneline.txt", "fixtures/twolines.txt", "fixtures/sixlines.txt"}
	num_files, err := library.AddSourceDocuments(filenames)

	if err != nil || num_files != 3 {
		t.Errorf("Received an error adding documents to Library: %v", err)
	}
	lines_to_retrieve := 12
	lines := library.RetrieveLines(lines_to_retrieve)

	//fmt.Println(lines)

	if len(lines) != lines_to_retrieve {
		t.Errorf("Received %v lines instead of the request %v", len(lines), lines_to_retrieve)
	}
}

func TestCanReadAllFilesInADirectoryTree(t *testing.T) {
	library := Library{}

	directory := "fixtures"
	err := library.AddDirectory(directory)

	if err != nil {
		t.Errorf("Received an error adding documents to Library: %v", err)
	}

	expected_number_of_documents := 4

	if expected_number_of_documents != library.Size() {
		t.Errorf("Expected %v files to be added but library contains: %v", expected_number_of_documents, library.Size())
	}

}
