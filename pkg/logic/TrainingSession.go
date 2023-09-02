package logic

import "time"

type TrainingSession struct {
	Turns    map[PlayerName]Turn
	Tasks    Tasks
	Deck     Deck
	DrawSize int
	//todo need to implement the ability to delete old expired training sessions
	ExpirationTime time.Time
}

func NewTrainingSession(drawSize int, seed int64) (t *TrainingSession) {
	t = &TrainingSession{
		DrawSize: drawSize,
		Turns:    make(map[PlayerName]Turn),
	}

	t.Deck = newDeck(seed)
	t.Deck.Shuffle()
	t.Tasks, _ = t.Deck.Deal(t.DrawSize)
	t.ExpirationTime = time.Now().Add(1 * time.Hour)
	return
}

func (g *TrainingSession) PlayTurn(playerName PlayerName, dim Dimension) {
	score, bonus, errors := ScoreTurn(g.Tasks, dim)
	g.Turns[playerName] = Turn{
		PlayerName:     playerName,
		Dimension:      dim,
		Score:          score,
		Bonus:          bonus,
		TaskViolations: errors,
	}
	g.ExpirationTime = time.Now().Add(1 * time.Hour)

	return
}

func (g *TrainingSession) Regenerate() {
	g.Tasks, _ = g.Deck.Deal(g.DrawSize)
	g.Turns = make(map[PlayerName]Turn)
	return
}
