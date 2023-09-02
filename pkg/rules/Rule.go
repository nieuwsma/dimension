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
		{"QUANTITY-1-G", 1, "The quantity of GREEN should be exactly 1."},
		{"QUANTITY-1-O", 1, "The quantity of ORANGE should be exactly 1."},
		{"QUANTITY-1-K", 1, "The quantity of BLACK should be exactly 1."},
		{"QUANTITY-1-W", 1, "The quantity of WHITE should be exactly 1."},
		{"QUANTITY-1-B", 1, "The quantity of BLUE should be exactly 1."},
		{"QUANTITY-2-G", 1, "The quantity of GREEN should be exactly 2."},
		{"QUANTITY-2-O", 1, "The quantity of ORANGE should be exactly 2."},
		{"QUANTITY-2-K", 1, "The quantity of BLACK should be exactly 2."},
		{"QUANTITY-2-W", 1, "The quantity of WHITE should be exactly 2."},
		{"QUANTITY-2-B", 1, "The quantity of BLUE should be exactly 2."},
		{"BOTTOM-G", 1, "No sphere of any color may be below GREEN."},
		{"BOTTOM-O", 1, "No sphere of any color may be below ORANGE."},
		{"BOTTOM-K", 1, "No sphere of any color may be below BLACK."},
		{"BOTTOM-W", 1, "No sphere of any color may be below WHITE."},
		{"BOTTOM-B", 1, "No sphere of any color may be below BLUE."},
		{"TOP-G", 1, "No sphere of any color may be on top of GREEN."},
		{"TOP-O", 1, "No sphere of any color may be on top of ORANGE."},
		{"TOP-K", 1, "No sphere of any color may be on top of BLACK."},
		{"TOP-W", 1, "No sphere of any color may be on top of WHITE."},
		{"TOP-B", 1, "No sphere may be above BLUE."},
		{"TOUCH-G-G", 1, "Every GREEN must touch another GREEN."},
		{"TOUCH-O-O", 1, "Every ORANGE must touch another ORANGE."},
		{"TOUCH-K-K", 1, "Every BLACK must touch another BLACK."},
		{"TOUCH-W-W", 1, "Every WHITE must touch another WHITE."},
		{"TOUCH-B-B", 1, "Every BLUE must touch another BLUE."},
		{"NOTOUCH-G-G", 1, "Every GREEN must not touch another GREEN."},
		{"NOTOUCH-O-O", 1, "Every ORANGE must not touch another ORANGE."},
		{"NOTOUCH-K-K", 1, "Every BLACK must not touch another BLACK."},
		{"NOTOUCH-W-W", 1, "Every WHITE must not touch another WHITE."},
		{"NOTOUCH-B-B", 1, "Every BLUE must not touch another BLUE."},
		{"TOUCH-G-O", 1, "Every GREEN must touch another ORANGE."},
		{"TOUCH-G-B", 1, "Every GREEN must touch another BLUE."},
		{"TOUCH-G-W", 1, "Every GREEN must touch another WHITE."},
		{"TOUCH-G-K", 1, "Every GREEN must touch another BLACK."},
		{"TOUCH-O-B", 1, "Every ORANGE must touch another BLUE."},
		{"TOUCH-O-W", 1, "Every ORANGE must touch another WHITE."},
		{"TOUCH-O-K", 1, "Every ORANGE must touch another BLACK."},
		{"TOUCH-K-W", 1, "Every BLACK must touch another WHITE."},
		{"TOUCH-K-B", 1, "Every BLACK must touch another BLUE."},
		{"TOUCH-B-W", 1, "Every BLUE must touch another WHITE."},
		{"NOTOUCH-G-O", 1, "Every GREEN must not touch another ORANGE."},
		{"NOTOUCH-G-B", 1, "Every GREEN must not touch another BLUE."},
		{"NOTOUCH-G-W", 1, "Every GREEN must not touch another WHITE."},
		{"NOTOUCH-G-K", 1, "Every GREEN must not touch another BLACK."},
		{"NOTOUCH-O-B", 1, "Every ORANGE must not touch another BLUE."},
		{"NOTOUCH-O-W", 1, "Every ORANGE must not touch another WHITE."},
		{"NOTOUCH-O-K", 1, "Every ORANGE must not touch another BLACK."},
		{"NOTOUCH-K-W", 1, "Every BLACK must not touch another WHITE."},
		{"NOTOUCH-K-B", 1, "Every BLACK must not touch another BLUE."},
		{"NOTOUCH-B-W", 1, "Every BLUE must not touch another WHITE."},
		{"SUM-4-G-K", 1, "The sum of GREEN and BLACK must be exactly 4."},
		{"SUM-4-K-W", 1, "The sum of BLACK and WHITE must be exactly 4."},
		{"SUM-4-W-O", 1, "The sum of WHITE and ORANGE must be exactly 4."},
		{"SUM-4-O-B", 1, "The sum of ORANGE and BLUE must be exactly 4."},
		{"SUM-4-B-G", 1, "The sum of BLUE and GREEN must be exactly 4."},
		{"GT-W-G", 1, "The count of WHITE must be greater than GREEN."},
		{"GT-G-O", 1, "The count of GREEN must be greater than ORANGE."},
		{"GT-O-K", 1, "The count of ORANGE must be greater than BLACK."},
		{"GT-K-B", 1, "The count of BLACK must be greater than BLUE."},
		{"GT-B-W", 1, "The count of BLUE must be greater than WHITE."},
	},
}
