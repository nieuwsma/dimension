package tasks

import (
	"container/list"
	"errors"
	"fmt"
	"github.com/nieuwsma/dimension/pkg/logic"
)

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

func removeColors(slice []logic.Color, s logic.Color) []logic.Color {
	var result []logic.Color
	for _, item := range slice {
		if item != s {
			result = append(result, item)
		}
	}
	return result
}

func removeTask(slice []string, s string) []string {
	var result []string
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
