package climate

// WorldClimate WorldClimate
type WorldClimate struct {
	Length   int32
	GridSize float32
	Cells    []*Cell
}

// Cell Cell
type Cell struct {
	CellID  int32
	Climate Climate
}

// Climate Climate
type Climate struct {
	Latitude        float32
	WaterPercentage float32
	Elevation       float32
	Tempature       float32
	Humidity        float32
	Rainfall        float32
}
