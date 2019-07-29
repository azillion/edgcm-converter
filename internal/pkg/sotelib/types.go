package sotelib

type SotEDataFile struct {
	Length int32
	Cells []*SotEData
}

type SotEData struct {
    Water float32
    Elevation float32
    Tempature float32
    Humidity float32
    Rainfall float32
}

// gonna add lat and long here
type Climate struct {
    Water float32
    Elevation float32
    Tempature float32
    Humidity float32
    Rainfall float32
}