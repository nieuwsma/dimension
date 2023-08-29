package logic

import (
	"dimension/pkg/rules"
	"errors"
	"math/rand"
)

type Deck struct {
	DrawPile      Tasks
	NextTaskIndex int
	Seed          int64
	RuleSetName   string
}

type Tasks []Task
type Task string

func (d *Deck) Deal(size int) (tasks Tasks, err error) {

	if d.NextTaskIndex+size <= len(d.DrawPile) { //todo this might be off by one
		d.Shuffle()
		d.NextTaskIndex = 0
	}
	for i := 0; i < size; i++ {
		if d.NextTaskIndex < len(d.DrawPile) {
			tasks = append(tasks, d.DrawPile[d.NextTaskIndex])
			d.NextTaskIndex++
		} else {
			err = errors.New("ran out of cards")
			return nil, err
		}

	}

	return tasks, err
}

func newDeck(seed int64) (d Deck) {

	defaultRules, ruleSetName := rules.GetRuleSet(rules.Default)
	d.RuleSetName = ruleSetName
	for _, r := range defaultRules.Set {
		d.DrawPile = append(d.DrawPile, Task(r.Name))
	}
	d.Seed = seed
	return d
}

func (d *Deck) Shuffle() {

	d.NextTaskIndex = 0

	rand.New(rand.NewSource(d.Seed))
	rand.Shuffle(len(d.DrawPile), func(i, j int) { d.DrawPile[i], d.DrawPile[j] = d.DrawPile[j], d.DrawPile[i] })
}
