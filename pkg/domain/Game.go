package domain

import (
	"errors"
	"fmt"
	"time"
)

//TODO consider adding a game history and tracking players over time! #goldplatting
//type GameHistory []Game

const RoundLength = 6

type Leaderboard struct {
	Round  int
	Scores map[PlayerName]ScoreRecord
}

type ScoreRecord struct {
	Points      int
	BonusTokens int //TODO bonus payout is: 0 = -6, 1 = -3, 2 = -1, 3 = 0, 4 = 1, 5 = 3, 6 = 6
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
	Dimension      Dimension
	Score          int
	Bonus          bool
	TaskViolations error
}

type Game struct {
	Players        map[PlayerName]Player //name of player
	Rounds         []Round               //tracks round ID to a play deck // TODO only 6 rounds are played
	Deck           Deck
	DrawSize       int
	HourglassLimit time.Duration
}

func NewGame(drawSize int, hourGlassLimit time.Duration, seed int64) (g *Game) {
	g = &Game{
		DrawSize:       drawSize,
		HourglassLimit: hourGlassLimit,
		Players:        make(map[PlayerName]Player),
	}

	g.Deck = newDeck(seed)
	g.Deck.Shuffle()
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
// todo max out at 6 founds
func (g *Game) NextRound() (activeTasks Tasks, err error) {

	if len(g.Rounds) > 0 && !g.Rounds[len(g.Rounds)-1].Resolved {
		err = fmt.Errorf("current round must be resolved before next round can begin")
		return
	}

	if len(g.Rounds) >= RoundLength {
		err = fmt.Errorf("this game is over")
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
	g.Rounds = append(g.Rounds, round)
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
			g.Rounds[len(g.Rounds)-1].Resolved = true
		}
	}
	return
}

func (g *Game) GetLeaderboard() (leaderboard Leaderboard) {
	leaderboard.Scores = make(map[PlayerName]ScoreRecord)
	leaderboard.Round = len(g.Rounds) + 1 //rounds actually start at 1, not zero!
	for _, player := range g.Players {
		leaderboard.Scores[player.PlayerName] = player.ScoreRecord
	}
	return
}

// todo somehow I need to account for the fact that naughty players might not play a turn each round, which they need to do!
// need to also account for not being able to start next round without resolving the last round?
func (g *Game) PlayTurn(playerName PlayerName, dim Dimension) (turn Turn) {

	score, bonus, errors := ScoreTurn(g.Rounds[len(g.Rounds)-1].Tasks, dim)
	turn = Turn{
		Dimension:      dim,
		Score:          score,
		Bonus:          bonus,
		TaskViolations: errors,
	}
	playerRecord := g.Players[playerName]
	playerRecord.Turns[len(g.Rounds)] = turn
	playerRecord.ScoreRecord.Points += score
	if bonus {
		playerRecord.ScoreRecord.BonusTokens++
	}
	g.Players[playerName] = playerRecord

	return
}
