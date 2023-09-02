package storage

import (
	"github.com/nieuwsma/dimension/pkg/logic"
)

type GameProvider interface {
	GetGames() (games map[string]logic.Game, err error)
	GetGame(gameID string) (game logic.Game, err error)
	StoreGame(gameID string, game logic.Game) (err error)
	DeleteGame(gameID string) (err error)

	GetTrainingSessions() (ts map[string]logic.TrainingSession, err error)
	GetTrainingSession(trainID string) (ts logic.TrainingSession, err error)
	StoreTrainingSession(trainID string, session logic.TrainingSession) (err error)
	DeleteTrainingSession(trainID string) (err error)

	DeleteExpiredTrainingSessions() (err error)
}
