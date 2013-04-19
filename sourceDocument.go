package guerillaradio

import "bufio"
import "os"
import "math"
import "math/rand"

type SourceDocument struct {
	FileName string
	Lines    []string
}

func (source *SourceDocument) ReadFile() (err error) {
	file, err := os.Open(source.FileName)
	reader := bufio.NewReader(file)
	for err == nil {
		line, err := reader.ReadString('\n')
		if line != "" {
			source.Lines = append(source.Lines, line)
		}
		if err != nil {
			break
		}
	}

	return
}

func (source *SourceDocument) RetrieveLines(number_of_lines int) (lines []string, lines_returned int) {
	index := rand.Intn(len(source.Lines))
	end := math.Min(float64(len(source.Lines)), float64(index+number_of_lines))
	lines = source.Lines[index:int(end)]
	lines_returned = len(lines)
	return
}
