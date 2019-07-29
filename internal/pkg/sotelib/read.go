package sotelib

import (
    "bufio"
    "math"
    "fmt"
    // "io"
    // "io/ioutil"
    "bytes"
    "encoding/binary"
    "os"
)

//     wr.Write(w.climateElevation.Length); //int
// for (int i = 0; i < w.climateElevation.Length; i++)
// {
// 	wr.Write(w.climateWaterPercentage[i]); //float, 0-1, percentage
// 	wr.Write(w.climateElevation[i]); //float, meters
// 	wr.Write(w.climateJanTemperature[i]); //float, degrees Celsius
// 	wr.Write(w.climateJanHumidity[i]); //float, relative humidity, 0-1, perentage
// 	wr.Write(w.climateJanRainfall[i]); //float, absolute amount, mm
// }
// data is in a 2 dimensional array
// with the same number of rows and columns


// Reading files requires checking most calls for errors.
// This helper will streamline our error checks below.
func check(e error) {
    if e != nil {
        panic(e)
    }
}

// ReadDat read a dat file
func ReadDat(f *os.File) []*SotEData {
    defer f.Close()

    reader := bufio.NewReader(f)

    // num of rows in the file
    var numOfCells int32
    numOfCellsBytes := readNextBytes(reader, 4)
    buffer := bytes.NewBuffer(numOfCellsBytes)

    // convert buffer to int32
    err := binary.Read(buffer, binary.LittleEndian, &numOfCells)
    check(err)

    gridsize := math.Sqrt(float64(numOfCells))
    fmt.Printf("len: %d, gridsize: %f", numOfCells, gridsize)

    soteCells := []*SotEData{}
    for i := 0; i < int(numOfCells); i++ {
        soteClimate := new(SotEData)
        data := readNextBytes(reader, 20)
        buffer2 := bytes.NewBuffer(data)
        err := binary.Read(buffer2, binary.LittleEndian, soteClimate)
        check(err)

        soteCells = append(soteCells, soteClimate)
    }

    return soteCells
}

func parseSotEData(soteData *SotEData, data []byte) *SotEData {
    // soteData := new(SotEData)
    // data := readNextBytes(reader, 4)
    buffer := bytes.NewBuffer(data)
    // data = readNextBytes(reader, 4)
    // buffer2.Write(data)
    // data = readNextBytes(reader, 4)
    // buffer2.Write(data)
    // data = readNextBytes(reader, 4)
    // buffer2.Write(data)
    // data = readNextBytes(reader, 4)
    // buffer2.Write(data)
   
    err := binary.Read(buffer, binary.LittleEndian, soteData)
    check(err)

    // fmt.Printf("soteData:%+v\n", soteData)
    return soteData
}

func readNextBytes(b *bufio.Reader, number int) []byte {
	bytes := make([]byte, number)

	_, err := b.Read(bytes)
	check(err)

	return bytes
}