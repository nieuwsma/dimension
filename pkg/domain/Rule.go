package domain

import (
	_ "embed"
	"encoding/json"
	"math/rand"
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

type Game struct {
	Players        map[PlayerName]Player //name of player
	PlayerTurns    map[PlayerName]Turn   //players play a dimension in a round; do i need to track round score? probably
	Rounds         []Round               //tracks round ID to a play deck // TODO only 6 rounds are played
	RoundCounter   int
	Deck           Deck
	DrawSize       int
	HourglassLimit time.Duration
}

func NewGame(drawSize int, hourGlassLimit time.Duration, seed int64) (g *Game) {
	g = &Game{
		DrawSize:       drawSize,
		HourglassLimit: hourGlassLimit,
	}

	g.Deck = newDeck(seed)
	g.Deck.Shuffle()
	return
}

func newDeck(seed int64) (d Deck) {
	d.DrawPile = DefaultTasks
	d.Seed = seed
	return
}

func (d *Deck) Shuffle() {
	//return everything to the draw pile and shuffle
	d.DrawPile = append(d.DrawPile, d.DiscardPile...)
	d.DrawPile = append(d.DrawPile, d.Active...)
	d.DiscardPile = Tasks{}
	d.Active = Tasks{}

	rand.New(rand.NewSource(d.Seed))
	rand.Shuffle(len(d.DrawPile), func(i, j int) { d.DrawPile[i], d.DrawPile[j] = d.DrawPile[j], d.DrawPile[i] })
}

func (g *Game) NextRound() (err error) {
	return
}

func (g *Game) GetCurrentRoundTasks() {}

func (g *Game) PlayTurn(playerName PlayerName, dim Dimension) (turn Turn) {
	return
}

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

type Deck struct {
	DrawPile    Tasks
	DiscardPile Tasks
	Active      Tasks
	Seed        int64
}

type Tasks []Task
type Task string

//type DimensionGame interface {
//	NewGame() (g Game)
//	EndGame(g *Game)
//	AddPlayer()
//	RemovePlayer()
//	PlayTurn()
//	GetRound()
//}
