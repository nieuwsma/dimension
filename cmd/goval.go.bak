package main

import (
	"fmt"
	"github.com/Knetic/govaluate"
)

func main() {
	expressionStr := `
		(b || !b) &&
		(c || !c) &&
		(d || !d) &&
		(e || !e) &&
		(f || !f) &&
		(g || !g) &&
		((h && b && c && a) || !h) &&
		((i && c && d && a) || !i) &&
		((j && d && e && a) || !j) &&
		((k && e && f && a) || !k) &&
		((l && f && g && a) || !l) &&
		((m && g && b && a) || !m) &&
		((!h && !i && !j && !k && !l && !m) || a) &&
		(!n || (n && ((h && j && l) || (i && k && m)) && a && b && c && d && e && f && g)) &&
		(((h && j && l) || (i && k && m) || !n) || ((h && k && !i && !j && !l && !m) || (i && l && !h && !j && !k && !m) || (j && m && !h && !i && !k && !l)))
	`

	parameters := make(map[string]interface{}, 8)
	parameters["a"] = true // Or false, depending on the configuration
	parameters["b"] = true // And so on for the rest of the parameters

	expression, _ := govaluate.NewEvaluableExpression(expressionStr)
	result, _ := expression.Evaluate(parameters)

	fmt.Println(result)
}
