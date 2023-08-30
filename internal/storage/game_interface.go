package storage

import (
	"dimension/pkg/logic"
)

//TODO consider adding a game history and tracking players over time! #goldplatting
//type GameHistory []Game

type GameProvider interface {
	GetGames() (games map[string]logic.Game, err error)
	GetGame(gameID string) (game logic.Game, err error)
	StoreGame(gameID string, game logic.Game) (err error)
	DeleteGame(gameID string) (err error)

	GetTrainingSession(trainID string) (ts logic.TrainingSession, err error)
	StoreTrainingSession(trainID string, session logic.TrainingSession) (err error)
	DeleteTrainingSession(trainID string) (err error)

	//AddPlayer(gameID string, name logic.PlayerName) (err error)
	//RemovePlayer(gameID string, name logic.PlayerName) (err error)
	//
	//NextRound(gameID string) (activeTasks logic.Tasks, err error)
	//GetCurrentRound(gameID string) (roundCount int, round logic.Round, err error)
	//
	//EndRound(gameID string, force bool) (err error)
	//GetLeaderboard(gameID string) (leaderboard logic.Leaderboard)
	//
	//PlayTurn(gameID string, playerName logic.PlayerName, dim logic.Dimension) (turn logic.Turn, err error)
	//EndGame(force bool) (err error)
	//
	//CreateTrainingSession(ts logic.TrainingSession) (err error)
	//PlayTrainingSession(dim logic.Dimension)
	//RegenerateTrainingSession()
}
