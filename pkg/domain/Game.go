package domain

import (
	"math/rand"
	"time"
)

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
