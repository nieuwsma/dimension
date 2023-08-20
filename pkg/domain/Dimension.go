package domain

import (
	"errors"
	"fmt"
	"sort"
	"strings"
)

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

func NewDimension(pairs ...SpherePair) (*Dimension, error) {
	dim := &Dimension{
		Dimension: make(map[string]Sphere),
	}

	for _, pair := range pairs {
		dim.Dimension[string(pair.ID)] = pair.Sphere
	}

	err := dim.ValidateGeometry()
	if err == nil {
		err = dim.ValidateSpheres()
	}

	return dim, err
}

type Dimension struct {
	Dimension map[string]Sphere
}

func (d *Dimension) String() string {
	var entries []string

	// Loop through the dimension map and collect all entries in a slice
	for id, sphere := range d.Dimension {
		entries = append(entries, fmt.Sprintf("%s: %s", id, sphere.Color.LongHand()))

	}

	// Sort the entries for consistency (if needed)
	sort.Strings(entries)

	// Join all the entries with a newline separator and return
	return strings.Join(entries, "\n")
}

func (d *Dimension) ValidateSpheres() error {
	colorCounts := make(map[Color]int)

	for _, sphere := range d.Dimension {

		// Increment the color count for the sphere's color
		colorCounts[sphere.Color]++

		// Check for any color exceeding the limit of 3
		if colorCounts[sphere.Color] > 3 {
			//exceeded color count
			return fmt.Errorf("%s has %d spheres, maximum is 3", sphere.Color.LongHand(), colorCounts[sphere.Color])
		}

	}

	return nil
}

func (d *Dimension) ValidateGeometry() error {
	//the length cannot be bigger than 11
	if len(d.Dimension) > 11 {
		return errors.New("too many spheres")
	}

	// Rule 1: a can always be present
	// (This rule doesn't need explicit code)

	// Rule 2: Check spheres from the first outer ring
	// (This rule also doesn't need explicit code as they always can be present)

	// Rule 3: Check Tropical Ring
	tropicalRing := []string{"h", "i", "j", "k", "l", "m"}
	tropicalPresent := make(map[string]bool)
	for _, t := range tropicalRing {
		if _, ok := d.Dimension[t]; ok {
			tropicalPresent[t] = true
		}
	}
	if len(tropicalPresent) > 0 {
		if _, ok := d.Dimension["a"]; !ok {
			return errors.New("missing center sphere")
		}
	}

	// Check if required neighbors are present
	requiredNeighbors := map[string][]string{
		"h": {"b", "c"},
		"i": {"c", "d"},
		"j": {"d", "e"},
		"k": {"e", "f"},
		"l": {"f", "g"},
		"m": {"g", "b"},
	}
	for t, _ := range tropicalPresent {
		neighbors := requiredNeighbors[t]
		for _, neighbor := range neighbors {
			if _, ok := d.Dimension[neighbor]; !ok {
				return fmt.Errorf("%s is missing a required neighbor %s", t, neighbor)
			}
		}
	}

	// Check the configuration of the tropical ring
	validConfigs := [][]string{
		{},                                       // No spheres are present
		{"h"}, {"j"}, {"l"}, {"i"}, {"k"}, {"m"}, // Only one sphere is present
		{"h", "j", "l"},
		{"i", "k", "m"},
		{"h", "k"},
		{"i", "l"},
		{"j", "m"},
	}
	matchedConfig := false
	for _, config := range validConfigs {
		allPresent := true
		for _, s := range config {
			if _, ok := tropicalPresent[s]; !ok {
				allPresent = false
				break
			}
		}
		if allPresent && len(tropicalPresent) == len(config) {
			matchedConfig = true
			break
		}
	}
	if !matchedConfig {
		return errors.New("invalid tropical ring configuration")
	}

	// Rule 4: Check Top Sphere
	if _, ok := d.Dimension["n"]; ok {
		if _, ok := d.Dimension["a"]; !ok {
			return errors.New("missing center sphere")
		}
		for _, s := range []string{"b", "c", "d", "e", "f", "g"} {
			if _, ok := d.Dimension[s]; !ok {
				return fmt.Errorf("missing a equitorial ring sphere neighbor %s", s)
			}
		}
		if !matchedConfig {
			return errors.New("invalid tropical ring configuration")
		}
	}

	return nil
}
