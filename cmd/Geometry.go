package main

import (
	"errors"
	"strconv"
	"strings"
)

type Geometry struct {
	PolarAngle       float64 `json:"polarAngle"`
	InclinationAngle float64 `json:"inclinationAngle"`
	RadialDistance   float64 `json:"radialDistance"`
	ID               int     `json:"id"`
	Adjacency        string  `json:"adjacency"`
	Required         *string `json:"required"`   // using pointer since it can be null
	Prohibited       *string `json:"prohibited"` // using pointer since it can be null
}

type GeometryList struct {
	Geometry []Geometry `json:"Geometry"`
}

type NotExpr struct {
	Operand Expr
}

func (n *NotExpr) Evaluate(dim *Dimension) bool {
	return !n.Operand.Evaluate(dim)
}

// Expr represents a boolean expression.
type Expr interface {
	Evaluate(dim *Dimension) bool
}

type AndExpr struct {
	Left, Right Expr
}

func (a *AndExpr) Evaluate(dim *Dimension) bool {
	if a.Left == nil || a.Right == nil {
		return false
	}
	return a.Left.Evaluate(dim) && a.Right.Evaluate(dim)
}

type XorExpr struct {
	Left, Right Expr
}

func (x *XorExpr) Evaluate(dim *Dimension) bool {
	return x.Left.Evaluate(dim) != x.Right.Evaluate(dim)
}

type OrExpr struct {
	Left, Right Expr
}

func (o *OrExpr) Evaluate(dim *Dimension) bool {
	if o.Left == nil || o.Right == nil {
		return false
	}
	return o.Left.Evaluate(dim) || o.Right.Evaluate(dim)
}

type VarExpr struct {
	Value int
}

func (v *VarExpr) Evaluate(dim *Dimension) bool {
	sphere, exists := dim.Dimension[v.Value]
	return exists && sphere != nil
}

func Parse(expr string) (Expr, error) {
	expr = strings.ReplaceAll(expr, " ", "")
	return parseRecursively(expr)
}

////corrrect but no xor
//func parseRecursively(expr string) (Expr, error) {
//	expr = strings.TrimSpace(expr)
//
//	// Check for surrounding parentheses
//	if strings.HasPrefix(expr, "(") && strings.HasSuffix(expr, ")") {
//		// Check if the entire expression is surrounded by parentheses
//		depth := 0
//		for i := 0; i < len(expr); i++ {
//			switch expr[i] {
//			case '(':
//				depth++
//			case ')':
//				depth--
//				if depth == 0 && i < len(expr)-1 {
//					goto NoSurroundingParens
//				}
//			}
//		}
//		return parseRecursively(expr[1 : len(expr)-1])
//	}
//
//NoSurroundingParens:
//	// Parse "OR" expression, ensuring we're getting the outermost '|' not enclosed in parentheses
//	for i, depth := 0, 0; i < len(expr); i++ {
//		switch expr[i] {
//		case '(':
//			depth++
//		case ')':
//			depth--
//		case '|':
//			if depth == 0 {
//				left, err := parseRecursively(expr[:i])
//				if err != nil {
//					return nil, err
//				}
//				right, err := parseRecursively(expr[i+1:])
//				if err != nil {
//					return nil, err
//				}
//				return &OrExpr{left, right}, nil
//			}
//		}
//	}
//
//	// Parse "AND" expression, ensuring we're getting the outermost '&' not enclosed in parentheses
//	for i, depth := 0, 0; i < len(expr); i++ {
//		switch expr[i] {
//		case '(':
//			depth++
//		case ')':
//			depth--
//		case '&':
//			if depth == 0 {
//				left, err := parseRecursively(expr[:i])
//				if err != nil {
//					return nil, err
//				}
//				right, err := parseRecursively(expr[i+1:])
//				if err != nil {
//					return nil, err
//				}
//				return &AndExpr{left, right}, nil
//			}
//		}
//	}
//
//	// If neither '|' nor '&' are found, it should be a value
//	value, err := strconv.Atoi(expr)
//	if err != nil {
//		return nil, errors.New("invalid expression")
//	}
//	return &VarExpr{Value: value}, nil
//}
//
////incorrect w xor
//func parseRecursively(expr string) (Expr, error) {
//	expr = strings.TrimSpace(expr)
//
//	// Strip surrounding parentheses, if any
//	if strings.HasPrefix(expr, "(") && strings.HasSuffix(expr, ")") {
//		return parseRecursively(expr[1 : len(expr)-1])
//	}
//
//	// Handle XOR
//	if idx := strings.Index(expr, "^"); idx != -1 {
//		left, err := parseRecursively(expr[:idx])
//		if err != nil {
//			return nil, err
//		}
//
//		right, err := parseRecursively(expr[idx+1:])
//		if err != nil {
//			return nil, err
//		}
//
//		return &XorExpr{left, right}, nil
//	}
//
//	// Handle OR
//	if idx := strings.Index(expr, "|"); idx != -1 {
//		left, err := parseRecursively(expr[:idx])
//		if err != nil {
//			return nil, err
//		}
//
//		right, err := parseRecursively(expr[idx+1:])
//		if err != nil {
//			return nil, err
//		}
//
//		return &OrExpr{left, right}, nil
//	}
//
//	// Handle AND
//	if idx := strings.Index(expr, "&"); idx != -1 {
//		left, err := parseRecursively(expr[:idx])
//		if err != nil {
//			return nil, err
//		}
//
//		right, err := parseRecursively(expr[idx+1:])
//		if err != nil {
//			return nil, err
//		}
//
//		return &AndExpr{left, right}, nil
//	}
//
//	value, err := strconv.Atoi(expr)
//	if err != nil {
//		return nil, errors.New("invalid expression")
//	}
//
//	return &VarExpr{Value: value}, nil
//}

// hopefully correct w xor
//func parseRecursively(expr string) (Expr, error) {
//	expr = strings.TrimSpace(expr)
//
//	// Check for surrounding parentheses
//	if strings.HasPrefix(expr, "(") && strings.HasSuffix(expr, ")") {
//		// Check if the entire expression is surrounded by parentheses
//		depth := 0
//		for i := 0; i < len(expr); i++ {
//			switch expr[i] {
//			case '(':
//				depth++
//			case ')':
//				depth--
//				if depth == 0 && i < len(expr)-1 {
//					goto NoSurroundingParens
//				}
//			}
//		}
//		return parseRecursively(expr[1 : len(expr)-1])
//	}
//
//NoSurroundingParens:
//
//	// Parse "XOR" expression, ensuring we're getting the outermost '^' not enclosed in parentheses
//	for i, depth := 0, 0; i < len(expr); i++ {
//		switch expr[i] {
//		case '(':
//			depth++
//		case ')':
//			depth--
//		case '^':
//			if depth == 0 {
//				left, err := parseRecursively(expr[:i])
//				if err != nil {
//					return nil, err
//				}
//				right, err := parseRecursively(expr[i+1:])
//				if err != nil {
//					return nil, err
//				}
//				return &XorExpr{left, right}, nil
//			}
//		}
//	}
//
//	// Parse "OR" expression, ensuring we're getting the outermost '|' not enclosed in parentheses
//	for i, depth := 0, 0; i < len(expr); i++ {
//		switch expr[i] {
//		case '(':
//			depth++
//		case ')':
//			depth--
//		case '|':
//			if depth == 0 {
//				left, err := parseRecursively(expr[:i])
//				if err != nil {
//					return nil, err
//				}
//				right, err := parseRecursively(expr[i+1:])
//				if err != nil {
//					return nil, err
//				}
//				return &OrExpr{left, right}, nil
//			}
//		}
//	}
//
//	// Parse "AND" expression, ensuring we're getting the outermost '&' not enclosed in parentheses
//	for i, depth := 0, 0; i < len(expr); i++ {
//		switch expr[i] {
//		case '(':
//			depth++
//		case ')':
//			depth--
//		case '&':
//			if depth == 0 {
//				left, err := parseRecursively(expr[:i])
//				if err != nil {
//					return nil, err
//				}
//				right, err := parseRecursively(expr[i+1:])
//				if err != nil {
//					return nil, err
//				}
//				return &AndExpr{left, right}, nil
//			}
//		}
//	}
//
//	// If neither '|', '&', nor '^' are found, it should be a value
//	value, err := strconv.Atoi(expr)
//	if err != nil {
//		return nil, errors.New("invalid expression")
//	}
//	return &VarExpr{Value: value}, nil
//}

// hopefully correct w xor and not
func parseRecursively(expr string) (Expr, error) {
	expr = strings.TrimSpace(expr)

	// Check for surrounding parentheses
	if strings.HasPrefix(expr, "(") && strings.HasSuffix(expr, ")") {
		// Check if the entire expression is surrounded by parentheses
		depth := 0
		for i := 0; i < len(expr); i++ {
			switch expr[i] {
			case '(':
				depth++
			case ')':
				depth--
				if depth == 0 && i < len(expr)-1 {
					goto NoSurroundingParens
				}
			}
		}
		return parseRecursively(expr[1 : len(expr)-1])
	}

NoSurroundingParens:

	// Handle NOT
	if strings.HasPrefix(expr, "!") {
		operand, err := parseRecursively(expr[1:])
		if err != nil {
			return nil, err
		}
		return &NotExpr{Operand: operand}, nil
	}

	// Parse "XOR" expression, ensuring we're getting the outermost '^' not enclosed in parentheses
	for i, depth := 0, 0; i < len(expr); i++ {
		switch expr[i] {
		case '(':
			depth++
		case ')':
			depth--
		case '^':
			if depth == 0 {
				left, err := parseRecursively(expr[:i])
				if err != nil {
					return nil, err
				}
				right, err := parseRecursively(expr[i+1:])
				if err != nil {
					return nil, err
				}
				return &XorExpr{left, right}, nil
			}
		}
	}

	// Parse "OR" expression, ensuring we're getting the outermost '|' not enclosed in parentheses
	for i, depth := 0, 0; i < len(expr); i++ {
		switch expr[i] {
		case '(':
			depth++
		case ')':
			depth--
		case '|':
			if depth == 0 {
				left, err := parseRecursively(expr[:i])
				if err != nil {
					return nil, err
				}
				right, err := parseRecursively(expr[i+1:])
				if err != nil {
					return nil, err
				}
				return &OrExpr{left, right}, nil
			}
		}
	}

	// Parse "AND" expression, ensuring we're getting the outermost '&' not enclosed in parentheses
	for i, depth := 0, 0; i < len(expr); i++ {
		switch expr[i] {
		case '(':
			depth++
		case ')':
			depth--
		case '&':
			if depth == 0 {
				left, err := parseRecursively(expr[:i])
				if err != nil {
					return nil, err
				}
				right, err := parseRecursively(expr[i+1:])
				if err != nil {
					return nil, err
				}
				return &AndExpr{left, right}, nil
			}
		}
	}

	// If neither '|', '&', '^', nor '!' are found, it should be a value
	value, err := strconv.Atoi(expr)
	if err != nil {
		return nil, errors.New("invalid expression")
	}
	return &VarExpr{Value: value}, nil
}
