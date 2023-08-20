package main

type Geometry struct {
	PolarAngle       float64  `json:"polarAngle"`
	InclinationAngle float64  `json:"inclinationAngle"`
	RadialDistance   float64  `json:"radialDistance"`
	ID               string   `json:"id"`
	Neighbors        []string `json:"neighbors"`
}

type Geometries struct {
	Geometry []Geometry `json:"Geometry"`
}
