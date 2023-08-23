package domain

import (
	"errors"
	"fmt"
	"sort"
)

func Dispatcher(rule string, dim Dimension) {
	colorCounts := make(map[Color]int)
	for _, v := range dim.Dimension {
		colorCounts[v.Color]++
	}
}

func ValidateQuantity(quantity int, color Color, colorCounts map[Color]int) (err error) {

	if colorCounts[color] != quantity {
		return fmt.Errorf("expected %d, got %d", quantity, colorCounts[color])
	}
	return nil
}

func ValidateRatio(sum int, colors []Color, colorCounts map[Color]int) (err error) {
	runningTotal := 0
	for _, color := range colors {
		runningTotal += colorCounts[color]
	}
	if runningTotal > sum {
		return fmt.Errorf("expected sum of %d, got %d", sum, runningTotal)
	}
	return nil
}

func ValidateGreaterThan(gt Color, lt Color, colorCounts map[Color]int) (err error) {
	if colorCounts[lt] >= colorCounts[gt] {
		return fmt.Errorf("count of %s exceeds count of %s", lt.String(), gt.String())
	}
	return nil
}

// start with IT MUST touch
// todo im not sure this boolean inversion logic is correct for both cases, need to unit test!
// I think what is missing is that a MUST touch has to check every instance has a touch, vs a NOT touch has to make sure that no instance touches
func ValidateTouch(dim Dimension, mustTouch bool, a Color, b Color, neighbors Neighbors) (err error) {

	for position, sphere := range dim.Dimension {
		if !sphere.Color.Equals(a) && !sphere.Color.Equals(b) { // lets check against A first
			continue
		}
		currentColor := sphere.Color
		var matchColor Color
		if currentColor.Equals(a) {
			matchColor = b
		} else {
			matchColor = a
		}
		touch := !mustTouch //false = !true #it must touch in my walkthrough
		//true = !false

		neighborhood := neighbors[position]
		for _, neighbor := range neighborhood {
			if neighborSphere, exists := dim.Dimension[neighbor]; exists {
				if neighborSphere.Color.Equals(matchColor) {
					touch = mustTouch // true = true |
				}
			}
			if touch != mustTouch { //false
				err = errors.Join(fmt.Errorf("%s which is %s failed to find a neighbor who is %s", position, currentColor.LongHand(), matchColor.LongHand()))
			}
		}
	}
	return err
}

// Ensure no sphere of any color may be above color sphere
//todo need to play around with how to iterate through this
func ValidateOnTop(dim Dimension, a Color) (err error) {

	//make a slice of the dimension keys, put it in order
	keys := make([]string, 0, len(dim.Dimension))
	for k := range dim.Dimension {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	keyPosition:= make(map[string]int)
	for index, position := range keys {
		keyPosition[position]=index
	}

	for position, sphere := range dim.Dimension {
		//for every sphere, if its the color, then nothing can be in a later position that would put it in a higher ring

		if sphere.Color.Equals(a) {
			//now check position.
			if position
		}
	}

	return err
}
