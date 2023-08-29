package logic

import (
	"errors"
	"fmt"
	"sort"
	"strings"
)

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

	// Task 2: Check spheres from the first outer ring
	// (This rule also doesn't need explicit code as they always can be present)

	// Task 3: Check Tropical Ring
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
		{"h"}, {"i"}, {"j"}, {"k"}, {"l"}, {"m"}, // Only one sphere is present
		{"h", "j", "l"},
		{"i", "k", "m"},

		{"h", "k"},
		{"h", "j"},
		{"h", "l"},

		{"i", "k"},
		{"i", "l"},
		{"i", "m"},

		{"j", "l"},
		{"j", "m"},
		{"j", "h"}, //duplicate

		{"k", "m"},
		{"k", "h"}, //duplicate
		{"k", "i"}, //duplicate

		{"l", "h"}, //duplicate
		{"l", "i"}, //duplicate
		{"l", "j"}, //duplicate

		{"m", "i"}, //duplicate
		{"m", "j"}, //duplicate
		{"m", "k"}, //duplicate
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

	// Task 4: Check Top Sphere
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
