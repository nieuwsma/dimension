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

type Task struct {
	Name        string
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

//TODO END OF MOVE TO API

type Game struct {
	Players  map[PlayerName]Player //name of player
	Rounds   map[RoundNumber]Round //tracks round ID to a play deck
	Deck     Deck
	DrawSize int
}

type PlayerName string
type RoundNumber int

type Player struct {
	Points      int
	BonusTokens int
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
