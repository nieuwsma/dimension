package domain

import (
	_ "embed"
	"encoding/json"
)

type Rules map[int]Rule

type Rule struct {
	Name        string
	Quantity    int
	Description string
}

//go:embed rules.json
var RulesData []byte

func GetAllRules() (r Rules) {
	_ = json.Unmarshal(RulesData, &r)
	return
}

//var Hand Rules
//var Deck Rules
//var Discard Rules
