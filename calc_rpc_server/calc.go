package main

import "errors"

type Operation func(Number1, Number2 float64) float64

func Add(Number1, Number2 float64) float64 {
	return Number1 + Number2
}

func Sub(Number1, Number2 float64) float64 {
	return Number1 - Number2
}

func Mul(Number1, Number2 float64) float64 {
	return Number1 * Number2
}

func Div(Number1, Number2 float64) float64 {
	return Number1 / Number2
}

var Operators = map[string]Operation {
	"+": Add,
	"-": Sub,
	"*": Mul,
	"/": Div,
}

func CreateOperation(operator string) (Operation, error) {
	var oper Operation
	if oper, ok := Operators[operator]; ok {
		return oper, nil
	}
	return oper, errors.New("operator not found")
}