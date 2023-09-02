package middleware

import (
	"github.com/nieuwsma/dimension/pkg/geometry"
	"github.com/nieuwsma/dimension/pkg/logic"
	"github.com/nieuwsma/dimension/pkg/rules"
)

// Rules Route

func GetGameRules() (rules.RuleSet, logic.Colors, geometry.Geometries) {
	// Logic to get game rules
	rules, _ := rules.GetRuleSet(rules.Default)
	colors := logic.GetColors()
	geometries := geometry.GetGeometry()

	return rules, colors, geometries
}
