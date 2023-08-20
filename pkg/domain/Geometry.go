package domain

import (
	_ "embed"
	"encoding/json"
)

type Geometry struct {
	PolarAngle       int      `json:"polarAngle"`
	InclinationAngle int      `json:"inclinationAngle"`
	RadialDistance   int      `json:"radialDistance"`
	ID               string   `json:"id"`
	Neighbors        []string `json:"neighbors"`
}

type Geometries struct {
	Geometry []Geometry `json:"Geometry"`
}

//go:embed geometry.json
var GeometryData []byte

func GetGeometry() (g Geometries) {
	_ = json.Unmarshal(GeometryData, &g)
	return
}
