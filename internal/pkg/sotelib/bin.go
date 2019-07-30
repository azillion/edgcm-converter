package sotelib



// ClimateReader read a SotE Bin file
func (f *SotEBinFile) Read() {
    reader := bufio.NewReader(f.sotEFile.InputFile)

    // num of rows in the file
    var numOfCells int32
    numOfCellsBytes := readNextBytes(reader, 4)
    buffer := bytes.NewBuffer(numOfCellsBytes)

    // convert buffer to int32
    err := binary.Read(buffer, binary.LittleEndian, &numOfCells)
    check(err)

    f.sotEFile.WorldClimate.GridSize = math.Sqrt(float64(numOfCells))
    f.sotEFile.WorldClimate.Length = numOfCells
    fmt.Printf("len: %d\ngridsize: %f\n", f.sotEFile.WorldClimate.Length, f.sotEFile.WorldClimate.GridSize)

    f.sotEFile.WorldClimate.Cells = []*Cell{}
    for i := 0; i < int(c.Length); i++ {
        cell := new(Cell)
        data := readNextBytes(reader, 24)
        buffer2 := bytes.NewBuffer(data)
        err := binary.Read(buffer2, binary.LittleEndian, cell)
        check(err)
		// fmt.Printf("cell:%+v\n", cell)
		cell.CellId = i
        f.sotEFile.WorldClimate.Cells = append(f.sotEFile.WorldClimate.Cells, cell)
    }
}

func SotEBinFile(f *os.File) *SotEBinFile {
	soteFile := SotEBinFile{File: f}
	return &soteFile
}