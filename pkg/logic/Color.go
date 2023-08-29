package logic

import (
	_ "embed"
	"strings"
)

func NewColorShort(c string) (color Color) {
	c = strings.ToUpper(c)
	switch c {
	case "G":
		return Green
	case "B":
		return Blue
	case "K":
		return Black
	case "O":
		return Orange
	case "W":
		return White
	default:
		return Empty
	}
}

func NewColorLong(c string) (color Color) {
	c = strings.ToUpper(c)
	switch c {
	case "GREEN":
		return Green
	case "BLUE":
		return Blue
	case "BLACK":
		return Black
	case "ORANGE":
		return Orange
	case "WHITE":
		return White
	default:
		return Empty
	}
}

func (g Color) LongHand() string {
	switch g {
	case Green:
		return "GREEN"
	case Blue:
		return "BLUE"
	case Black:
		return "BLACK"
	case Orange:
		return "ORANGE"
	case White:
		return "WHITE"
	default:
		return ""
	}
}

type Color int

const (
	Empty  Color = 0
	Green  Color = 1
	Blue   Color = 2
	White  Color = 3
	Orange Color = 4
	Black  Color = 5
)

func (s Color) String() string {
	switch s {
	case Green:
		return "G"
	case Blue:
		return "B"
	case Black:
		return "K"
	case Orange:
		return "O"
	case White:
		return "W"
	default:
		return ""
	}
}

func (s Color) Equals(other Color) bool {
	return s == other
}

type Colors []ColorRecord

type ColorRecord struct {
	Name string `json:"name"`
	Code string `json:"code"`
}

var colors = Colors{
	{
		Name: "GREEN",
		Code: "G",
	},
	{
		Name: "ORANGE",
		Code: "O",
	},
	{
		Name: "BLACK",
		Code: "K",
	},
	{
		Name: "WHITE",
		Code: "W",
	},
	{
		Name: "BLUE",
		Code: "B",
	},
}

func GetColors() Colors {
	return colors
}
