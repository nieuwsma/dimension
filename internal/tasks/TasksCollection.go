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
	NoTouch                     map[string]NoTouch
	Touch                       map[string]Touch
	Top                         map[string]Top
	Bottom                      map[string]Bottom
	GreaterThan                 map[string]GreaterThan
	Sum                         map[string]ColorPair
	Quantity                    map[string]Quantity
	Tasks                       logic.Tasks
	RequiredQuantity            map[logic.Color]int
	RequiredGreaterThanLessThan map[logic.Color]int
	RequiredSums                map[logic.Color]Sums
	RequiredTop                 []logic.Color
	RequiredBottom              []logic.Color
	AllowedTouches              map[logic.Color][]logic.Color
	RequiredTouches             map[logic.Color][]logic.Color
	ColorTaskDependency         map[logic.Color][]string
	Relations                   [][]string
}

func NewTasksCollection(tasks logic.Tasks) (t *TasksCollection, err error) {
	t = &TasksCollection{
		NoTouch:                     make(map[string]NoTouch),
		Touch:                       make(map[string]Touch),
		Top:                         make(map[string]Top),
		Bottom:                      make(map[string]Bottom),
		GreaterThan:                 make(map[string]GreaterThan),
		Sum:                         make(map[string]ColorPair),
		Quantity:                    make(map[string]Quantity),
		Tasks:                       tasks,
		RequiredQuantity:            make(map[logic.Color]int),
		RequiredGreaterThanLessThan: make(map[logic.Color]int),
		RequiredSums:                make(map[logic.Color]Sums),
		//RequiredTop:                 make([]logic.Color),
		//RequiredBottom:              []logic.Color
		AllowedTouches:      make(map[logic.Color][]logic.Color),
		RequiredTouches:     make(map[logic.Color][]logic.Color),
		ColorTaskDependency: make(map[logic.Color][]string),
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

	t.Tasks = tasks

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

	t.RequiredQuantity = t.fillRequiredQuantities()
	t.RequiredSums = t.fillRequiredSums()
	t.RequiredGreaterThanLessThan = t.fillRequiredGreaterThanLessThan()
	t.AllowedTouches = t.fillAllowedTouches()
	t.RequiredTouches = t.fillRequiredTouches()
	t.RequiredBottom = t.fillBottom()
	t.RequiredTop = t.fillTop()
	t.Relations = t.mapRelationsRecursively()
	t.ColorTaskDependency = t.fillColorTasksDependency()
	return
}

func (t *TasksCollection) allowedCounts() (counts map[logic.Color]Counts) {

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
		requiredQuantities := t.fillRequiredQuantities()
		requiredSums := t.fillRequiredSums()

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
		sums := t.fillRequiredSums()
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
		q := t.fillRequiredQuantities()
		for color, number := range q {
			counts[color] = Counts{
				Max: number,
				Min: number,
			}
		}
	}

	return
}

// Assuming logic.Color's String method is implemented
// (I will not re-implement it here since you mentioned it exists)

func (t *TasksCollection) String() string {
	var sb strings.Builder

	// TaskRelations
	sb.WriteString("\nTasks:\n")
	for _, v := range t.Tasks {
		sb.WriteString(fmt.Sprintf("\t%v\n", v))
	}
	// NoTouch
	sb.WriteString("NoTouch:\n")
	for k, v := range t.NoTouch {
		sb.WriteString(fmt.Sprintf("\t%s: { A: %s, B: %s }\n", k, v.Colors.A, v.Colors.B))
	}

	// Touch
	sb.WriteString("\nTouch:\n")
	for k, v := range t.Touch {
		sb.WriteString(fmt.Sprintf("\t%s: { A: %s, B: %s }\n", k, v.Colors.A, v.Colors.B))
	}

	// Top
	sb.WriteString("\nTop:\n")
	for k, v := range t.Top {
		sb.WriteString(fmt.Sprintf("\t%s: { Color: %s }\n", k, v.Color))
	}

	// Bottom
	sb.WriteString("\nBottom:\n")
	for k, v := range t.Bottom {
		sb.WriteString(fmt.Sprintf("\t%s: { Color: %s }\n", k, v.Color))
	}

	// GreaterThan
	sb.WriteString("\nGreaterThan:\n")
	for k, v := range t.GreaterThan {
		sb.WriteString(fmt.Sprintf("\t%s: { GreaterColor: %s, LesserColor: %s }\n", k, v.GreaterColor, v.LesserColor))
	}

	// Sum
	sb.WriteString("\nSum:\n")
	for k, v := range t.Sum {
		sb.WriteString(fmt.Sprintf("\t%s: { A: %s, B: %s }\n", k, v.A, v.B))
	}

	// Quantity
	sb.WriteString("\nQuantity:\n")
	for k, v := range t.Quantity {
		sb.WriteString(fmt.Sprintf("\t%s: { Number: %d, Color: %s }\n", k, v.Number, v.Color))
	}

	// RequiredQuantity
	sb.WriteString("\nRequiredQuantity:\n")
	for k, v := range t.RequiredQuantity {
		sb.WriteString(fmt.Sprintf("\t%s: %d\n", k, v))
	}

	// RequiredGreaterThanLessThan
	sb.WriteString("\nRequiredGreaterThanLessThan:\n")
	for k, v := range t.RequiredGreaterThanLessThan {
		sb.WriteString(fmt.Sprintf("\t%s: %d\n", k, v))
	}

	// RequiredSums
	sb.WriteString("\nRequiredSums:\n")
	for k, v := range t.RequiredSums {
		sb.WriteString(fmt.Sprintf("\t%s: { Color: %s, Counts: { Max: %d, Min: %d }, Chain: %v }\n", k, v.Color, v.Counts.Max, v.Counts.Min, v.Chain))
	}

	// RequiredTop
	sb.WriteString("\nRequiredTop:\n")
	for _, color := range t.RequiredTop {
		sb.WriteString(fmt.Sprintf("\t%s\n", color))
	}

	// RequiredBottom
	sb.WriteString("\nRequiredBottom:\n")
	for _, color := range t.RequiredBottom {
		sb.WriteString(fmt.Sprintf("\t%s\n", color))
	}

	// AllowedTouches
	sb.WriteString("\nAllowedTouches:\n")
	for k, v := range t.AllowedTouches {
		sb.WriteString(fmt.Sprintf("\t%s: %v\n", k, v))
	}

	// RequiredTouches
	sb.WriteString("\nRequiredTouches:\n")
	for k, v := range t.RequiredTouches {
		sb.WriteString(fmt.Sprintf("\t%s: %v\n", k, v))
	}

	// Relations
	sb.WriteString("\nRelations:\n")
	for _, v := range t.Relations {

		sb.WriteString(fmt.Sprintf("\t["))

		for _, v1 := range v {
			sb.WriteString(fmt.Sprintf("%v, ", v1))
		}
		sb.WriteString(fmt.Sprintf("]\n"))

	}

	// ColorTaskDependency
	sb.WriteString("\nColorTaskDependency:\n")
	for k, v := range t.ColorTaskDependency {
		sb.WriteString(fmt.Sprintf("\t%s: %v\n", k, v))
	}

	return sb.String()
}

// A couple of ways to think about this... by color; a map of colors to Tasks that they are involved in
// by task; first order connections (which is based on color)
// by task; second+ order connections (basically a connection of all Tasks that are related to the nth degree)
func (t *TasksCollection) fillColorTasksDependency() (colors map[logic.Color][]string) {
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

func (t *TasksCollection) fillRequiredQuantities() (colors map[logic.Color]int) {
	colors = make(map[logic.Color]int)
	for _, v := range t.Quantity {
		colors[v.Color] = v.Number
	}

	return
}

func (t *TasksCollection) fillRequiredGreaterThanLessThan() (colors map[logic.Color]int) {
	colorsList := list.New()

	for _, v := range t.GreaterThan {
		err := addRelation(v.GreaterColor, v.LesserColor, colorsList)
		if err != nil {
			//fmt.Println("Error:", err)
		}
	}
	colors = colorCounts(colorsList)

	return
}

func (t *TasksCollection) fillRequiredSums() (sums map[logic.Color]Sums) {
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

		SUM2, exists := sums[v.B]
		if exists {
			SUM2.Counts = Counts{
				Max: 3,
				Min: 1,
			}
			SUM2.Chain = append(SUM2.Chain, v.A)
		} else {

			SUM2 = Sums{
				Color:  v.B,
				Counts: Counts{Min: 1, Max: 3},
				Chain:  []logic.Color{v.A},
			}
		}
		sums[v.B] = SUM2
	}
	for k, v := range sums {
		v.Chain = deduplicateColors(v.Chain)
		sums[k] = v
	}
	return
}

func (t *TasksCollection) fillTop() (colors []logic.Color) {
	for _, v := range t.Top {
		colors = append(colors, v.Color)
	}
	return
}

func (t *TasksCollection) fillBottom() (colors []logic.Color) {
	for _, v := range t.Bottom {
		colors = append(colors, v.Color)
	}
	return
}

// all possible color interactions are allowed, unless explicitly in NOTOUCH. always search for affirmative connections
func (t *TasksCollection) fillAllowedTouches() (allowedColorTouches map[logic.Color][]logic.Color) {
	allowedColorTouches = GetAllTouches()

	for _, v := range t.NoTouch {

		allowedColorTouches[v.Colors.A] = removeColors(allowedColorTouches[v.Colors.A], v.Colors.B)
		allowedColorTouches[v.Colors.B] = removeColors(allowedColorTouches[v.Colors.B], v.Colors.A)
	}
	return
}

// no color relationship are required unless in TOUCH. always search for affirmative connections
func (t *TasksCollection) fillRequiredTouches() (requiredColorTouches map[logic.Color][]logic.Color) {
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
