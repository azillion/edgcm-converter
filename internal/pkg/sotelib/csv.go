package sotelib

import (
	"encoding/csv"
	"os"
	"strconv"

	"github.com/azillion/edgcm-converter/climate"
	log "github.com/sirupsen/logrus"
)

// SotECSVReader SotECSVReader
type SotECSVReader struct {
	file *os.File
}

// NewSotECSVReader NewSotECSVReader
func NewSotECSVReader(f *SotEFile) *SotECSVReader {
	csvFile := SotECSVReader{file: f.File}
	return &csvFile
}

func (f *SotECSVReader) Read() (*climate.WorldClimate, error) {
	wc := new(climate.WorldClimate)

	lines, err := csv.NewReader(f.file).ReadAll()
	if err != nil {
		return nil, err
	}

	wc.Length = int32(len(lines[1:]))
	wc.GridSize = lengthToGridSize(wc.Length)
	log.Debugf("File Cell Length: %+v\n", wc.Length)
	log.Debugf("File GridSize: %+v\n", wc.GridSize)

	// Loop through lines & turn into object
	for i, line := range lines[1:] {
		cell := climate.Cell{
			CellID:  int32(i),
			Climate: parseCSVLine(line, int32(i+1)),
		}
		wc.Cells = append(wc.Cells, &cell)
	}
	log.Debug("fully parsed")

	return wc, nil
}

func parseCSVLine(line []string, lineNumber int32) climate.Climate {
	c := climate.Climate{
		Latitude:           parseCSVFloat(line[0], lineNumber, 0),
		WaterPercentage:    parseCSVFloat(line[1], lineNumber, 1),
		Elevation:          parseCSVFloat(line[2], lineNumber, 2),
		Temperature:        parseCSVFloat(line[3], lineNumber, 3),
		HumidityPercentage: parseCSVFloat(line[4], lineNumber, 4),
		Rainfall:           parseCSVFloat(line[5], lineNumber, 5),
	}
	return c
}

func parseCSVFloat(line string, lineNumber, columnNumber int32) float32 {
	f, err := strconv.ParseFloat(line, 32)
	if err != nil {
		log.Debugf("error parsing CSV line: [%d][%d]> %s", lineNumber, columnNumber, line)
		return 0
	}
	return float32(f)
}
