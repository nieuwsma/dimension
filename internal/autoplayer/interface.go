package autoplayer

import "github.com/nieuwsma/dimension/pkg/logic"

type AutoPlayer interface {
	Solve(tasks logic.Tasks) (solution logic.Dimension)
}

func GetDefaultColors() (colors map[logic.Color]int) {
	colors = make(map[logic.Color]int)
	colors[logic.Green] = 3
	colors[logic.Blue] = 3
	colors[logic.Black] = 3
	colors[logic.White] = 3
	colors[logic.Orange] = 3
	return
}
