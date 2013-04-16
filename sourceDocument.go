package guerillaRadio

import "io/ioutil"

type SourceDocument struct {
	FileName string
	Lines []byte
}

func (source *SourceDocument) ReadFile() (err error) {
 source.Lines , err = ioutil.ReadFile(source.FileName)
 return
}


