package autoplayer

import (
	"github.com/nieuwsma/dimension/internal/tasks"
	"github.com/nieuwsma/dimension/pkg/logic"
)

// Mk1 is very simple, it sees if there are any color without tasks, then plays them; it will also assign minimum quantities
type Mk1 struct {
	TaskCollection  tasks.TasksCollection
	availableColors map[logic.Color]int
}

func (b *Mk1) Solve(submittedTasks logic.Tasks) (solution logic.Dimension) {
	TaskCollection, _ := tasks.NewTasksCollection(submittedTasks)
	b.TaskCollection = *TaskCollection
	b.availableColors = GetDefaultColors()
	useableColors := make(map[logic.Color]int)

	var selectableColors []logic.Color

	for color, count := range b.availableColors {
		if _, exists := b.TaskCollection.ColorTaskDependency[color]; !exists {
			useableColors[color] = count
		}
	}

	for color, count := range useableColors {
		for i := 0; i < count; i++ {
			selectableColors = append(selectableColors, color)
		}
	}

	for color, count := range TaskCollection.RequiredQuantity {
		for i := 0; i < count; i++ {
			selectableColors = append(selectableColors, color)
		}
	}

	spherePairs := generateSpherePairs(selectableColors)
	a, _ := logic.NewDimension(spherePairs...)
	solution = *a

	return
}
