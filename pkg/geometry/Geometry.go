package geometry

import (
	_ "embed"
)

var geometryData Geometries = []Geometry{
	{0, 0, 1, "b", []string{"c", "a", "g", "m", "h"}},
	{60, 0, 1, "c", []string{"b", "a", "d", "h", "i"}},
	{120, 0, 1, "d", []string{"c", "a", "e", "i", "j"}},
	{180, 0, 1, "e", []string{"d", "a", "f", "j", "k"}},
	{240, 0, 1, "f", []string{"g", "a", "e", "k", "l"}},
	{300, 0, 1, "g", []string{"b", "a", "f", "l", "m"}},
	{30, 45, 1, "h", []string{"a", "b", "c", "n", "l", "j"}},
	{90, 45, 1, "i", []string{"a", "c", "d", "m", "k", "n"}},
	{150, 45, 1, "j", []string{"a", "e", "d", "h", "l", "n"}},
	{210, 45, 1, "k", []string{"a", "e", "f", "i", "m", "n"}},
	{270, 45, 1, "l", []string{"g", "f", "a", "h", "j", "n"}},
	{330, 45, 1, "m", []string{"g", "b", "a", "k", "i", "n"}},
	{0, 0, 0, "a", []string{"b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m"}},
	{0, 90, 2, "n", []string{"h", "i", "j", "k", "l", "m"}},
}

type Geometry struct {
	PolarAngle       int      `json:"polarAngle"`
	InclinationAngle int      `json:"inclinationAngle"`
	RadialDistance   int      `json:"radialDistance"`
	ID               string   `json:"id"`
	Neighbors        []string `json:"neighbors"`
}

type Geometries []Geometry
type Neighbors map[string][]string

func GetGeometry() Geometries {
	return geometryData
}

func (g Geometries) GetNeighbors() (n Neighbors) {
	n = make(map[string][]string)
	for _, v := range g {
		n[v.ID] = v.Neighbors
	}
	return n
}
