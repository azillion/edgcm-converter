package sotelib

type SotECSVReader struct {}

func (csv *SotECSVFile) Read() *WorldClimate {
	return nil
}

func NewSotECSVFile(f *os.File) *SotECSVFile {
	soteFile := SotECSVFile{
		File: f
	}
	return &soteFile
}