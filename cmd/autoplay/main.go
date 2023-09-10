package main

import (
	"fmt"
	"github.com/nieuwsma/dimension/internal/tasks"
	"github.com/nieuwsma/dimension/pkg/logic"
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

	for i := 0; i < 1000; i++ {

		trainingSession := logic.NewTrainingSession(6, 12345)

		dimension, _ := logic.NewDimension()
		trainingSession.PlayTurn("autopilot", *dimension)

		tasksCollection, _ := tasks.NewTasksCollection(trainingSession.Tasks)
		if len(tasksCollection.Tasks) == 5 {
			fmt.Println(fmt.Sprintf("TEST CASE %v", i))
			fmt.Println(tasksCollection)
		}
	}
}
