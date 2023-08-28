package domain

import (
	"fmt"
	"testing"
)

func TestValidateGame(t *testing.T) {
	gameRecords := getGamesTestCases()

	newGames := make(map[int]Game)

	for gameID, game := range gameRecords {

		newGame := game.DeepCopy()
		newGames[gameID] = *newGame

		for roundID, round := range newGame.Rounds {
			for playerName, player := range newGame.Players {
				turn := player.Turns[roundID]
				score, bonus, err := ScoreTurn(round.Tasks, turn.Dimension)

				player.ScoreRecord.Points = 0
				player.ScoreRecord.BonusTokens = 0
				player.Turns[roundID] = turn
				newGame.Players[playerName] = player

				if score != turn.Score || bonus != turn.Bonus {
					printErr := fmt.Sprintf("FAIL: game: %v, round: %v, player: %v, expected score %v expected bonus %v, actual score %v, actual bonus %v. failed tasks %v", gameID, roundID, playerName, turn.Score, turn.Bonus, score, bonus, err)
					t.Errorf("failed test case %s", printErr)

				}
				//else {
				//	fmt.Println(fmt.Sprintf("PASS: game: %v, round: %v, player: %v, expected score %v expected bonus %v, actual score %v, actual bonus %v. failed tasks %v", gameID, roundID, playerName, turn.Score, turn.Bonus, score, bonus, err))
				//
				//}

				turn.Score = score
				turn.Bonus = bonus
				turn.TaskViolations = err

			}
		}

		for roundID, _ := range newGame.Rounds {
			for playerName, player := range newGame.Players {
				turn := player.Turns[roundID]
				player.ScoreRecord.Points += turn.Score
				if turn.Bonus {
					player.ScoreRecord.BonusTokens++
				}
				newGame.Players[playerName] = player
			}
		}

		for playerName, player := range newGame.Players {
			if !player.ScoreRecord.Equals(game.Players[playerName].ScoreRecord) {
				printErr := fmt.Sprintf("GameID: %v, player: %v, expected final score %v expected final bonus %v, actual final score %v, actual final bonus %v. ", gameID, playerName, player.ScoreRecord.Points, player.ScoreRecord.BonusTokens, game.Players[playerName].ScoreRecord.Points, game.Players[playerName].ScoreRecord.BonusTokens)
				t.Errorf("failed test case %s", printErr)
			}
		}
	}
	//lb := game.GetLeaderboard()

	//fmt.Printf("%+v\n", lb)

}
