package calculation

import (
	"alexander/caller/store"
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

func Divide(calculation *store.Calculation, value int) error {
	if value == 0 {
		return errors.New("error: division by zero is disallowed")
	}
	extend(calculation, "/", value)
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
