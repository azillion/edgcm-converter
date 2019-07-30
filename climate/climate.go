package climate

type WorldClimate struct {
    Length int32
    GridSize float32
	Cells []*Cell
}

type Cell struct {
    CellId int32
    Climate Climate
}

type Climate struct {
    Latitude float32
    WaterPercentage float32
    Elevation float32
    Tempature float32
    Humidity float32
    Rainfall float32
}