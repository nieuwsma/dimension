package main

import (
	"fmt"
	"github.com/nieuwsma/dimension/internal/autoplayer"
	"github.com/nieuwsma/dimension/pkg/logic"
)

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

	maxScore := 0
	for i := 0; i < 1000; i++ {

		trainingSession := logic.NewTrainingSession(6, 12345)

		var ai autoplayer.Mk0
		dimension := ai.Solve(trainingSession.Tasks)
		turn := trainingSession.PlayTurn("Mk0-Autoplayer", dimension)

		if turn.Score > maxScore {
			maxScore = turn.Score

			fmt.Println(fmt.Sprintf("TEST CASE %v", i))
			fmt.Println(ai.TaskCollection.String())
			fmt.Println(turn)

		}
	}
}
