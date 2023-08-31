package middleware

import (
	"dimension/internal/storage"
	"dimension/pkg/logic"
)

var GameProvider storage.GameProvider

// Game Management Routes

// Game Management Routes

func CreateGame() (gameID string, err error) {

	return
}

func DeleteGame(gameID string) (err error) {
	// Logic to delete a game
	return
}

func GetGameDetails(gameID string) (game logic.Game, err error) {

	return
}

func RemovePlayerFromGame(gameID string, playerID string) (err error) {
	// Logic to remove player from game
	return
}

func AddPlayerToGame(gameID string, name logic.PlayerName) (err error) {
	// Logic to add player to game

	return
}

// Rounds Routes

func ForceStartNewRound(gameID string) (err error) {
	// Logic to start a new round

	return
}

func GetRounds(gameID string) (err error) {
	// Logic to get rounds

	return
}

func GetSpecificRoundStatus(gameID string, roundID int) (err error) {
	// Logic to get specific round status

	return
}

func ForceRoundCompletion(gameID string) (err error) {
	// Logic to force round completion

	return
}

// Players Routes

func PlayerTakeTurn(gameID string, playerID string, roundID int, dimension logic.Dimension) (turn logic.Turn, err error) {
	// Logic for player to take a turn

	return
}

// Rules Route

func GetGameRules() (err error) {
	// Logic to get game rules

	return
}

// Training Routes

func StartTrainingSession(ommitedTypes logic.Tasks) (trainID string, tasks logic.Tasks, err error) {
	// Logic to start a training session

	//todo need to monkey around with the tasks
	//todo need to generate a new training ID
	trainID = "test"
	trainingSession := logic.NewTrainingSession(6, 12345)
	err = GameProvider.StoreTrainingSession(trainID, *trainingSession)

	tasks = trainingSession.Tasks

	return
}

func PlayTrainingSession(trainID string, dimension logic.Dimension) (trainingSession logic.TrainingSession, err error) {
	// Logic to play the training session
	trainingSession, err = GameProvider.GetTrainingSession(trainID)
	if err != nil {
		return
	}

	trainingSession.PlayTurn(dimension)

	err = GameProvider.StoreTrainingSession(trainID, trainingSession)

	return
}

func RetrieveTrainingStatus(trainID string) (trainingSession logic.TrainingSession, err error) {
	// Logic to retrieve training status
	trainingSession, err = GameProvider.GetTrainingSession(trainID)
	return
}

func RegenerateTrainingSession(trainID string) (tasks logic.Tasks, err error) {
	// Logic to retrieve training status
	trainingSession, err := GameProvider.GetTrainingSession(trainID)

	if err != nil {
		return
	}

	trainingSession.NextRound()

	err = GameProvider.StoreTrainingSession(trainID, trainingSession)
	if err != nil {
		return
	}

	tasks = trainingSession.Tasks
	return
}
