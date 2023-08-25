package domain

import (
	_ "embed"
	"encoding/json"
	"time"
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

//TODO END OF MOVE TO API

//TODO consider adding a game history and tracking players over time! #goldplatting
//type GameHistory []Game

type Game struct {
	Players        map[PlayerName]Player //name of player
	Rounds         map[RoundNumber]Round //tracks round ID to a play deck // TODO only 6 rounds are played
	Deck           Deck
	DrawSize       int
	HourglassLimit time.Duration
}

type PlayerName string
type RoundNumber int

type Player struct {
	PlayerName  PlayerName
	Points      int
	BonusTokens int //TODO bonus payout is: 0 = -6, 1 = -3, 2 = -1, 3 = 0, 4 = 1, 5 = 3, 6 = 6
}

type Round struct {
	PlayerTurns map[PlayerName]Turn //players play a dimension in a round; do i need to track round score? probably
	Tasks       Tasks               // here are the rules all players are playing by in the round; it is their constraints!
}

type Turn struct {
	Dimension      Dimension
	Score          int
	TaskViolations error
}

type Deck struct {
	DrawPile    Tasks
	DiscardPile Tasks
}

type Tasks []Task
type Task string

type DimensionGame interface {
	NewGame(players []Player) (g Game)
}
