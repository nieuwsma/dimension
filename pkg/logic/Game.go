package logic

import (
	"errors"
	"fmt"
	"strings"
	"time"
)

const RoundLength = 6

type Leaderboard struct {
	Round  int
	Scores map[PlayerName]ScoreRecord
	Alive  bool //is the game still active?
}

type ScoreRecord struct {
	Points      int
	BonusTokens int // bonus payout is: 0 = -6, 1 = -3, 2 = -1, 3 = 0, 4 = 1, 5 = 3, 6 = 6
}

func (s ScoreRecord) Equals(other ScoreRecord) bool {
	return s.Points == other.Points && s.BonusTokens == other.BonusTokens
}

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
	PlayerName     PlayerName
	Dimension      Dimension
	Score          int
	Bonus          bool
	TaskViolations []string
}

func (t Turn) String() string {
	bonusStr := "No"
	if t.Bonus {
		bonusStr = "Yes"
	}

	violationsStr := "None"
	if len(t.TaskViolations) > 0 {
		violationsStr = strings.Join(t.TaskViolations, "; ")
	}

	return fmt.Sprintf(
		"Player: %s\nDimension: %s\nScore: %d\nBonus: %s\nTask Violations: %s",
		t.PlayerName,
		t.Dimension,
		t.Score,
		bonusStr,
		violationsStr,
	)
}

type Game struct {
	Players        map[PlayerName]Player //name of player
	Rounds         map[int]Round         //tracks round ID to a play deck //  only 6 rounds are played
	Deck           Deck
	DrawSize       int
	HourglassLimit time.Duration
	Alive          bool
}

func NewGame(drawSize int, hourGlassLimit time.Duration, seed int64) (g *Game) {
	g = &Game{
		DrawSize:       drawSize,
		HourglassLimit: hourGlassLimit,
		Players:        make(map[PlayerName]Player),
		Rounds:         make(map[int]Round),
	}

	g.Deck = newDeck(seed)
	g.Deck.Shuffle()
	g.Alive = true
	return
}

func NewPlayer(name PlayerName) (player Player) {
	player = Player{
		PlayerName: name,
		Turns:      make(map[int]Turn),
	}
	return
}
func (g *Game) AddPlayer(name PlayerName) (err error) {
	if len(g.Rounds) > 0 {
		err = fmt.Errorf("cannot add %s to game in progress", name)
	}

	if _, exists := g.Players[name]; exists {
		err = fmt.Errorf("player %s already exists", name)
	} else {

		g.Players[name] = NewPlayer(name)
	}
	return
}

func (g *Game) RemovePlayer(name PlayerName) (err error) {
	if len(g.Rounds) > 0 {
		err = fmt.Errorf("cannot remove %s from game in progress", name)
	}

	if _, exists := g.Players[name]; exists {
		delete(g.Players, name)
	} else {
		err = fmt.Errorf("player %s doesn't exist", name)
	}
	return
}

// Used to start the game as well!

func (g *Game) NextRound() (activeTasks Tasks, err error) {

	if g.Alive {
		if len(g.Rounds) > 0 && !g.Rounds[len(g.Rounds)-1].Resolved {
			err = fmt.Errorf("current round must be resolved before next round can begin")
			return
		}

		if len(g.Rounds)+1 > RoundLength {
			g.Alive = false
			err = fmt.Errorf("the game has ended")
			return
		}
		activeTasks, err = g.Deck.Deal(g.DrawSize)
		if err != nil {
			//todo something!
		}
		round := Round{
			Tasks:    activeTasks,
			Resolved: false,
		}
		g.Rounds[len(g.Rounds)] = round //if the length WAS 5 our NEW round goes into [5]; which is now the sixth
	} else {
		err = fmt.Errorf("the game has ended")
	}

	return activeTasks, err
}

func (g *Game) GetCurrentRound() (roundCount int, round Round, err error) {
	if len(g.Rounds) == 0 {
		err = fmt.Errorf("no rounds exist")
	}
	round = g.Rounds[len(g.Rounds)-1]
	roundCount = len(g.Rounds)

	return roundCount, round, err
}

// will close a round if all players have taken a turn; can force delinquent players to forefit the round if they havent played
func (g *Game) EndRound(force bool) (err error) {
	if !g.Rounds[len(g.Rounds)-1].Resolved {
		delinquentPlayers := 0
		//then we need to resolve it!
		roundCount := len(g.Rounds)
		for playername, player := range g.Players {
			if len(player.Turns) < roundCount {
				if force {
					//play an empty dimension for the player
					emptyDimension, _ := NewDimension()
					g.PlayTurn(playername, *emptyDimension)
					err = errors.Join(fmt.Errorf("%s has not submitted a turn for round, they have forfitted this round! %d", playername, roundCount))
				} else {
					//this player is delinquent!
					delinquentPlayers++
					err = errors.Join(fmt.Errorf("%s has not submitted a turn for round %d", playername, roundCount))
				}
			}

		}
		if delinquentPlayers == 0 {
			round := g.Rounds[len(g.Rounds)-1]
			round.Resolved = true
			g.Rounds[len(g.Rounds)-1] = round
		}
	}
	return
}

func (g *Game) GetLeaderboard() (leaderboard Leaderboard) {
	leaderboard.Scores = make(map[PlayerName]ScoreRecord)
	leaderboard.Round = len(g.Rounds) //rounds actually start at 1, not zero!
	for _, player := range g.Players {
		leaderboard.Scores[player.PlayerName] = player.ScoreRecord
	}
	leaderboard.Alive = g.Alive
	return
}

func (g *Game) PlayTurn(playerName PlayerName, dim Dimension) (turn Turn, err error) {
	if g.Alive {
		score, bonus, taskViolations, errors := ScoreTurn(g.Rounds[len(g.Rounds)-1].Tasks, dim)
		if errors != nil {
			return turn, errors
		}
		turn = Turn{
			Dimension:      dim,
			Score:          score,
			Bonus:          bonus,
			TaskViolations: taskViolations,
		}
		playerRecord := g.Players[playerName]
		playerRecord.Turns[len(g.Rounds)] = turn
		playerRecord.ScoreRecord.Points += score
		if bonus {
			playerRecord.ScoreRecord.BonusTokens++
		}
		g.Players[playerName] = playerRecord
	} else {
		err = fmt.Errorf("the game has ended")
	}
	return
}

func (g *Game) EndGame(force bool) (err error) {

	if len(g.Rounds) == RoundLength && g.Rounds[RoundLength-1].Resolved {
		g.Alive = false
	} else {
		g.Alive = true
	}

	if force {
		g.Alive = false
		for roundID, round := range g.Rounds {
			round.Resolved = true
			g.Rounds[roundID] = round
		}
	}

	return
}

func (g *Game) DeepCopy() *Game {
	// Copy basic fields
	newGame := &Game{
		DrawSize:       g.DrawSize,
		HourglassLimit: g.HourglassLimit,
		Alive:          g.Alive,
		Deck: Deck{
			NextTaskIndex: g.Deck.NextTaskIndex,
			Seed:          g.Deck.Seed,
			RuleSetName:   g.Deck.RuleSetName,
		},
	}

	// Copy Players map
	newGame.Players = make(map[PlayerName]Player)
	for k, v := range g.Players {
		newTurns := make(map[int]Turn)
		for turnKey, turnVal := range v.Turns {
			newTurns[turnKey] = turnVal
		}

		newGame.Players[k] = Player{
			PlayerName:  v.PlayerName,
			Turns:       newTurns,
			ScoreRecord: v.ScoreRecord,
		}
	}

	// Copy Rounds map
	newGame.Rounds = make(map[int]Round)
	for k, v := range g.Rounds {
		newGame.Rounds[k] = v
	}

	// Copy Deck.DrawPile (Tasks slice)
	newTasks := make(Tasks, len(g.Deck.DrawPile))
	copy(newTasks, g.Deck.DrawPile)
	newGame.Deck.DrawPile = newTasks

	return newGame
}
