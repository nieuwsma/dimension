package autoplayer

import (
	"github.com/nieuwsma/dimension/internal/tasks"
	"github.com/nieuwsma/dimension/pkg/logic"
)

// Mk2 is more complex, it tries to figure out the acceptable list of quantities for each color,
type Mk2 struct {
	TaskCollection      tasks.TasksCollection
	ColorTaskTypeCounts map[logic.Color]TaskTypeCounts
	AvailableQuantities map[logic.Color]Counts
	SelectableColors    map[logic.Color]int
}

type TaskTypeCounts struct {
	Quantity int
	Sum      int
	GT       int
	Top      int
	Bottom   int
	Touch    int
	NoTouch  int
}

func (b *Mk2) Name() string {
	return "Mk2-Autoplayer"
}

func (b *Mk2) Solve(submittedTasks logic.Tasks) (solution logic.Dimension) {
	TaskCollection, _ := tasks.NewTasksCollection(submittedTasks)
	b.TaskCollection = *TaskCollection
	b.SelectableColors = make(map[logic.Color]int)

	b.ColorTaskTypeCounts = make(map[logic.Color]TaskTypeCounts)
	for color, taskTuples := range TaskCollection.ColorTaskDependency {
		for _, taskTuple := range taskTuples {
			record := b.ColorTaskTypeCounts[color]

			switch taskTuple.Type {
			case tasks.QUANTITY:
				record.Quantity++
			case tasks.SUM:
				record.Sum++
			case tasks.GT:
				record.GT++
			case tasks.TOP:
				record.Top++
			case tasks.BOTTOM:
				record.Bottom++
			case tasks.TOUCH:
				record.Touch++
			case tasks.NOTOUCH:
				record.NoTouch++

			}

			b.ColorTaskTypeCounts[color] = record
		}
	}

	tasklessColors := GetTasklessColors(GetDefaultColors(), *TaskCollection)
	allQuantities := GetDefaultColors()
	availableQuantities := make(map[logic.Color]Counts)

	for color, maximum := range allQuantities {
		count := availableQuantities[color]
		count.DefaultMaximum = maximum
		availableQuantities[color] = count
	}

	for color, counts := range TaskCollection.RequiredGreaterThanLessThan {
		count := availableQuantities[color]
		count.GtLt = counts
		availableQuantities[color] = count

	}

	for color, requiredQuantity := range TaskCollection.RequiredQuantity {
		count := availableQuantities[color]
		count.RequiredQuantity = requiredQuantity
		availableQuantities[color] = count
	}

	for color, requiredSum := range TaskCollection.RequiredSums {
		count := availableQuantities[color]
		count.Sum = requiredSum.Counts
		availableQuantities[color] = count
	}

	for color, specialAvailability := range tasklessColors {
		count := availableQuantities[color]
		count.TasklessColorsAvailability = specialAvailability
		availableQuantities[color] = count
	}

	b.AvailableQuantities = availableQuantities

	//todo now do something with data!!!
	b.processQuantityElement()
	//todo, now we need to figure out positioning
	selectedColorsSlice := ConvertUseableColorsMapToSlice(b.SelectableColors)
	spherePairs := generateSpherePairs(selectedColorsSlice)
	a, _ := logic.NewDimension(spherePairs...)
	solution = *a

	return
}

type Counts struct {
	DefaultMaximum             int          //The default for every color is 3
	RequiredQuantity           int          //if there is a quantity set, this is it
	GtLt                       tasks.Counts //The range of counts it can be if its engaged in GTLT
	TasklessColorsAvailability int          //The maximum allowed if the color is Taskless
	Sum                        tasks.Counts
}

// todo work on this
// first pass; set quantity; set max GT if its an option, set sum = 2 if its an option
func (t *Mk2) processQuantityElement() {

	var GTwQ = make(map[logic.Color]bool)
	var SUMwQ = make(map[logic.Color]bool)

	for color, counts := range t.AvailableQuantities {
		GTwQ[color] = false
		SUMwQ[color] = false
		if counts.GtLt.Zero() && counts.Sum.Zero() && counts.RequiredQuantity == 0 && counts.TasklessColorsAvailability == 0 {
			t.SelectableColors[color] = counts.DefaultMaximum
		} else if counts.GtLt.Zero() && counts.Sum.Zero() && counts.RequiredQuantity == 0 && counts.TasklessColorsAvailability != 0 { // a bit redundant
			t.SelectableColors[color] = counts.TasklessColorsAvailability
		} else if counts.RequiredQuantity > 0 && counts.GtLt.Zero() && counts.Sum.Zero() { //easiest case
			t.SelectableColors[color] = counts.RequiredQuantity
		} else if counts.RequiredQuantity > 0 && !counts.GtLt.Zero() { //needs followup
			t.SelectableColors[color] = counts.RequiredQuantity
			GTwQ[color] = true
		} else if counts.RequiredQuantity > 0 && !counts.Sum.Zero() { //needs followup
			t.SelectableColors[color] = counts.RequiredQuantity
			SUMwQ[color] = true
		} else if counts.RequiredQuantity > 0 && !counts.Sum.Zero() && !counts.GtLt.Zero() { //needs followup
			t.SelectableColors[color] = counts.RequiredQuantity
			SUMwQ[color] = true
			GTwQ[color] = true
			//} else if counts.RequiredQuantity == 0 && !counts.Sum.Zero() && !counts.GtLt.Zero() { //needs followup
			//	t.SelectableColors[color] = 2
			//	SUMwQ[color] = true
			//	GTwQ[color] = true
			//} else if counts.RequiredQuantity == 0 && !counts.Sum.Zero() && counts.GtLt.Zero() { //needs followup
			//	SUMwQ[color] = true
			//	t.SelectableColors[color] = counts.GtLt.Max
			//} else if counts.RequiredQuantity == 0 && counts.Sum.Zero() && !counts.GtLt.Zero() { //needs followup
			//	GTwQ[color] = true
			//	t.SelectableColors[color] = 2

		}
	}

	/*
		There are a few scenarios; when A has a Q; B !Q; A GT B || A SUM B
		A & B !Q; but A GT B || A SUM B

	*/
	//if TasklessColorsAvaialbity == DefaultMaximum (USE that number)
	// else if tasklessColorsAvailabily == 0 && GT &7 SUM && Quantity == 0 ;
	//	DefaultMaximum == 3
	//else if RequiredQuantity > 0; use that quantity
	//else if greaterthan max/min != 0; use max? depends on other colors?
	// else if sum max/min != 0; use max? depends on other colors

	//for _, t := range t.ColorTaskTypeCounts {

	////TOUCHES
	//if t.Touch >= 1 && t.NoTouch >= 1 {
	//	//NoTouch and Touch
	//} else if t.NoTouch > 1 {
	//	//multiple NoTouches
	//} else if t.Touch > 1 {
	//	//Multiple Touches
	//}
	//
	////SITS
	//if t.Top >= 1 && t.Bottom >= 1 {
	//	//Top and Bottom
	//} else if t.Top > 1 {
	//	//multiple Top
	//} else if t.Bottom > 1 {
	//	//Multiple Bottom
	//}

	//COUNTS
	//if t.GT >= 1 && t.Quantity >= 1 && t.Sum >= 1 {
	//	//GreaterThan and Quantity and sum
	//} else if t.GT >= 1 && t.Quantity >= 1 {
	//	//GreaterThan and Quantity
	//} else if t.Sum >= 1 && t.GT >= 1 {
	//	//Sum and GreaterThan
	//} else if t.Quantity >= 1 && t.Sum >= 1 {
	//	//Quantity and Sum
	//} else if t.GT > 1 {
	//	//Multiple GreaterThan
	//} else if t.Sum > 1 {
	//	//multiple Sum
	//} else if t.Quantity > 1 {
	//	//Multiple Quantity
	//}

	//}
	return
}
