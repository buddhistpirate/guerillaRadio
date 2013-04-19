package guerillaradio

import "bufio"
import "os"

type SourceDocument struct {
	FileName string
	Lines    []string
}

func (source *SourceDocument) ReadFile() (err error) {
	file, err := os.Open(source.FileName)
	reader := bufio.NewReader(file)
	for err == nil {
		line, err := reader.ReadString('\n')
		source.Lines = append(source.Lines, line)
		if err == nil {
			break
		}
	}

	return
}
