package api

import (
	"dimension/pkg/logic"
	"errors"
	"time"
)

// Game represents a game object
type Game struct {
	GameID string `json:"gameId"`
}

// Player represents a player object
type Player struct {
	Name string `json:"name"`
}

// RoundStatus represents the status of a round
type RoundStatus struct {
	Tasks    []string       `json:"tasks"`
	Players  []PlayerStatus `json:"players"`
	IsActive bool           `json:"isActive"`
}

// PlayerStatus represents the status of a player
type PlayerStatus struct {
	PlayerName string `json:"playerName"`
	PlayerID   string `json:"playerId"`
	TurnTaken  bool   `json:"turnTaken"`
}

// GameDetails represents details of a game
type GameDetails struct {
	Leaderboard []LeaderboardEntry `json:"leaderboard"`
	Rounds      []RoundSummary     `json:"rounds"`
}

// LeaderboardEntry represents an entry in the leaderboard
type LeaderboardEntry struct {
	PlayerName string `json:"playerName"`
	PlayerID   string `json:"playerId"`
	Score      int    `json:"score"`
}

// RoundSummary represents a summary of a round
type RoundSummary struct {
	RoundID        int `json:"roundId"`
	TasksCompleted int `json:"tasksCompleted"`
}

// ForceCompletion represents an action to force completion
type ForceCompletion struct {
	ForceComplete bool   `json:"forceComplete"`
	Reason        string `json:"reason"`
}

// PlayerCreated represents a newly created player
type PlayerCreated struct {
	PlayerID  string `json:"playerId"`
	AuthToken string `json:"authToken"`
}

// RoundCreated represents a newly created round
type RoundCreated struct {
	RoundID int `json:"roundId"`
}

// Dimension represents dimension properties
type Dimension struct {
	A string `json:"a"`
	B string `json:"b"`
	C string `json:"c"`
	D string `json:"d"`
	E string `json:"e"`
	F string `json:"f"`
	G string `json:"g"`
	H string `json:"h"`
	I string `json:"i"`
	J string `json:"j"`
	K string `json:"k"`
	L string `json:"l"`
	M string `json:"m"`
	N string `json:"n"`
}

type DimensionResponse struct {
	Dimension map[string]string
}

// Task represents a task
type Task struct {
	Name        string `json:"Name"`
	Quantity    string `json:"Quantity"`
	Description string `json:"Description"`
}

// RulesResponse represents a response for game rules
type RulesResponse struct {
	Tasks      []Task         `json:"Tasks"`
	Geometries []GeometryItem `json:"Geometries"`
	Colors     []Color        `json:"Colors"`
}

// GeometryItem represents a geometry item
type GeometryItem struct {
	PolarAngle       float64  `json:"polarAngle"`
	InclinationAngle float64  `json:"inclinationAngle"`
	RadialDistance   float64  `json:"radialDistance"`
	ID               string   `json:"id"`
	Neighbors        []string `json:"neighbors"`
}

// Color represents a color
type Color struct {
	Name string `json:"Name"`
	Code string `json:"Code"`
}

type PostTrainingSessionRequest struct {
	types []string `json:"taskTypes"`
}

type PostTrainingSessionResponse struct {
	TrainID string      `json:"trainID"`
	Tasks   logic.Tasks `json:"tasks"`
}

type PostRegenerateTrainingSessionResponse struct {
	Tasks logic.Tasks `json:"tasks"`
}

type GetTrainingSessionResponse struct {
	Score              int               `json:"score"`
	BonusPoints        bool              `json:"bonusPoints"`
	SubmittedDimension DimensionResponse `json:"submittedDimension"`
	Tasks              logic.Tasks       `json:"tasks"`
	TaskViolations     []string          `json:"taskViolations"`
	ExpirationTime     CustomTime        `json:"expirationTime"`
}

type CustomTime struct {
	time.Time
}

const ctLayout = "2006-01-02T15:04:05.999Z"

func (ct *CustomTime) UnmarshalJSON(b []byte) (err error) {
	s := string(b)
	s = s[1 : len(s)-1] // Strip quotes
	ct.Time, err = time.Parse(ctLayout, s)
	return
}

func (ct *CustomTime) MarshalJSON() ([]byte, error) {
	return []byte(`"` + ct.Time.Format(ctLayout) + `"`), nil
}

func Unwrap(err error) (errs []string) {
	for err != nil {
		errs = append(errs, err.Error())
		err = errors.Unwrap(err)
	}
	return

}

func NewDimensionResponse(dimension logic.Dimension) (ndr DimensionResponse) {
	ndr.Dimension = make(map[string]string)

	for k, v := range dimension.Dimension {
		ndr.Dimension[k] = v.Color.String()
	}
	return
}

func (reqDim *Dimension) ToLogicDimension() (dimension *logic.Dimension, err error) {
	var pairs []logic.SpherePair

	// Helper function to generate SpherePair from SphereID and Color string
	addSpherePair := func(id logic.SphereID, colorStr string) {
		// If colorStr is empty, skip creating SpherePair
		if colorStr == "" {
			return
		}

		color := logic.NewColorShort(colorStr)
		pairs = append(pairs, *logic.NewSpherePair(id, color))
	}

	// Convert each field from REQUESTDimension
	addSpherePair("a", reqDim.A)
	addSpherePair("b", reqDim.B)
	addSpherePair("c", reqDim.C)
	addSpherePair("d", reqDim.D)
	addSpherePair("e", reqDim.E)
	addSpherePair("f", reqDim.F)
	addSpherePair("g", reqDim.G)
	addSpherePair("h", reqDim.H)
	addSpherePair("i", reqDim.I)
	addSpherePair("j", reqDim.J)
	addSpherePair("k", reqDim.K)
	addSpherePair("l", reqDim.L)
	addSpherePair("m", reqDim.M)
	addSpherePair("n", reqDim.N)

	// Create the dimension from SpherePairs
	dimension, err = logic.NewDimension(pairs...)
	if err != nil {
		return nil, errors.New("failed to create dimension: " + err.Error())
	}

	return dimension, nil
}
