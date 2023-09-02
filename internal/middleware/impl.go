package middleware

import (
	"dimension/internal/storage"
	"dimension/pkg/geometry"
	"dimension/pkg/logic"
	"dimension/pkg/rules"
)

var GameProvider storage.GameProvider

// Rules Route

func GetGameRules() (rules.RuleSet, logic.Colors, geometry.Geometries) {
	// Logic to get game rules
	rules, _ := rules.GetRuleSet(rules.Default)
	colors := logic.GetColors()
	geometries := geometry.GetGeometry()

	return rules, colors, geometries
}

// Training Routes

func StartTrainingSession(ommitedTypes logic.Tasks) (trainID string, tasks logic.Tasks, err error) {
	// Logic to start a training session

	//todo need to monkey around with the tasks
	trainID, err = randomStringWithPrefix()
	if err != nil {
		return
	}
	trainingSession := logic.NewTrainingSession(6, 12345)
	err = GameProvider.StoreTrainingSession(trainID, *trainingSession)

	tasks = trainingSession.Tasks

	return
}

func PlayTrainingSession(trainID string, playerName string, dimension logic.Dimension) (trainingSession logic.TrainingSession, err error) {
	// Logic to play the training session
	trainingSession, err = GameProvider.GetTrainingSession(trainID)
	if err != nil {
		return
	}

	trainingSession.PlayTurn(logic.PlayerName(playerName), dimension)

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

	trainingSession.Regenerate()

	err = GameProvider.StoreTrainingSession(trainID, trainingSession)
	if err != nil {
		return
	}

	tasks = trainingSession.Tasks
	return
}

func RetrieveTrainingSessions() (trainingSessions map[string]logic.TrainingSession, err error) {
	trainingSessions, _ = GameProvider.GetTrainingSessions()
	return
}
