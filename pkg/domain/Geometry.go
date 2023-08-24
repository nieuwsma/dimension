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

type Neighbors map[string][]string

//go:embed geometry.json
var GeometryData []byte

func GetGeometry() (g Geometries) {
	_ = json.Unmarshal(GeometryData, &g)
	return
}

func (g Geometries) GetNeighbors() (n Neighbors) {
	n = make(map[string][]string)
	for _, v := range g.Geometry {
		n[v.ID] = v.Neighbors
	}
	return n
}
