package domain

type Deck struct {
	DrawPile    Tasks
	DiscardPile Tasks
	Active      Tasks
	Seed        int64
}

type Tasks []Task
type Task string
