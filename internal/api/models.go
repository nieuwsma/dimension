package api

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

// GeometryResponse represents a response for geometry
type GeometryResponse struct {
	Geometry []GeometryItem `json:"Geometry"`
}

// Color represents a color
type Color struct {
	Name string `json:"Name"`
	Code string `json:"Code"`
}
