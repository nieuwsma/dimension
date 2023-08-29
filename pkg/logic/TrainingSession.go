package logic

import "time"

type TrainingSession struct {
	Turn           Turn
	Tasks          Tasks
	Deck           Deck
	DrawSize       int
	HourglassLimit time.Duration
	Alive          bool
}

func NewTrainingSession(drawSize int, hourGlassLimit time.Duration, seed int64) (t *TrainingSession) {
	t = &TrainingSession{
		DrawSize:       drawSize,
		HourglassLimit: hourGlassLimit,
	}
	t.Tasks, _ = t.Deck.Deal(t.DrawSize)
	t.Deck = newDeck(seed)
	t.Deck.Shuffle()
	t.Alive = true
	return
}

func (g *TrainingSession) PlayTurn(dim Dimension) {
	if g.Alive {
		score, bonus, errors := ScoreTurn(g.Tasks, dim)
		g.Turn = Turn{
			Dimension:      dim,
			Score:          score,
			Bonus:          bonus,
			TaskViolations: errors,
		}
	}
	return
}

func (g *TrainingSession) NextRound() {
	if g.Alive {
		g.Tasks, _ = g.Deck.Deal(g.DrawSize)
		g.Turn = Turn{}
	}
	return
}
