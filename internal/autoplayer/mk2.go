package autoplayer

import (
	"github.com/nieuwsma/dimension/internal/tasks"
	"github.com/nieuwsma/dimension/pkg/logic"
)

// Mk2 is very simple, it sees if there are any color without tasks, then plays them; it will also assign minimum quantities
type Mk2 struct {
	TaskCollection tasks.TasksCollection
}

func (b *Mk2) Name() string {
	return "Mk2-Autoplayer"
}

func (b *Mk2) Solve(submittedTasks logic.Tasks) (solution logic.Dimension) {
	TaskCollection, _ := tasks.NewTasksCollection(submittedTasks)
	b.TaskCollection = *TaskCollection

	tasklessColors := GetTasklessColors(GetDefaultColors(), *TaskCollection)
	//todo get rid of this
	selectableColors := ConvertUseableColorsMapToSlice(tasklessColors)
	//
	//for color, count := range TaskCollection.RequiredQuantity {
	//	for i := 0; i < count; i++ {
	//		selectableColors = append(selectableColors, color)
	//	}
	//}

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

	for color, specialAvailability := range tasklessColors {
		count := availableQuantities[color]
		count.TasklessColorsAvailability = specialAvailability
		availableQuantities[color] = count
	}

	//todo need to preserve the linked list from the greater than less than chain
	spherePairs := generateSpherePairs(selectableColors)
	a, _ := logic.NewDimension(spherePairs...)
	solution = *a

	return
}

type Counts struct {
	DefaultMaximum             int          //The default for every color is 3
	RequiredQuantity           int          //if there is a quantity set, this is it
	GtLt                       tasks.Counts //The range of counts it can be if its engaged in GTLT
	TasklessColorsAvailability int          //The maximum allowed if the color is Taskless
}
