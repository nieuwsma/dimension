package domain

type SphereID string

const (
	A SphereID = "a"
	B SphereID = "b"
	C SphereID = "c"
	D SphereID = "d"
	E SphereID = "e"
	F SphereID = "f"
	G SphereID = "g"
	H SphereID = "h"
	I SphereID = "i"
	J SphereID = "j"
	K SphereID = "k"
	L SphereID = "l"
	M SphereID = "m"
	N SphereID = "n"
)

type SpherePair struct {
	ID     SphereID
	Sphere Sphere
}

func NewSpherePair(ID SphereID, color Color) *SpherePair {
	return &SpherePair{Sphere: Sphere{Color: color}, ID: ID}
}

func NewSphere(color Color) *Sphere {
	return &Sphere{Color: color}
}

type Sphere struct {
	Color Color
}
