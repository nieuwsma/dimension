package main

import (
	"fmt"
	"github.com/nieuwsma/dimension/pkg/logic"
	"strconv"
	"strings"
)

// a few thoughts on approaches
// i could start by submitting an empty dimension, and seeing what i fail
// i could start by trying to figure out what rules are in conflict
// i could start by always working with quantity rules first: exact, sum, gt

// some good heuristics,
// if a color is repeated a lot across many colors, it might be easier to just omit it
// if there is a touch-GreaterThan-K; then a good pattern is to do GKGKGK around the equator.
func main() {

	trainingSession := logic.NewTrainingSession(6, 12345)

	dimension, _ := logic.NewDimension()
	trainingSession.PlayTurn("autopilot", *dimension)

	categorizedTasks, _ := CategorizeTasks(trainingSession.Tasks)
	fmt.Println(categorizedTasks)

	fmt.Println(trainingSession.Turns["autopilot"])

}

func ProcessCategorizedTasks(t TasksCollection) (interactions map[string][]string) {
	//TOUCHES
	if len(t.Touch) > 1 {
		//Multiple Touches
	}
	if len(t.NoTouch) > 1 {
		//multiple NoTouches
	}

	if len(t.Touch) >= 1 && len(t.NoTouch) >= 1 {
		//NoTouch and Touch
	}

	//SITS
	if len(t.Bottom) > 1 {
		//Multiple Bottom
	}
	if len(t.Top) > 1 {
		//multiple Top
	}

	if len(t.Top) >= 1 && len(t.Bottom) >= 1 {
		//Top and Bottom
	}

	//COUNTS
	if len(t.Quantity) > 1 {
		//Multiple Quantity
	}
	if len(t.Sum) > 1 {
		//multiple Sum
	}

	if len(t.GreaterThan) > 1 {
		//Multiple GreaterThan
	}

	if len(t.Quantity) >= 1 && len(t.Sum) >= 1 {
		//Quantity and Sum
	}

	if len(t.Sum) >= 1 && len(t.GreaterThan) >= 1 {
		//Sum and GreaterThan
	}

	if len(t.GreaterThan) >= 1 && len(t.Quantity) >= 1 {
		//GreaterThan and Quantity
	}

	if len(t.GreaterThan) >= 1 && len(t.Quantity) >= 1 && len(t.Sum) >= 1 {
		//GreaterThan and Quantity and sum
	}
}

//there are 7 types of tasks, what are the ones that interact with eachother?

//touch, notouch
//Touch & NoTouch interact if the are the same, or if any color is the same
//Touch & Touch interacts if any color is the same
//NoTouch & NoTouch interacts if any color is the same

//bottom,top
// Bottom & Top interact if the same color is listed
// Bottom & Bottom interact if there arent any more places
// Top & Top interact if there arent any more places

//quantity, sum, gt
// Quantity & Sum interact if any color in sum is Quantity color
// Quantity & Quantity only interact if they are the same color
// Sum & Sum interact if any color is in both pairs
// Quantity & GreaterThan interact if any color is in Quantity
// GreaterThan & Sum interact if any color is in either
// GreaterThan & GreaterThan interact if any color is in both pairs

//AH, I realize that these are a priori interactions, I know that touches can impact each other,
//the other type of interactions are based on color itself; as that is the common thread, I really care as much,
//if not more about all the places where G has a rule for it

// there are a few types of rules, SHALL DO, like quantity, sum, and greater than, you MUST do these; the only way you meet these if you MUST place at least one color
// there are rules that are MAY do, meaning if you place the color criteria, it must conform to the rules, e.g. top, bottom, touch, no touch.
// top and bottom apply even if there is just one sphere of a color, -> they apply to the whole color
// touch and notouch only apply if there are more than 2 speheres placed: ACOLOR & BCOLOR

type Quantity struct {
	Number int
	Color  logic.Color
}

type Sum struct {
	Total  int
	Colors []logic.Color
}

type GreaterThan struct {
	GreaterColor logic.Color
	LesserColor  logic.Color
}

type Bottom struct {
	Color logic.Color
}

type Top struct {
	Color logic.Color
}

type Touch struct {
	Colors []logic.Color
}

type NoTouch struct {
	Colors []logic.Color
}

type TasksCollection struct {
	NoTouch     map[string]NoTouch
	Touch       map[string]Touch
	Top         map[string]Top
	Bottom      map[string]Bottom
	GreaterThan map[string]GreaterThan
	Sum         map[string]Sum
	Quantity    map[string]Quantity
}

func NewTasksCollection() (t *TasksCollection) {
	t = &TasksCollection{
		NoTouch:     make(map[string]NoTouch),
		Touch:       make(map[string]Touch),
		Top:         make(map[string]Top),
		Bottom:      make(map[string]Bottom),
		GreaterThan: make(map[string]GreaterThan),
		Sum:         make(map[string]Sum),
		Quantity:    make(map[string]Quantity),
	}
	return
}

func CategorizeTasks(tasks logic.Tasks) (t TasksCollection, err error) {

	t = *NewTasksCollection()

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
				return t, err
			}
			color := logic.NewColorShort(parts[2])

			t.Quantity[string(task)] = Quantity{
				Number: quantity,
				Color:  color,
			}

		case strings.Contains(string(task), "BOTTOM"):
			t.Bottom[string(task)] = Bottom{Color: logic.NewColorShort(parts[1])}
		case strings.Contains(string(task), "TOP"):
			t.Top[string(task)] = Top{Color: logic.NewColorShort(parts[1])}

		case strings.Contains(string(task), "TOUCH"):
			if strings.Contains(string(task), "NOTOUCH") {
				noTouch := NoTouch{make([]logic.Color, 0)}
				noTouch.Colors = append(noTouch.Colors, logic.NewColorShort(parts[1]))
				noTouch.Colors = append(noTouch.Colors, logic.NewColorShort(parts[2]))
				t.NoTouch[string(task)] = noTouch
			} else {
				touch := Touch{make([]logic.Color, 0)}
				touch.Colors = append(touch.Colors, logic.NewColorShort(parts[1]))
				touch.Colors = append(touch.Colors, logic.NewColorShort(parts[2]))
				t.Touch[string(task)] = touch
			}
		case strings.Contains(string(task), "SUM"):
			var colors []logic.Color
			colors = append(colors, logic.NewColorShort(parts[2]))
			colors = append(colors, logic.NewColorShort(parts[3]))
			t.Sum[string(task)] = Sum{
				Total:  4,
				Colors: colors,
			}

		case strings.Contains(string(task), "GT"):
			t.GreaterThan[string(task)] = GreaterThan{
				GreaterColor: logic.NewColorShort(parts[1]),
				LesserColor:  logic.NewColorShort(parts[2]),
			}

		default:
			err = fmt.Errorf("Could not parse task %s", task)
			if err != nil {
				return
			}
		}
	}

	return
}
