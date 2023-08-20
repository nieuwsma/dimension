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

func parseRecursively(expr string) (Expr, error) {
	if idx := strings.Index(expr, "|"); idx != -1 {
		left, err := parseRecursively(expr[:idx])
		if err != nil {
			return nil, err
		}

		right, err := parseRecursively(expr[idx+1:])
		if err != nil {
			return nil, err
		}

		return &OrExpr{left, right}, nil
	}

	if idx := strings.Index(expr, "&"); idx != -1 {
		left, err := parseRecursively(expr[:idx])
		if err != nil {
			return nil, err
		}

		right, err := parseRecursively(expr[idx+1:])
		if err != nil {
			return nil, err
		}

		return &AndExpr{left, right}, nil
	}

	if strings.HasPrefix(expr, "(") && strings.HasSuffix(expr, ")") {
		return parseRecursively(expr[1 : len(expr)-1])
	}

	value, err := strconv.Atoi(expr)
	if err != nil {
		return nil, errors.New("invalid expression")
	}

	return &VarExpr{Value: value}, nil
}
