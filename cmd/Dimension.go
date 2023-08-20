package main

import (
	"fmt"
	"sort"
	"strings"
)

type Sphere struct {
	Color Color
}

// NewDimension is a factory function for creating a Dimension with its map initialized.
func NewDimension() *Dimension {
	return &Dimension{
		Dimension: make(map[string]*Sphere),
	}
}

type Dimension struct {
	Dimension map[string]*Sphere
}

func (d *Dimension) String() string {
	var entries []string

	// Loop through the dimension map and collect all entries in a slice
	for id, sphere := range d.Dimension {
		if sphere != nil {
			entries = append(entries, fmt.Sprintf("%d: %s", id, sphere.Color.LongHand()))
		} else {
			entries = append(entries, fmt.Sprintf("%d: nil", id))
		}
	}

	// Sort the entries for consistency (if needed)
	sort.Strings(entries)

	// Join all the entries with a newline separator and return
	return strings.Join(entries, "\n")
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

func (d *Dimension) IsValid() bool {
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
			return false
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
				return false
			}
		}
	}

	// Check the configuration of the tropical ring
	validConfigs := [][]string{
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
		return false
	}

	// Rule 4: Check Top Sphere
	if _, ok := d.Dimension["n"]; ok {
		if _, ok := d.Dimension["a"]; !ok {
			return false
		}
		for _, s := range []string{"b", "c", "d", "e", "f", "g"} {
			if _, ok := d.Dimension[s]; !ok {
				return false
			}
		}
		if !matchedConfig {
			return false
		}
	}

	return true
}
