package calculation

import (
	"alexander/main/store"
	"errors"
)

func extend(calculation *store.Calculation, operator string, value int) {
	calculation.Operations = append(calculation.Operations, store.Operation{Operator: operator, Val: value})
}

func Add(calculation *store.Calculation, value int) error {
	extend(calculation, "+", value)
	return nil
}

func Subtract(calculation *store.Calculation, value int) error {
	extend(calculation, "-", value)
	return nil
}

func Multiply(calculation *store.Calculation, value int) error {
	extend(calculation, "*", value)
	return nil
}

type DivisionByZeroError struct {
}

func (e DivisionByZeroError) Error() string {
	return "division by zero is disallowed"
}

func Divide(calculation *store.Calculation, value int) error {
	if value == 0 {
		return DivisionByZeroError{}
	}
	extend(calculation, "/", value)
	return nil
}

type UndoError struct {
}

func (e UndoError) Error() string {
	return "no operations to undo"
}

func Undo(calculation *store.Calculation) error {
	numOperators := len(calculation.Operations)
	if numOperators == 0 {
		return UndoError{}
	}
	calculation.Operations = calculation.Operations[:numOperators-1]
	return nil
}

func Enter(calculation *store.Calculation) (float64, error) {

	newList := make([]float64, 0)
	var value float64 = float64(calculation.InitialValue)

	// Evaluate all subexpressions with division and multiplication first
	for _, v := range calculation.Operations {
		switch v.Operator {
		case "+":
			newList = append(newList, value)
			value = float64(v.Val)
		case "-":
			newList = append(newList, value)
			value = float64(-v.Val)
		case "*":
			value = value * float64(v.Val)
		case "/":
			value = value / float64(v.Val)
		default:
			return 0, errors.New("error: unknown operator in expression")
		}
	}
	newList = append(newList, value)

	// Then process addition and subtraction
	var total float64 = 0
	for _, v := range newList {
		total = total + v
	}

	return total, nil
}
