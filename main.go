package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"

	"github.com/azillion/edgcm-converter/climate"

	"github.com/azillion/edgcm-converter/internal/pkg/sotelib"

	log "github.com/sirupsen/logrus"
)

var (
	debug    bool
	filePath string
)

func init() {
	flag.BoolVar(&debug, "d", false, "debug mode")
	flag.StringVar(&filePath, "f", "./input-data/sote.csv", "input data file path")
}

func main() {
	flag.Parse()

	if debug == true {
		log.SetLevel(log.DebugLevel)
		log.Debug("Debug Mode Enabled")
	}

	worldClimate := parseClimateFile(filePath)

	exportClimateData(worldClimate, "./sote.out")
}

func parseClimateFile(path string) *climate.WorldClimate {
	f, err := os.Open(path)
	if err != nil {
		log.Fatalf("Failed to open the input file: %s", path)
	}
	defer f.Close()

	soteFile, err := sotelib.NewSotEFile(f)
	if err != nil {
		log.Fatal(err)
	}
	log.Debugf("Created SotEFile: %+v", soteFile)

	worldClimate, err := soteFile.Reader.Read()
	if err != nil {
		log.Fatal(err)
	}
	log.Infof("Number of Parsed Cells: %d", len(worldClimate.Cells))
	return worldClimate
}

func exportClimateData(wc *climate.WorldClimate, path string) {
	f2, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer f2.Close()

	datawriter := bufio.NewWriter(f2)

	s1 := fmt.Sprintf("World Climate\nLength: %+v\n", wc.Length)
	_, _ = datawriter.WriteString(s1)
	s2 := fmt.Sprintf("GridSize: %+v\n", wc.GridSize)
	_, _ = datawriter.WriteString(s2)
	for _, cell := range wc.Cells {
		s := fmt.Sprintf("Cell[%d]:%+v\n", cell.CellID, cell.Climate)
		_, _ = datawriter.WriteString(s)
	}

	datawriter.Flush()
	fstat, err := f2.Stat()
	if err != nil {
		log.Fatal(err)
	}
	log.Debugf("Exported Climate Data: %s", fstat.Name())
}
