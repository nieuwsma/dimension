package tasks

import (
	"container/list"
	"errors"
	"fmt"
	"github.com/nieuwsma/dimension/pkg/logic"
)

func removeColors(slice []logic.Color, s logic.Color) []logic.Color {
	var result []logic.Color
	for _, item := range slice {
		if item != s {
			result = append(result, item)
		}
	}
	return result
}

func removeTask(slice []TaskTuple, s TaskTuple) []TaskTuple {
	var result []TaskTuple
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

func colorCounts(l *list.List) map[logic.Color]Counts {
	colorMap := make(map[logic.Color]Counts)
	counter := 0

	for e := l.Back(); e != nil; e = e.Prev() {
		color := e.Value.(logic.Color)
		if color.Equals(logic.Empty) {
			counter = 0
			continue
		}

		counts := colorMap[color]
		counts.Min = counter
		colorMap[color] = counts
		counter++
	}

	counter = 3

	for e := l.Front(); e != nil; e = e.Next() {
		color := e.Value.(logic.Color)
		if color.Equals(logic.Empty) {
			counter = 3
			continue
		}

		counts := colorMap[color]
		counts.Max = counter
		colorMap[color] = counts
		counter--
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

func deduplicateTasks(tasks []TaskTuple) []TaskTuple {
	seen := make(map[TaskTuple]bool)
	result := []TaskTuple{}

	for _, task := range tasks {
		if _, exists := seen[task]; !exists {
			seen[task] = true
			result = append(result, task)
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

// todo work on this
func ProcessCategorizedTasks(t TasksCollection) (interactions map[string][]string) {
	//TOUCHES
	if len(t.Touch) >= 1 && len(t.NoTouch) >= 1 {
		//NoTouch and Touch
	} else if len(t.NoTouch) > 1 {
		//multiple NoTouches
	} else if len(t.Touch) > 1 {
		//Multiple Touches
	}

	//SITS
	if len(t.Top) >= 1 && len(t.Bottom) >= 1 {
		//Top and Bottom
	} else if len(t.Top) > 1 {
		//multiple Top
	} else if len(t.Bottom) > 1 {
		//Multiple Bottom
	}

	//COUNTS
	if len(t.GreaterThan) >= 1 && len(t.Quantity) >= 1 && len(t.Sum) >= 1 {
		//GreaterThan and Quantity and sum
	} else if len(t.GreaterThan) >= 1 && len(t.Quantity) >= 1 {
		//GreaterThan and Quantity
	} else if len(t.Sum) >= 1 && len(t.GreaterThan) >= 1 {
		//Sum and GreaterThan
	} else if len(t.Quantity) >= 1 && len(t.Sum) >= 1 {
		//Quantity and Sum
	} else if len(t.GreaterThan) > 1 {
		//Multiple GreaterThan
	} else if len(t.Sum) > 1 {
		//multiple Sum
	} else if len(t.Quantity) > 1 {
		//Multiple Quantity
	}
	return
}
