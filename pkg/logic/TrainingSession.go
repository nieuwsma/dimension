package logic

import "time"

type TrainingSession struct {
	Turn  Turn
	Tasks Tasks
	Deck  Deck

	DrawSize       int
	ExpirationTime time.Time
}

func NewTrainingSession(drawSize int, seed int64) (t *TrainingSession) {
	t = &TrainingSession{
		DrawSize: drawSize,
	}
	t.Deck = newDeck(seed)
	t.Deck.Shuffle()
	t.Tasks, _ = t.Deck.Deal(t.DrawSize)
	t.ExpirationTime = time.Now().Add(1 * time.Hour)
	return
}

func (g *TrainingSession) PlayTurn(dim Dimension) {
	score, bonus, errors := ScoreTurn(g.Tasks, dim)
	g.Turn = Turn{
		Dimension:      dim,
		Score:          score,
		Bonus:          bonus,
		TaskViolations: errors,
	}
	g.ExpirationTime = time.Now().Add(1 * time.Hour)

	return
}

func (g *TrainingSession) NextRound() {
	g.Tasks, _ = g.Deck.Deal(g.DrawSize)
	g.Turn = Turn{}
	return
}
