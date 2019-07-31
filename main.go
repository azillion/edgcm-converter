package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"

	"github.com/azillion/edgcm-converter/internal/pkg/sotelib"

	log "github.com/sirupsen/logrus"
)

var (
	debug    bool
	filePath string
)

func init() {
	flag.BoolVar(&debug, "d", false, "debug mode")
	flag.StringVar(&filePath, "f", "./input-data/sote.dat", "input data file path")
}

func main() {
	flag.Parse()

	if debug == true {
		log.SetLevel(log.DebugLevel)
		log.Debug("Debug Mode Enabled")
	}

	f, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("Failed to open the input file: %s", filePath)
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

	f2, err := os.OpenFile("./input-data/sote.out", os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer f2.Close()

	datawriter := bufio.NewWriter(f2)

	s1 := fmt.Sprintf("World Climate\nLength: %+v\n", worldClimate.Length)
	_, _ = datawriter.WriteString(s1)
	s2 := fmt.Sprintf("GridSize: %+v\n", worldClimate.GridSize)
	_, _ = datawriter.WriteString(s2)
	for _, cell := range worldClimate.Cells {
		s := fmt.Sprintf("Cell[%d]:%+v\n", cell.CellID, cell.Climate)
		_, _ = datawriter.WriteString(s)
	}

	datawriter.Flush()
}
