package domain

import (
	_ "embed"
	"encoding/json"
)

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

//go:embed colors.json
var ColorsData []byte

func GetColors() (c Colors) {
	_ = json.Unmarshal(ColorsData, &c)
	return
}
