package main

type Sphere struct {
	Color Color
}

// NewDimension is a factory function for creating a Dimension with its map initialized.
func NewDimension() *Dimension {
	return &Dimension{
		Dimension: make(map[int]*Sphere),
	}
}

type Dimension struct {
	Dimension map[int]*Sphere
}

// todo return an err instead of bool
func (d *Dimension) ValidateGeometry() bool {
	count := 0
	colorCounts := make(map[Color]int)

	for _, sphere := range d.Dimension {
		if sphere != nil {
			count++

			// Increment the color count for the sphere's color
			colorCounts[sphere.Color]++

			// Check for any color exceeding the limit of 3
			if colorCounts[sphere.Color] > 3 {
				//exceeded color count
				return false
			}
		}
	}

	if count > 11 {
		return false //too many spheres
	}

	return count <= 11
}

type Rules struct {
	Rules map[int]Rule
}
type Rule struct {
	Name        string
	Description string
}

func (g Color) LongHand() string {
	switch g {
	case Green:
		return "GREEN"
	case Blue:
		return "BLUE"
	case Black:
		return "BLACK"
	case Orange:
		return "ORANGE"
	case White:
		return "WHITE"
	default:
		return ""
	}
}

type Color int

const (
	Green  Color = 1
	Blue   Color = 2
	White  Color = 3
	Orange Color = 4
	Black  Color = 5
)

func (s Color) String() string {
	switch s {
	case Green:
		return "G"
	case Blue:
		return "B"
	case Black:
		return "K"
	case Orange:
		return "O"
	case White:
		return "W"
	default:
		return ""
	}
}

func (s Color) Equals(other Color) bool {
	return s == other
}
