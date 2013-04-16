package guerillaRadio

import "bufio"
import "os"
import "fmt"

type SourceDocument struct {
	FileName string
	Lines []string
}

func (source *SourceDocument) ReadFile() (err error) {
	file, err := os.Open(source.FileName)
	reader := bufio.NewReader(file)
	for err == nil {
		line, err := reader.ReadString('\n')
		source.Lines = append(source.Lines, line)
		if err == nil { break }
	}
	
	if err != nil {
		fmt.Sprintf("ERROR: ReadFile encountered: %v", err)
		return
	}

	return
}


