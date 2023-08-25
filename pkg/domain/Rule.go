package domain

import (
	_ "embed"
	"encoding/json"
)

// TODO move to API!!!!
type RulesArrayDataFormat []RulesDataFormat

type RulesDataFormat struct {
	Name        string
	Quantity    int
	Description string
}

// this is the master deck! //todo this needs to change because when I run main the relative path import is missing!
//
//go:embed rules.json
var RulesData []byte

func GetAllRules() (r RulesArrayDataFormat) {
	_ = json.Unmarshal(RulesData, &r)
	return
}

func (f *RulesArrayDataFormat) ToTasks() (tasks Tasks) {
	for _, rule := range *f {
		tasks = append(tasks, Task(rule.Name))
	}
	return
}

func init() {
	r := GetAllRules()
	masterTasks := r.ToTasks()
	DefaultTasks = masterTasks
}

// todo do I want different tasks? how does this differ from what i need to display in API?
var DefaultTasks Tasks

//TODO END OF MOVE TO API

//TODO consider adding a game history and tracking players over time! #goldplatting
//type GameHistory []Game

type PlayerName string

type Player struct {
	PlayerName  PlayerName
	Turns       map[int]Turn
	ScoreRecord ScoreRecord
}

type Round struct {
	Tasks    Tasks // here are the rules all players are playing by in the round; it is their constraints!
	Resolved bool
}

type Turn struct {
	Dimension      Dimension
	Score          int
	Bonus          bool
	TaskViolations error
}

//type DimensionGame interface {
//	NewGame() (g Game)
//	EndGame(g *Game)
//	AddPlayer()
//	RemovePlayer()
//	PlayTurn()
//	GetRound()
//}
