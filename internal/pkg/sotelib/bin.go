package sotelib

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"os"

	"github.com/azillion/edgcm-converter/climate"
	log "github.com/sirupsen/logrus"
)

// SotEBINReader SotEBINReader
type SotEBINReader struct {
	file *os.File
}

// NewSotEBINReader NewSotEBINReader
func NewSotEBINReader(f *SotEFile) *SotEBINReader {
	binFile := SotEBINReader{file: f.File}
	return &binFile
}

// ClimateReader read a SotE Bin file
func (f *SotEBINReader) Read() (*climate.WorldClimate, error) {
	wc := new(climate.WorldClimate)
	reader := bufio.NewReader(f.file)

	// num of rows in the file
	var numOfCells int32
	numOfCellsBytes, err := readNextBytes(reader, 4)
	if err != nil {
		return nil, err
	}
	buffer := bytes.NewBuffer(numOfCellsBytes)

	// convert buffer to int32
	err = binary.Read(buffer, binary.LittleEndian, &numOfCells)
	if err != nil {
		return nil, err
	}
	wc.Length = numOfCells

	// grid size is the square root of length
	wc.GridSize = lengthToGridSize(wc.Length)
	log.Debugf("File Cell Length: %+v\n", wc.Length)
	log.Debugf("File GridSize: %+v\n", wc.GridSize)

	for i := 0; i < int(wc.Length); i++ {
		climateInCell := new(climate.Climate)

		data, err := readNextBytes(reader, 24)
		if err != nil {
			return nil, err
		}
		buffer2 := bytes.NewBuffer(data)
		err = binary.Read(buffer2, binary.LittleEndian, climateInCell)
		if err != nil {
			return nil, err
		}

		cell := climate.Cell{
			CellID:  int32(i),
			Climate: *climateInCell,
		}
		// log.Debugf("cell:%+v\n", cell) // debug line will print every cell
		wc.Cells = append(wc.Cells, &cell)
	}

	return wc, nil
}

func readNextBytes(b *bufio.Reader, number int) ([]byte, error) {
	bytes := make([]byte, number)

	_, err := b.Read(bytes)
	if err != nil {
		return bytes, err
	}

	return bytes, nil
}
