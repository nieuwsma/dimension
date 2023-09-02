package middleware

import (
	"dimension/internal/storage"
	"dimension/pkg/logic"
)

var GameProvider storage.GameProvider

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
