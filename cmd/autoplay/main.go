package main

import (
	"fmt"
	"github.com/nieuwsma/dimension/internal/autoplayer"
	"github.com/nieuwsma/dimension/pkg/logic"
	"os"
	"strconv"
)

func main() {

	trainingSession := logic.NewTrainingSession(6, 12345)
	var TurnStatistics = make(map[string][]TurnStatistic)
	rounds := 100000

	setRounds := os.Getenv("ROUNDS")
	if setRounds != "" {
		if intRounds, err := strconv.Atoi(setRounds); err == nil {
			rounds = intRounds
		}
	}
	for i := 0; i < rounds; i++ {

		{
			var mk0 autoplayer.Mk0
			dimension := mk0.Solve(trainingSession.Tasks)
			turn := trainingSession.PlayTurn("Mk0-Autoplayer", dimension)
			TurnStatistics[string(turn.PlayerName)] = append(TurnStatistics[string(turn.PlayerName)], TurnStatistic{
				Score:           turn.Score,
				Bonus:           turn.Bonus,
				TaskViolations:  len(turn.TaskViolations),
				DimensionLength: len(dimension.Dimension),
			})
		}
		{
			var mk1 autoplayer.Mk1
			dimension := mk1.Solve(trainingSession.Tasks)
			turn := trainingSession.PlayTurn("Mk1-Autoplayer", dimension)
			TurnStatistics[string(turn.PlayerName)] = append(TurnStatistics[string(turn.PlayerName)], TurnStatistic{
				Score:           turn.Score,
				Bonus:           turn.Bonus,
				TaskViolations:  len(turn.TaskViolations),
				DimensionLength: len(dimension.Dimension),
			})
		}

		trainingSession.Regenerate()
	}
	stats := generateStatistics(TurnStatistics)
	fmt.Println(fmt.Sprintf("Rounds: %v", rounds))
	for player, s := range stats {
		fmt.Println(player, ":", s)
	}
}

type TurnStatistic struct {
	Score           int
	Bonus           bool
	TaskViolations  int
	DimensionLength int
}

func (s Statistic) String() string {
	return fmt.Sprintf("Avg Score: %.2f | Avg Dimension Length: %.2f | Avg Task Violations: %.2f | Total Bonuses: %d | Max Score: %d",
		s.AverageScore, s.AverageDimensionLength, s.AverageTaskViolations, s.TotalBonuses, s.MaxScore)
}

func generateStatistics(turnStats map[string][]TurnStatistic) map[string]Statistic {
	result := make(map[string]Statistic)
	for key, stats := range turnStats {
		var totalScore int
		var totalViolations int
		var totalBonuses int
		var maxScore int
		var totalDimensionLengths int

		for _, stat := range stats {
			totalScore += stat.Score
			if stat.Score > maxScore {
				maxScore = stat.Score
			}
			totalViolations += stat.TaskViolations
			totalDimensionLengths += stat.DimensionLength
			if stat.Bonus {
				totalBonuses++
			}
		}
		result[key] = Statistic{
			AverageScore:           float64(totalScore) / float64(len(stats)),
			AverageDimensionLength: float64(totalDimensionLengths) / float64(len(stats)),
			AverageTaskViolations:  float64(totalViolations) / float64(len(stats)),
			TotalBonuses:           totalBonuses,
			MaxScore:               maxScore,
		}
	}
	return result
}

type Statistic struct {
	AverageScore           float64
	AverageTaskViolations  float64
	TotalBonuses           int
	MaxScore               int
	AverageDimensionLength float64
}
