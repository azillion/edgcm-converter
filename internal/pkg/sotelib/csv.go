package sotelib

import (
	"fmt"
	"os"
	// "encoding/csv"

	"github.com/azillion/edgcm-converter/climate"
)

// SotECSVReader SotECSVReader
type SotECSVReader struct {
	file *os.File
}

func (f *SotECSVReader) Read() (*climate.WorldClimate, error) {
	return nil, fmt.Errorf("CSV is not implemented")
}

// NewSotECSVReader NewSotECSVReader
func NewSotECSVReader(f *SotEFile) *SotECSVReader {
	csvFile := SotECSVReader{file: f.File}
	return &csvFile
}
