package sotelib

import (
    "bufio"
    "math"
    "fmt"
    "bytes"
    "encoding/binary"
    "os"
    "strings"
	
	"github.com/azillion/edgcm-converter/internal/pkg/sotelib"
)

type FileReader interface {
    Read() *WorldClimate
}

type SotEFile struct {
    File *os.File
    Reader FileReader
}

func NewFileReader(f *os.File) FileReader {
    fileInfo, err := f.Stat()
    check(err)

    fileName := strings.ToLower(fileInfo.Name())
    if (strings.HasSuffix(fileName, ".csv") == false && strings.HasSuffix(fileName, ".dat") == false && strings.HasSuffix(fileName, ".bin") == false) {
        return nil
    }

    if (strings.HasSuffix(fileName, ".csv") == true) {
        reader := new(SotECSVFile)
    }
}

// Reading files requires checking most calls for errors.
// This helper will streamline our error checks below.
func check(e error) {
    if e != nil {
        panic(e)
    }
}

func readNextBytes(b *bufio.Reader, number int) []byte {
	bytes := make([]byte, number)

	_, err := b.Read(bytes)
	check(err)

	return bytes
}