package main

import (
	"errors"
	"fmt"
	"github.com/nieuwsma/dimension/pkg/geometry"
	"github.com/nieuwsma/dimension/pkg/logic"
	"strconv"
	"strings"
)

// a few thoughts on approaches
// i could start by submitting an empty dimension, and seeing what i fail
// i could start by trying to figure out what rules are in conflict
// i could start by always working with quantity rules first: exact, sum, gt

//some good heuristics,
// if a color is repeated a lot across many colors, it might be easier to just omit it
// if there is a touch-G-K; then a good pattern is to do GKGKGK around the equator.
func main() {

	trainingSession := logic.NewTrainingSession(6, 12345)

	dimension, _ := logic.NewDimension()
	trainingSession.PlayTurn("autopilot", *dimension)

	fmt.Println(trainingSession.Turns["autopilot"])

}

//there are 7 types of tasks, what are the ones that interact with eachother?

//touch, notouch
//U & N interact if the are the same, or if any color is the same
//U & U interacts if any color is the same
//N & N interacts if any color is the same

//bottom,top
// B & T interact if the same color is listed
// B & B interact if there arent any more places
// T & T interact if there arent any more places

//quantity, sum, gt
// Q & S interact if any color in sum is Q color
// Q & Q only interact if they are the same color
// S & S interact if any color is in both pairs
// Q & G interact if any color is in Q
// G & S interact if any color is in either
// G & G interact if any color is in both pairs


func CategorizeTasks(tasks logic.Tasks (score int, bonus bool, taskViolations []string, errs error) {


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
		var newTasks logic.Tasks
		if quantity == 3 {
			find1 := fmt.Sprintf("QUANTITY-1-%s", color)
			find2 := fmt.Sprintf("QUANTITY-2-%s", color)
			for _, task := range tasks {
				if !strings.Contains(string(task), find1) && !strings.Contains(string(task), find2) {
					newTasks = append(newTasks, task)
				}
			}
			newTask := logic.Task(fmt.Sprintf("QUANTITY-3-%s", color))
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
				//todo need a different error struct for actual errors, not task violations
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
