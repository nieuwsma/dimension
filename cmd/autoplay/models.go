package main

import "github.com/nieuwsma/dimension/pkg/logic"

type Quantity struct {
	Number int
	Color  logic.Color
}

type Sum struct {
	Total  int
	Colors []logic.Color
}

type GreaterThan struct {
	GreaterColor logic.Color
	LesserColor  logic.Color
}

type Bottom struct {
	Color logic.Color
}

type Top struct {
	Color logic.Color
}

type Touch struct {
	Colors []logic.Color
}

type NoTouch struct {
	Colors []logic.Color
}

type TasksCollection struct {
	NoTouch     map[string]NoTouch
	Touch       map[string]Touch
	Top         map[string]Top
	Bottom      map[string]Bottom
	GreaterThan map[string]GreaterThan
	Sum         map[string]Sum
	Quantity    map[string]Quantity
}

func NewTasksCollection() (t *TasksCollection) {
	t = &TasksCollection{
		NoTouch:     make(map[string]NoTouch),
		Touch:       make(map[string]Touch),
		Top:         make(map[string]Top),
		Bottom:      make(map[string]Bottom),
		GreaterThan: make(map[string]GreaterThan),
		Sum:         make(map[string]Sum),
		Quantity:    make(map[string]Quantity),
	}
	return
}
