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
	Latitude float32
	// WaterPercentage 0-1, percentage
	WaterPercentage float32
	// Elevation meters
	Elevation float32
	// Temperature degrees Celsius
	Temperature float32
	// Humidity relative humidity, 0-1, perentage
	HumidityPercentage float32
	// Rainfall absolute amount, millimeter
	Rainfall float32
}

// SeaIce Naively determine percentage of cell is sea ice
func (c *Climate) SeaIce() float32 {
	if c.WaterPercentage > 0.0 && c.Temperature < -1.8 {
		return c.WaterPercentage
	}
	return 0.0
}

// IsLandIce Naively determine if cell has land ice
func (c *Climate) IsLandIce() bool {
	if c.WaterPercentage == 0.0 && c.Temperature < -1.8 && c.Rainfall > 0.0 {
		return true
	}
	return false
}
