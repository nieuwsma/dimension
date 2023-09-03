package main

import (
	"container/list"
	"errors"
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

	//colors := list.New()
	//
	//addRelation(logic.NewColorShort("G"), logic.NewColorShort("B"), colors)
	//addRelation(logic.NewColorShort("O"), logic.NewColorShort("W"), colors)
	//printList(colors) // GREEN -> BLUE -> EMPTY -> ORANGE -> WHITE -> EMPTY
	//
	//counts := colorCounts(colors)
	//for color, count := range counts {
	//	fmt.Printf("%s: %d\n", color.LongHand(), count)
	//}
	//
	//colors2 := list.New()
	//addRelation(logic.NewColorShort("G"), logic.NewColorShort("B"), colors2)
	//addRelation(logic.NewColorShort("B"), logic.NewColorShort("O"), colors2)
	//addRelation(logic.NewColorShort("O"), logic.NewColorShort("W"), colors2)
	//printList(colors2) // GREEN -> BLUE -> ORANGE -> WHITE -> EMPTY
	//
	//counts2 := colorCounts(colors2)
	//for color, count := range counts2 {
	//	fmt.Printf("%s: %d\n", color.LongHand(), count)
	//}

	for i := 0; i < 10; i++ {
		fmt.Println(fmt.Sprintf("TEST CASE %v", i))

		trainingSession := logic.NewTrainingSession(6, 12345)

		dimension, _ := logic.NewDimension()
		trainingSession.PlayTurn("autopilot", *dimension)

		tasksCollection, _ := CategorizeTasks(trainingSession.Tasks)

		colorMap := tasksCollection.GetColorDependencyTasks()
		rq := tasksCollection.GetRequiredQuantities()
		gt := tasksCollection.GetRequiredGreaterThanLessThan()
		su := tasksCollection.GetRequiredSums()
		at := tasksCollection.GetAllowedTouches()
		rt := tasksCollection.GetRequiredTouches()
		fmt.Println(fmt.Sprintf("tasks: %v", trainingSession.Tasks))
		fmt.Println(fmt.Sprintf("task collection \n%v", tasksCollection.String()))
		fmt.Println(fmt.Sprintf("colorMap : %v", colorMap))
		fmt.Println(fmt.Sprintf("required touch: %v \nallowed touch: %v \nrequired quantities: %v \nA > B: %v \nSums : %v", rt, at, rq, gt, su))
		fmt.Println()
	}
}

func (t *TasksCollection) String() string {
	var s string
	s = fmt.Sprintf("NoTouch: %v\nTouch: %v\nTop: %v\nBottom: %v\nGreaterThan: %v\nSum: %v\nQuantity: %v", t.NoTouch, t.Touch, t.Top, t.Bottom, t.GreaterThan, t.Sum, t.Quantity)
	return s
}

//A couple of ways to think about this... by color; a map of colors to tasks that they are involved in
// by task; first order connections (which is based on color)
// by task; second+ order connections (basically a connection of all tasks that are related to the nth degree)

func (t *TasksCollection) GetColorDependencyTasks() (colors map[logic.Color][]string) {
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
		for _, c := range v.Colors {
			colors[c] = append(colors[c], task)
		}
	}
	for task, v := range t.Touch {
		for _, c := range v.Colors {
			colors[c] = append(colors[c], task)
		}
	}
	for task, v := range t.NoTouch {
		for _, c := range v.Colors {
			colors[c] = append(colors[c], task)
		}
	}

	for color, tasks := range colors {
		colors[color] = deduplicateTasks(tasks)
	}
	return
}

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

// no color relationship are required unless in TOUCH. always search for affirmative connections
func (t *TasksCollection) GetRequiredTouches() (requiredColorTouches map[logic.Color][]logic.Color) {
	requiredColorTouches = make(map[logic.Color][]logic.Color)
	for _, v := range t.Touch {
		requiredColorTouches[v.Colors[0]] = append(requiredColorTouches[v.Colors[0]], v.Colors[1])
		requiredColorTouches[v.Colors[1]] = append(requiredColorTouches[v.Colors[1]], v.Colors[0])
	}

	for k, v := range requiredColorTouches {
		requiredColorTouches[k] = deduplicateColors(v)
	}
	return
}

//todo this doesnt preserve the relationships well enough
//func (t *TasksCollection) GetRequiredSums() (colors map[logic.Color]int) {
//	colorsList := list.New()
//
//	for _, v := range t.Sum {
//
//		err := addRelation(v.Colors[0], v.Colors[1], colorsList)
//		if err != nil {
//			fmt.Println("Error:", err)
//		}
//	}
//	colors = accountForColor(colorsList)
//
//	return
//}

func (t *TasksCollection) GetRequiredSums() (sums map[logic.Color][]logic.Color) {
	sums = make(map[logic.Color][]logic.Color)

	for _, v := range t.Sum {
		sums[v.Colors[0]] = append(sums[v.Colors[0]], v.Colors[1])
		sums[v.Colors[1]] = append(sums[v.Colors[1]], v.Colors[0])
	}
	for k, v := range sums {
		sums[k] = deduplicateColors(v)
	}
	return
}

// all possible color interactions are allowed, unless explicitly in NOTOUCH. always search for affirmative connections
func (t *TasksCollection) GetAllowedTouches() (allowedColorTouches map[logic.Color][]logic.Color) {
	allowedColorTouches = GetAllTouches()

	for _, v := range t.NoTouch {
		colorA := v.Colors[0]
		colorB := v.Colors[1]

		allowedColorTouches[colorA] = removeColors(allowedColorTouches[colorA], colorB)
		allowedColorTouches[colorB] = removeColors(allowedColorTouches[colorB], colorA)
	}
	return
}

func removeColors(slice []logic.Color, s logic.Color) []logic.Color {
	var result []logic.Color
	for _, item := range slice {
		if item != s {
			result = append(result, item)
		}
	}
	return result
}

func GetAllTouches() (allColorTouches map[logic.Color][]logic.Color) {
	allColorTouches = make(map[logic.Color][]logic.Color)

	for i := logic.Green; i <= logic.Black; i++ {
		for j := logic.Green; j <= logic.Black; j++ {
			allColorTouches[i] = append(allColorTouches[i], j)
		}
	}
	return
}

func accountForColor(l *list.List) map[logic.Color]int {
	colorMap := make(map[logic.Color]int)

	for e := l.Back(); e != nil; e = e.Prev() {
		color := e.Value.(logic.Color)
		if color.Equals(logic.Empty) {
			continue
		}

		colorMap[color] = 1
	}

	return colorMap
}

func colorCounts(l *list.List) map[logic.Color]int {
	colorMap := make(map[logic.Color]int)
	counter := 0

	for e := l.Back(); e != nil; e = e.Prev() {
		color := e.Value.(logic.Color)
		if color.Equals(logic.Empty) {
			counter = 0
			continue
		}

		colorMap[color] = counter
		counter++
	}

	return colorMap
}

func printList(l *list.List) {
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Print(e.Value.(logic.Color).LongHand()) // Use LongHand for printing
		if e.Next() != nil {
			fmt.Print(" -> ")
		}
	}
	fmt.Println()
}

func findElement(color logic.Color, l *list.List) *list.Element {
	for e := l.Front(); e != nil; e = e.Next() {
		if e.Value.(logic.Color).Equals(color) {
			return e
		}
	}
	return nil
}

func ensureEmptyAtEnd(l *list.List) {
	if !l.Back().Value.(logic.Color).Equals(logic.Empty) {
		l.PushBack(logic.Empty)
	}
}

func addRelation(ancestor logic.Color, descendant logic.Color, l *list.List) error {
	ancestorElement := findElement(ancestor, l)
	descendantElement := findElement(descendant, l)

	if ancestorElement != nil && descendantElement != nil {
		if !ancestorElement.Next().Value.(logic.Color).Equals(descendant) {
			return errors.New("conflicting rule")
		} else {
			if ancestorElement.Next().Value.(logic.Color).Equals(logic.Empty) {
				l.Remove(ancestorElement.Next())
			}
			ensureEmptyAtEnd(l)
			return nil
		}
	}

	if ancestorElement == nil && descendantElement == nil {
		l.PushBack(ancestor)
		l.PushBack(descendant)
	} else if ancestorElement != nil {
		l.InsertAfter(descendant, ancestorElement)
		if ancestorElement.Next().Value.(logic.Color).Equals(logic.Empty) {
			l.Remove(ancestorElement.Next())
		}
	} else if descendantElement != nil {
		l.InsertBefore(ancestor, descendantElement)
		if descendantElement.Prev().Value.(logic.Color).Equals(logic.Empty) {
			l.Remove(descendantElement.Prev())
		}
	}

	ensureEmptyAtEnd(l)
	return nil
}

func deduplicateTasks(strings []string) []string {
	seen := make(map[string]bool)
	result := []string{}

	for _, str := range strings {
		if _, exists := seen[str]; !exists {
			seen[str] = true
			result = append(result, str)
		}
	}

	return result
}

func deduplicateColors(colors []logic.Color) []logic.Color {
	seen := make(map[logic.Color]bool)
	result := []logic.Color{}

	for _, str := range colors {
		if _, exists := seen[str]; !exists {
			seen[str] = true
			result = append(result, str)
		}
	}

	return result
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
	return
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
