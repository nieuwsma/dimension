package autoplayer

import (
	"github.com/nieuwsma/dimension/internal/tasks"
	"github.com/nieuwsma/dimension/pkg/logic"
)

// Mk1 is very simple, it sees if there are any color without tasks, then plays them; it will also assign minimum quantities
type Mk1 struct {
	TaskCollection tasks.TasksCollection
}

func (b *Mk1) Name() string {
	return "Mk1-Autoplayer"
}
func (b *Mk1) Solve(submittedTasks logic.Tasks) (solution logic.Dimension) {
	TaskCollection, _ := tasks.NewTasksCollection(submittedTasks)
	b.TaskCollection = *TaskCollection

	usableColors := GetTasklessColors(GetDefaultColors(), *TaskCollection)
	selectableColors := ConvertUseableColorsMapToSlice(usableColors)

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
