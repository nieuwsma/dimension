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

// this is the master deck!
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
type RoundNumber int

type Player struct {
	PlayerName  PlayerName
	Turns       []Turn
	Points      int
	BonusTokens int //TODO bonus payout is: 0 = -6, 1 = -3, 2 = -1, 3 = 0, 4 = 1, 5 = 3, 6 = 6
}

type Round struct {
	Tasks Tasks // here are the rules all players are playing by in the round; it is their constraints!
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
