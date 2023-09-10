package autoplayer

import (
	"github.com/nieuwsma/dimension/internal/tasks"
	"github.com/nieuwsma/dimension/pkg/logic"
)

// Mk0 is very simple, it sees if there are any color without tasks, then plays them
type Mk0 struct {
	TaskCollection  tasks.TasksCollection
	availableColors map[logic.Color]int
}

func (b *Mk0) Name() string {
	return "Mk0-Autoplayer"
}

func (b *Mk0) Solve(submittedTasks logic.Tasks) (solution logic.Dimension) {
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

	spherePairs := generateSpherePairs(selectableColors)
	a, _ := logic.NewDimension(spherePairs...)
	solution = *a

	return
}
