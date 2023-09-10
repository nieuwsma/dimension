package tasks

import (
	_ "embed"
)

type TaskType int

const (
	EMPTY    TaskType = 0
	QUANTITY TaskType = 1
	GT       TaskType = 2 //greater than
	SUM      TaskType = 3
	TOP      TaskType = 4
	BOTTOM   TaskType = 5
	TOUCH    TaskType = 6
	NOTOUCH  TaskType = 7
)

func (s TaskType) String() string {
	switch s {
	case QUANTITY:
		return "QUANTITY"
	case GT:
		return "GT"
	case SUM:
		return "SUM"
	case TOP:
		return "TOP"
	case BOTTOM:
		return "BOTTOM"
	case TOUCH:
		return "TOUCH"
	case NOTOUCH:
		return "NOTOUCH"
	default:
		return "!"
	}
}

func (s TaskType) Equals(other TaskType) bool {
	return s == other
}
