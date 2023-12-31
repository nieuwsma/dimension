package logic

import (
	"errors"
	"fmt"
	"github.com/nieuwsma/dimension/pkg/geometry"
	"strconv"
	"strings"
)

func ScoreTurn(tasks Tasks, dim Dimension) (score int, bonus bool, taskViolations []string, errs error) {
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
		if strings.Contains(string(task), "QUANTITY") {
			parts := strings.Split(string(task), "-")
			count, _ := strconv.Atoi(parts[1])
			colorQuantity[parts[2]] += count
		}

	}
	for color, quantity := range colorQuantity {
		var newTasks Tasks
		if quantity == 3 {
			find1 := fmt.Sprintf("QUANTITY-1-%s", color)
			find2 := fmt.Sprintf("QUANTITY-2-%s", color)
			for _, task := range tasks {
				if !strings.Contains(string(task), find1) && !strings.Contains(string(task), find2) {
					newTasks = append(newTasks, task)
				}
			}
			newTask := Task(fmt.Sprintf("QUANTITY-3-%s", color))
			newTasks = append(newTasks, newTask)
			tasks = newTasks
		}
	}

	for _, task := range tasks {

		parts := strings.Split(string(task), "-")

		switch {
		case strings.Contains(string(task), "QUANTITY"):
			quantity, err := strconv.Atoi(parts[1])
			if err != nil {
				taskViolations = append(taskViolations, fmt.Sprintf("Could not parse task %s", task))
				score -= 2
			} else {
				err = CheckQuantity(quantity, NewColorShort(parts[2]), colorCounts)
				if err != nil {
					taskViolations = append(taskViolations, fmt.Sprintf(err.Error()))
					score -= 2
				}
			}
		case strings.Contains(string(task), "BOTTOM"):
			err := CheckTopBottom(dim, false, NewColorShort(parts[1]), geometry.GetGeometry().GetNeighbors())
			if err != nil {
				taskViolations = append(taskViolations, fmt.Sprintf(err.Error()))
				score -= 2
			}
		case strings.Contains(string(task), "TOP"):
			err := CheckTopBottom(dim, true, NewColorShort(parts[1]), geometry.GetGeometry().GetNeighbors())
			if err != nil {
				taskViolations = append(taskViolations, fmt.Sprintf(err.Error()))
				score -= 2
			}
		case strings.Contains(string(task), "TOUCH"):
			var err error
			if strings.Contains(string(task), "NOTOUCH") {
				err = CheckTouch(dim, colorCounts, false, NewColorShort(parts[1]), NewColorShort(parts[2]), geometry.GetGeometry().GetNeighbors())
			} else {
				err = CheckTouch(dim, colorCounts, true, NewColorShort(parts[1]), NewColorShort(parts[2]), geometry.GetGeometry().GetNeighbors())
			}

			if err != nil {
				taskViolations = append(taskViolations, fmt.Sprintf(err.Error()))
				score -= 2
			}
		case strings.Contains(string(task), "SUM"):
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
					taskViolations = append(taskViolations, fmt.Sprintf(err.Error()))
					score -= 2
				}
			}
		case strings.Contains(string(task), "GT"):
			err := CheckGreaterThan(NewColorShort(parts[1]), NewColorShort(parts[2]), colorCounts)
			if err != nil {
				taskViolations = append(taskViolations, fmt.Sprintf(err.Error()))
				score -= 2
			}
		default:
			err := fmt.Errorf("Could not parse task %s", task)
			if err != nil {
				taskViolations = append(taskViolations, fmt.Sprintf(err.Error()))
				score -= 2
			}
		}
	}

	//a bonus is awarded if all tasks were successfully completed and if all 5 colors were used.
	if len(taskViolations) == 0 && len(colorCounts) == 5 { //then
		bonus = true

	}

	if score < 0 {
		score = 0
	}
	return score, bonus, taskViolations, errs
}

func CheckQuantity(quantity int, color Color, colorCounts map[Color]int) (err error) {

	if colorCounts[color] != quantity {
		err = fmt.Errorf("expected %d %s, got %d", quantity, color.LongHand(), colorCounts[color])
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
		return fmt.Errorf("expected sum of %v = %d, got %d", colors, sum, runningTotal)
	}
	return nil
}

func CheckGreaterThan(gt Color, lt Color, colorCounts map[Color]int) (err error) {
	if colorCounts[lt] >= colorCounts[gt] {
		return fmt.Errorf("count of %s must be greater than count of %s", gt.String(), lt.String())
	}
	return nil
}

// the offical rules state that for touch or no-touch that only playing one color is allowed,
// so only playing one W on a W-W or only playing one W and no K on  a W-K scenario is acceptable
func CheckTouch(dim Dimension, colorCounts map[Color]int, mustTouch bool, a Color, b Color, neighbors geometry.Neighbors) (err error) {

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
			err = errors.Join(fmt.Errorf("'%s' which is %s %s who is %s", position, currentColor.LongHand(), message, matchColor.LongHand()))
		}
	}
	return err
}

// Ensure no sphere of any color may be above & touching color sphere
func CheckTopBottom(dim Dimension, top bool, a Color, neighbors geometry.Neighbors) (err error) {

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
						return fmt.Errorf("position '%s' (color %s), has neighbor '%s'  that is on top of it", position, sphere.Color.LongHand(), neighbor)
					} else if !top && neighborPlane < currentPlane {
						return fmt.Errorf("position '%s' (color %s), has neighbor '%s' that is below it", position, sphere.Color.LongHand(), neighbor)
					}
				}
			}
		}
	}

	return err
}
