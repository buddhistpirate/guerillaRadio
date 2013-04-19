package guerillaradio

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
	for _,filename := range filenames {
		err = library.AddSourceDocument(filename)
		if err != nil {
			return
		}
		num_files++
	}
	return
}
