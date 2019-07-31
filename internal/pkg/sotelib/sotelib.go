package sotelib

import (
	"fmt"
	"os"
	"strings"

	"github.com/azillion/edgcm-converter/climate"
)

// FileReader FileReader
type FileReader interface {
	Read() (*climate.WorldClimate, error)
}

// SotEFile SotEFile
type SotEFile struct {
	File   *os.File
	Reader FileReader
}

// NewSotEFile NewSotEFile
func NewSotEFile(f *os.File) (*SotEFile, error) {
	// TODO: Pass file path instead of file pointer
	fileInfo, err := f.Stat()
	if err != nil {
		return nil, err
	}

	fileName := strings.ToLower(fileInfo.Name())
	if strings.HasSuffix(fileName, ".csv") == false && strings.HasSuffix(fileName, ".dat") == false && strings.HasSuffix(fileName, ".bin") == false {
		return nil, fmt.Errorf("%s is not a supported file type <.csv, .dat, .bin>", fileName)
	}

	soteFile := SotEFile{File: f}

	if strings.HasSuffix(fileName, ".csv") == true {
		reader := NewSotECSVReader(&soteFile)
		soteFile.Reader = FileReader(reader)
	} else {
		reader := NewSotEBINReader(&soteFile)
		soteFile.Reader = FileReader(reader)
	}

	return &soteFile, nil
}
