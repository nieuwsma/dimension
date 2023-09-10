package autoplayer

import (
	"github.com/nieuwsma/dimension/internal/tasks"
	"github.com/nieuwsma/dimension/pkg/logic"
)

func generateSpherePairs(colors []logic.Color) (pairs []logic.SpherePair) {
	const alphabet = "abcdefghjkn"
	for idx, color := range colors {
		char := alphabet[idx%len(alphabet)]
		pair := logic.NewSpherePair(logic.SphereID(string(char)), color)
		pairs = append(pairs, *pair)
	}
	return pairs
}

// GetTasklessColors - returns the colors that are not in any task, and therefore can be used freely
func GetTasklessColors(availableColors map[logic.Color]int, tc tasks.TasksCollection) (usableColors map[logic.Color]int) {
	usableColors = make(map[logic.Color]int)

	for color, count := range availableColors {
		if _, exists := tc.ColorTaskDependency[color]; !exists {
			usableColors[color] = count
		}
	}
	return
}

func ConvertUseableColorsMapToSlice(usableColors map[logic.Color]int) (selectableColors []logic.Color) {
	for color, count := range usableColors {
		for i := 0; i < count; i++ {
			selectableColors = append(selectableColors, color)
		}
	}
	return
}
