package tasks

import (
	"container/list"
	"fmt"
	"github.com/nieuwsma/dimension/pkg/logic"
	"strconv"
	"strings"
)

type Sums struct {
	Color  logic.Color
	Counts Counts
	Chain  []logic.Color
}

type Counts struct {
	Max int
	Min int
}

type Quantity struct {
	Number int
	Color  logic.Color
}

type ColorPair struct {
	A logic.Color
	B logic.Color
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
	Colors ColorPair
}

type NoTouch struct {
	Colors ColorPair
}

type TasksCollection struct {
	NoTouch     map[string]NoTouch
	Touch       map[string]Touch
	Top         map[string]Top
	Bottom      map[string]Bottom
	GreaterThan map[string]GreaterThan
	Sum         map[string]ColorPair
	Quantity    map[string]Quantity
	tasks       logic.Tasks
}

func NewTasksCollection(tasks logic.Tasks) (t *TasksCollection, err error) {
	t = &TasksCollection{
		NoTouch:     make(map[string]NoTouch),
		Touch:       make(map[string]Touch),
		Top:         make(map[string]Top),
		Bottom:      make(map[string]Bottom),
		GreaterThan: make(map[string]GreaterThan),
		Sum:         make(map[string]ColorPair),
		Quantity:    make(map[string]Quantity),
		tasks:       tasks,
	}

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
				noTouch := NoTouch{Colors: ColorPair{
					A: logic.NewColorShort(parts[1]),
					B: logic.NewColorShort(parts[2]),
				}}

				t.NoTouch[string(task)] = noTouch
			} else {
				touch := Touch{Colors: ColorPair{
					A: logic.NewColorShort(parts[1]),
					B: logic.NewColorShort(parts[2]),
				}}
				t.Touch[string(task)] = touch
			}
		case strings.Contains(string(task), "SUM"):
			t.Sum[string(task)] = ColorPair{
				A: logic.NewColorShort(parts[2]),
				B: logic.NewColorShort(parts[3]),
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

func (t *TasksCollection) AllowedCounts() (counts map[logic.Color]Counts) {

	counts = make(map[logic.Color]Counts)
	//COUNTS
	if len(t.GreaterThan) >= 1 && len(t.Quantity) >= 1 && len(t.Sum) >= 1 {
		//GreaterThan and Quantity and sum
	} else if len(t.GreaterThan) >= 1 && len(t.Quantity) >= 1 {
		//GreaterThan and Quantity
	} else if len(t.Sum) >= 1 && len(t.GreaterThan) >= 1 {
		//Sum and GreaterThan
	} else if len(t.Quantity) >= 1 && len(t.Sum) >= 1 {
		//Quantity and Sum
		requiredQuantities := t.GetRequiredQuantities()
		requiredSums := t.GetRequiredSums()

		for color, summedColors := range requiredSums {
			//if the color is already in the map; then we need to check if its peers are,
			//if the peers are in, the math is harder, if the peers are NOT in the list then do the simple math.
			if value, exists := requiredQuantities[color]; exists {
				for _, summedColor := range summedColors.Chain {
					if _, exist := counts[summedColor]; !exist {
						counts[summedColor] = Counts{
							Max: 4 - value,
							Min: 4 - value,
						}
					}
				}
			} else {
				//color is NOT in quantities
				for _, summedColor := range summedColors.Chain {
					if _, exist := counts[summedColor]; !exist {
						counts[summedColor] = Counts{
							Max: 4 - value,
							Min: 4 - value,
						}
					}
				}
			}

		}

	} else if len(t.GreaterThan) > 1 {
		//Multiple GreaterThan

	} else if len(t.Sum) > 1 {
		//multiple Sum
		sums := t.GetRequiredSums()
		//todo handle if len(sums) > 2; that means there is a 3,4,or 5 way split
		//todo need to make sure i cant violate the count totals
		for color, _ := range sums {
			counts[color] = Counts{
				Max: 3,
				Min: 1,
			}
		}
	} else if len(t.Quantity) > 1 {
		//Multiple Quantity
		q := t.GetRequiredQuantities()
		for color, number := range q {
			counts[color] = Counts{
				Max: number,
				Min: number,
			}
		}
	}

	return
}

func (t *TasksCollection) String() string {
	var s string
	s = fmt.Sprintf("NoTouch: %v\nTouch: %v\nTop: %v\nBottom: %v\nGreaterThan: %v\nSum: %v\nQuantity: %v", t.NoTouch, t.Touch, t.Top, t.Bottom, t.GreaterThan, t.Sum, t.Quantity)
	return s
}

// A couple of ways to think about this... by color; a map of colors to tasks that they are involved in
// by task; first order connections (which is based on color)
// by task; second+ order connections (basically a connection of all tasks that are related to the nth degree)
func (t *TasksCollection) GetColorTasksDependency() (colors map[logic.Color][]string) {
	colors = make(map[logic.Color][]string)

	for task, v := range t.Top {
		colors[v.Color] = append(colors[v.Color], task)
	}
	for task, v := range t.Bottom {
		colors[v.Color] = append(colors[v.Color], task)
	}
	for task, v := range t.Quantity {
		colors[v.Color] = append(colors[v.Color], task)
	}
	for task, v := range t.GreaterThan {
		colors[v.GreaterColor] = append(colors[v.GreaterColor], task)
		colors[v.LesserColor] = append(colors[v.LesserColor], task)

	}
	for task, v := range t.Sum {

		colors[v.A] = append(colors[v.A], task)
		colors[v.B] = append(colors[v.B], task)

	}
	for task, v := range t.Touch {
		colors[v.Colors.A] = append(colors[v.Colors.A], task)
		colors[v.Colors.B] = append(colors[v.Colors.B], task)

	}
	for task, v := range t.NoTouch {
		colors[v.Colors.A] = append(colors[v.Colors.A], task)
		colors[v.Colors.B] = append(colors[v.Colors.B], task)
	}

	for color, tasks := range colors {
		colors[color] = deduplicateTasks(tasks)
	}
	return
}

func (t *TasksCollection) KnownRelations() (relations map[string][]string) {
	relations = make(map[string][]string)
	colorTaskMap := t.GetColorTasksDependency()

	for _, task := range colorTaskMap {
		for _, subtask := range task {
			relation := relations[subtask]
			relation = append(relation, task...)
			relations[subtask] = relation
		}
	}

	for task, tasks := range relations {
		relations[task] = removeTask(deduplicateTasks(tasks), task)

	}

	return
}

//func (t *TasksCollection) GetTasksDependency() (bar map[string][]string, foo [][]string) {
//	bar = make(map[string][]string)
//
//	gctd := t.GetColorTasksDependency()
//
//	for _, tasks := range gctd {
//		for _, task := range tasks {
//			retrievedTasks, exist := bar[task]
//			if exist {
//
//			} else {
//
//			}
//		}
//	}
//	for task, v := range t.Top {
//		bar[v.Color] = append(bar[v.Color], task)
//	}
//	for task, v := range t.Bottom {
//		bar[v.Color] = append(bar[v.Color], task)
//	}
//	for task, v := range t.Quantity {
//		bar[v.Color] = append(bar[v.Color], task)
//	}
//	for task, v := range t.GreaterThan {
//		bar[v.GreaterColor] = append(bar[v.GreaterColor], task)
//		bar[v.LesserColor] = append(bar[v.LesserColor], task)
//
//	}
//	for task, v := range t.Sum {
//
//		bar[v.A] = append(bar[v.A], task)
//		bar[v.B] = append(bar[v.B], task)
//
//	}
//	for task, v := range t.Touch {
//		bar[v.Colors.A] = append(bar[v.Colors.A], task)
//		bar[v.Colors.B] = append(bar[v.Colors.B], task)
//
//	}
//	for task, v := range t.NoTouch {
//		bar[v.Colors.A] = append(bar[v.Colors.A], task)
//		bar[v.Colors.B] = append(bar[v.Colors.B], task)
//	}
//
//	for color, tasks := range bar {
//		bar[color] = deduplicateTasks(tasks)
//	}
//	return
//}

func (t *TasksCollection) GetRequiredQuantities() (colors map[logic.Color]int) {
	colors = make(map[logic.Color]int)
	for _, v := range t.Quantity {
		colors[v.Color] = v.Number
	}

	return
}

func (t *TasksCollection) GetRequiredGreaterThanLessThan() (colors map[logic.Color]int) {
	colorsList := list.New()

	for _, v := range t.GreaterThan {
		err := addRelation(v.GreaterColor, v.LesserColor, colorsList)
		if err != nil {
			fmt.Println("Error:", err)
		}
	}
	colors = colorCounts(colorsList)

	return
}

func (t *TasksCollection) GetRequiredSums() (sums map[logic.Color]Sums) {
	sums = make(map[logic.Color]Sums)

	for _, v := range t.Sum {
		SUM, exists := sums[v.A]
		if exists {
			SUM.Counts = Counts{
				Max: 3,
				Min: 1,
			}
			SUM.Chain = append(SUM.Chain, v.B)
		} else {

			SUM = Sums{
				Color:  v.A,
				Counts: Counts{Min: 1, Max: 3},
				Chain:  []logic.Color{v.B},
			}
		}
		sums[v.A] = SUM
	}
	for k, v := range sums {
		v.Chain = deduplicateColors(v.Chain)
		sums[k] = v
	}
	return
}

func (t *TasksCollection) GetTop() (colors []logic.Color) {
	for _, v := range t.Top {
		colors = append(colors, v.Color)
	}
	return
}

func (t *TasksCollection) GetBottom() (colors []logic.Color) {
	for _, v := range t.Bottom {
		colors = append(colors, v.Color)
	}
	return
}

// all possible color interactions are allowed, unless explicitly in NOTOUCH. always search for affirmative connections
func (t *TasksCollection) GetAllowedTouches() (allowedColorTouches map[logic.Color][]logic.Color) {
	allowedColorTouches = GetAllTouches()

	for _, v := range t.NoTouch {

		allowedColorTouches[v.Colors.A] = removeColors(allowedColorTouches[v.Colors.A], v.Colors.B)
		allowedColorTouches[v.Colors.B] = removeColors(allowedColorTouches[v.Colors.B], v.Colors.A)
	}
	return
}

// no color relationship are required unless in TOUCH. always search for affirmative connections
func (t *TasksCollection) GetRequiredTouches() (requiredColorTouches map[logic.Color][]logic.Color) {
	requiredColorTouches = make(map[logic.Color][]logic.Color)
	for _, v := range t.Touch {
		requiredColorTouches[v.Colors.A] = append(requiredColorTouches[v.Colors.A], v.Colors.B)
		requiredColorTouches[v.Colors.B] = append(requiredColorTouches[v.Colors.B], v.Colors.A)
	}

	for k, v := range requiredColorTouches {
		requiredColorTouches[k] = deduplicateColors(v)
	}
	return
}
