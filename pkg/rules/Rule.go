package rules

import (
	_ "embed"
)

type RuleSet struct {
	Set []Rule `json:"rules"`
}

type Rule struct {
	Name        string `json:"name"`
	Quantity    int    `json:"quantity"`
	Description string `json:"description"`
}

var rulesSets = make(map[string]RuleSet)

func init() {
	rulesSets["default"] = defaultRuleSet
}

// always return a rule set,
func GetRuleSet(name string) (rules RuleSet, ruleSetName string) {

	rules, exists := rulesSets[name]
	if !exists {
		rules = rulesSets[Default]
		return rules, Default
	}

	return rules, name
}

const Default string = "default"

var defaultRuleSet = RuleSet{
	Set: []Rule{
		{"QUANTITY-1-G", 1, "The quantity of green should be exactly 1."},
		{"QUANTITY-1-O", 1, "The quantity of orange should be exactly 1."},
		{"QUANTITY-1-K", 1, "The quantity of black should be exactly 1."},
		{"QUANTITY-1-W", 1, "The quantity of white should be exactly 1."},
		{"QUANTITY-1-B", 1, "The quantity of blue should be exactly 1."},
		{"QUANTITY-2-G", 1, "The quantity of green should be exactly 2."},
		{"QUANTITY-2-O", 1, "The quantity of orange should be exactly 2."},
		{"QUANTITY-2-K", 1, "The quantity of black should be exactly 2."},
		{"QUANTITY-2-W", 1, "The quantity of white should be exactly 2."},
		{"QUANTITY-2-B", 1, "The quantity of blue should be exactly 2."},
		{"BOTTOM-G", 1, "No sphere of any color may be below green."},
		{"BOTTOM-O", 1, "No sphere of any color may be below orange."},
		{"BOTTOM-K", 1, "No sphere of any color may be below black."},
		{"BOTTOM-W", 1, "No sphere of any color may be below white."},
		{"BOTTOM-B", 1, "No sphere of any color may be below blue."},
		{"TOP-G", 1, "No sphere of any color may be on top of green."},
		{"TOP-O", 1, "No sphere of any color may be on top of orange."},
		{"TOP-K", 1, "No sphere of any color may be on top of black."},
		{"TOP-W", 1, "No sphere of any color may be on top of white."},
		{"TOP-B", 1, "No sphere may be above blue."},
		{"TOUCH-G-G", 1, "Every green must touch another green."},
		{"TOUCH-O-O", 1, "Every orange must touch another orange."},
		{"TOUCH-K-K", 1, "Every black must touch another black."},
		{"TOUCH-W-W", 1, "Every white must touch another white."},
		{"TOUCH-B-B", 1, "Every blue must touch another blue."},
		{"NOTOUCH-G-G", 1, "Every green must not touch another green."},
		{"NOTOUCH-O-O", 1, "Every orange must not touch another orange."},
		{"NOTOUCH-K-K", 1, "Every black must not touch another black."},
		{"NOTOUCH-W-W", 1, "Every white must not touch another white."},
		{"NOTOUCH-B-B", 1, "Every blue must not touch another blue."},
		{"TOUCH-G-O", 1, "Every green must touch another orange."},
		{"TOUCH-G-B", 1, "Every green must touch another blue."},
		{"TOUCH-G-W", 1, "Every green must touch another white."},
		{"TOUCH-G-K", 1, "Every green must touch another black."},
		{"TOUCH-O-B", 1, "Every orange must touch another blue."},
		{"TOUCH-O-W", 1, "Every orange must touch another white."},
		{"TOUCH-O-K", 1, "Every orange must touch another black."},
		{"TOUCH-K-W", 1, "Every black must touch another white."},
		{"TOUCH-K-B", 1, "Every black must touch another blue."},
		{"TOUCH-B-W", 1, "Every blue must touch another white."},
		{"NOTOUCH-G-O", 1, "Every green must not touch another orange."},
		{"NOTOUCH-G-B", 1, "Every green must not touch another blue."},
		{"NOTOUCH-G-W", 1, "Every green must not touch another white."},
		{"NOTOUCH-G-K", 1, "Every green must not touch another black."},
		{"NOTOUCH-O-B", 1, "Every orange must not touch another blue."},
		{"NOTOUCH-O-W", 1, "Every orange must not touch another white."},
		{"NOTOUCH-O-K", 1, "Every orange must not touch another black."},
		{"NOTOUCH-K-W", 1, "Every black must not touch another white."},
		{"NOTOUCH-K-B", 1, "Every black must not touch another blue."},
		{"NOTOUCH-B-W", 1, "Every blue must not touch another white."},
		{"RATIO-4-G-K", 1, "The sum of green and black must be exactly 4."},
		{"RATIO-4-G-O", 1, "The sum of green and orange must be exactly 4."},
		{"RATIO-4-G-W", 1, "The sum of green and white must be exactly 4."},
		{"RATIO-4-G-B", 1, "The sum of green and blue must be exactly 4."},
		{"RATIO-4-O-K", 1, "The sum of orange and black must be exactly 4."},
		{"RATIO-4-O-W", 1, "The sum of orange and white must be exactly 4."},
		{"RATIO-4-O-B", 1, "The sum of orange and blue must be exactly 4."},
		{"RATIO-4-K-W", 1, "The sum of black and white must be exactly 4."},
		{"RATIO-4-K-B", 1, "The sum of black and blue must be exactly 4."},
		{"RATIO-4-W-B", 1, "The sum of white and blue must be exactly 4."},
		{"GT-W-G", 1, "The count of white must be greater than green."},
		{"GT-G-O", 1, "The count of green must be greater than orange."},
		{"GT-O-K", 1, "The count of orange must be greater than black."},
		{"GT-K-B", 1, "The count of black must be greater than blue."},
		{"GT-B-W", 1, "The count of blue must be greater than white."},
	},
}
