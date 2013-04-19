package guerillaradio

import "math/rand"
import "os"
import "path/filepath"
import "fmt"

type Library struct {
	documents []SourceDocument
}

func (library *Library) Size() int {
	return len(library.documents)
}

func (library *Library) AddSourceDocument(filename string) (err error) {
	source := SourceDocument{FileName: filename}
	err = source.ReadFile()

	if err != nil {
		return
	}

	library.documents = append(library.documents, source)
	return
}

func (library *Library) AddSourceDocuments(filenames []string) (num_files int, err error) {
	num_files = 0
	for _, filename := range filenames {
		err = library.AddSourceDocument(filename)
		if err != nil {
			return
		}
		num_files++
	}
	return
}

func (library *Library) RetrieveLines(number_of_lines int) (lines []string) {
	index := rand.Intn(library.Size())
	source := library.documents[index]
	lines, lines_returned := source.RetrieveLines(number_of_lines)
	if lines_returned < number_of_lines {
		lines = append(lines, library.RetrieveLines(number_of_lines-lines_returned)...)
	}
	return
}

func (library *Library) AddDirectory(name string) (err error) {
	walkTextTree := func(rootpath string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			err := library.AddSourceDocument(rootpath)
			if err != nil {
				fmt.Printf("Unable to add %v due to: %v", rootpath, err)
			}
		}
		return err
	}
	err = filepath.Walk(name, walkTextTree)
	return
}
