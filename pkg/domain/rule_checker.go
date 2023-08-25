package domain

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func ScoreTurn(tasks []string, dim Dimension) (score int, bonus bool, errs error) {
	maxScore := len(dim.Dimension)
	colorCounts := make(map[Color]int)
	for _, v := range dim.Dimension {
		colorCounts[v.Color]++
	}

	if err := dim.ValidateGeometry(); err != nil { // illegal dimension
		score = 0
		errs = errors.Join(err)
		return
	}

	if err := dim.ValidateSpheres(); err != nil { // illegal dimension
		score = 0
		errs = errors.Join(err)
		return
	}

	score = maxScore

	//special rule, if there is a 2 & 1 quantity task for the same color, add them!
	var colorQuantity = make(map[string]int)
	for _, task := range tasks {
		if strings.Contains(task, "QUANTITY") {
			parts := strings.Split(task, "-")
			count, _ := strconv.Atoi(parts[1])
			colorQuantity[parts[2]] += count
		}

	}
	for color, quantity := range colorQuantity {
		var newTasks []string
		if quantity == 3 {
			find1 := fmt.Sprintf("QUANTITY-1-%s", color)
			find2 := fmt.Sprintf("QUANTITY-2-%s", color)
			for _, task := range tasks {
				if !strings.Contains(task, find1) && !strings.Contains(task, find2) {
					newTasks = append(newTasks, task)
				}
			}
			newTask := fmt.Sprintf("QUANTITY-3-%s", color)
			newTasks = append(newTasks, newTask)
			tasks = newTasks
		}
	}

	for _, task := range tasks {

		parts := strings.Split(task, "-")

		switch {
		case strings.Contains(task, "QUANTITY"):
			//todo need to check for the special case that QUANTITY-1-n and QUANTITY-2-n have been played, if they do then sum is 3
			quantity, err := strconv.Atoi(parts[1])
			if err != nil {
				err = fmt.Errorf("Could not parse task %s", task)
				errs = errors.Join(err)
				score -= 2
			} else {
				err = CheckQuantity(quantity, NewColorShort(parts[2]), colorCounts)
				if err != nil {
					errs = errors.Join(err)
					score -= 2
				}
			}
		case strings.Contains(task, "BOTTOM"):
			err := CheckTopBottom(dim, false, NewColorShort(parts[1]), GetGeometry().GetNeighbors())
			if err != nil {
				errs = errors.Join(err)
				score -= 2
			}
		case strings.Contains(task, "TOP"):
			err := CheckTopBottom(dim, true, NewColorShort(parts[1]), GetGeometry().GetNeighbors())
			if err != nil {
				errs = errors.Join(err)
				score -= 2
			}
		case strings.Contains(task, "TOUCH"):
			var err error
			if strings.Contains(task, "NOTOUCH") {
				err = CheckTouch(dim, colorCounts, false, NewColorShort(parts[1]), NewColorShort(parts[2]), GetGeometry().GetNeighbors())
			} else {
				err = CheckTouch(dim, colorCounts, true, NewColorShort(parts[1]), NewColorShort(parts[2]), GetGeometry().GetNeighbors())
			}

			if err != nil {
				errs = errors.Join(err)
				score -= 2
			}
		case strings.Contains(task, "RATIO"):
			quantity, err := strconv.Atoi(parts[1])

			if err != nil {
				err = fmt.Errorf("Could not parse task %s", task)
				errs = errors.Join(err)
				score -= 2
			} else {
				var colors []Color
				colors = append(colors, NewColorShort(parts[2]))
				colors = append(colors, NewColorShort(parts[3]))
				err = CheckRatio(quantity, colors, colorCounts)
				if err != nil {
					errs = errors.Join(err)
					score -= 2
				}
			}
		case strings.Contains(task, "GT"):
			err := CheckGreaterThan(NewColorShort(parts[1]), NewColorShort(parts[2]), colorCounts)
			if err != nil {
				errs = errors.Join(err)
				score -= 2
			}
		default:
			err := fmt.Errorf("Could not parse task %s", task)
			if err != nil {
				errs = errors.Join(err)
				score -= 2
			}
		}
	}

	//a bonus is awarded if all tasks were successfully completed and if all 5 colors were used.
	if errs == nil && len(colorCounts) == 5 { //then
		bonus = true

	}
	return score, bonus, errs
}

func CheckQuantity(quantity int, color Color, colorCounts map[Color]int) (err error) {

	if colorCounts[color] != quantity {
		err = fmt.Errorf("expected %d, got %d", quantity, colorCounts[color])
		return
	}
	return nil
}

func CheckRatio(sum int, colors []Color, colorCounts map[Color]int) (err error) {
	runningTotal := 0
	for _, color := range colors {
		runningTotal += colorCounts[color]
	}
	if runningTotal != sum {
		return fmt.Errorf("expected sum of %d, got %d", sum, runningTotal)
	}
	return nil
}

func CheckGreaterThan(gt Color, lt Color, colorCounts map[Color]int) (err error) {
	if colorCounts[lt] >= colorCounts[gt] {
		return fmt.Errorf("count of %s exceeds count of %s", lt.String(), gt.String())
	}
	return nil
}

// the offical rules state that for touch or no-touch that only playing one color is allowed,
// so only playing one W on a W-W or only playing one W and no K on  a W-K scenario is acceptable
func CheckTouch(dim Dimension, colorCounts map[Color]int, mustTouch bool, a Color, b Color, neighbors Neighbors) (err error) {

	if a.Equals(b) && colorCounts[a] <= 1 {
		return nil
	} else if !a.Equals(b) && (colorCounts[a] == 0 || colorCounts[b] == 0) {
		return nil
	}
	for position, sphere := range dim.Dimension {
		if !sphere.Color.Equals(a) && !sphere.Color.Equals(b) { // this sphere has neither color we care about!
			continue
		}
		currentColor := sphere.Color
		var matchColor Color
		if currentColor.Equals(a) { //we care about currentColor vs matchColor, we can get to this from either a or b color
			matchColor = b
		} else {
			matchColor = a
		}
		touched := false //if mustTouch is true, then touch starts as false
		//if mustTouch is false, then touch starts as true

		//for each neighbor in the neighborhood
		//if the neighbor is set, check its color see if it's the match color(we don't need to check the geometry, that has already been done!)
		neighborhood := neighbors[position]
		for _, neighbor := range neighborhood {
			if neighborSphere, exists := dim.Dimension[neighbor]; exists {
				if neighborSphere.Color.Equals(matchColor) { //we have a color match
					touched = true
				}
			}

		}
		if touched != mustTouch { //check if the touched trigger was tripped, compare it to mustTouch
			var message string
			if mustTouch == false {
				message = "has a neighbor"
			} else {
				message = "failed to find a neighbor"
			}
			err = errors.Join(fmt.Errorf("%s which is %s %s who is %s", position, currentColor.LongHand(), message, matchColor.LongHand()))
		}
	}
	return err
}

// Ensure no sphere of any color may be above & touching color sphere
func CheckTopBottom(dim Dimension, top bool, a Color, neighbors Neighbors) (err error) {

	plane := make(map[string]int)
	plane["a"] = 0
	plane["b"] = 0
	plane["c"] = 0
	plane["d"] = 0
	plane["e"] = 0
	plane["f"] = 0
	plane["g"] = 0
	plane["h"] = 1
	plane["i"] = 1
	plane["j"] = 1
	plane["k"] = 1
	plane["l"] = 1
	plane["m"] = 1
	plane["n"] = 2
	//need to check the ontop rules.. is it on top ONLY if its touching? or if its in the higher order plane? answer, ONLY if its touching.
	//so a,b,c,f,h; f doesnt touch h; so it cannot break the top/bottom rules
	//this makes me think that If i check each sphere, and check its neighbors, that i can just say what comes after it,
	//that isnt a peer is on top, if it doesnt touch it, but is on top of it, then it doesn't matter

	for position, sphere := range dim.Dimension {
		//for every sphere, if its the color, check what is on top of it
		if sphere.Color.Equals(a) {
			//now check position.
			currentPlane := plane[position]
			neighborhood := neighbors[position]
			for _, neighbor := range neighborhood {
				neighborPlane := plane[neighbor]
				if _, exists := dim.Dimension[neighbor]; exists {
					if top && neighborPlane > currentPlane {
						return fmt.Errorf("position %s (color %s), has neighbor %s  that is on top of it", position, sphere.Color.LongHand(), neighbor)
					} else if !top && neighborPlane < currentPlane {
						return fmt.Errorf("position %s (color %s), has neighbor %s that is below it", position, sphere.Color.LongHand(), neighbor)
					}
				}
			}
		}
	}

	return err
}
