package calculation

import (
	"alexander/caller/model"
	"errors"
)

func extend(calculation *model.Calculation, operator string, value int) {
	calculation.Operations = append(calculation.Operations, model.Operation{Operator: operator, Val: value})
}

func Add(calculation *model.Calculation, value int) {
	extend(calculation, "+", value)
}

func Subtract(calculation *model.Calculation, value int) {
	extend(calculation, "-", value)
}

func Multiply(calculation *model.Calculation, value int) {
	extend(calculation, "*", value)
}

func Divide(calculation *model.Calculation, value int) {
	extend(calculation, "/", value)
}

func Enter(calculation *model.Calculation) (int, error) {
	// Because of presedence, we multiply and divide everything first.
	newList := make([]model.Operation, 0)
	value := calculation.InitialValue
	for _, v := range calculation.Operations {
		switch v.Operator {
		case "+":
			newList = append(newList, model.Operation{Operator: "+", Val: value})
			value = v.Val
		case "-":
			newList = append(newList, model.Operation{Operator: "-", Val: value})
			value = v.Val
		case "*":
			value = value * v.Val
		case "/":
			value = value / v.Val
		default:
			return 0, errors.New("error: unknown operator in expression")
		}
	}

	// Then process addition and subtraction
	for _, v := range newList {
		switch v.Operator {
		case "+":
			value = value + v.Val
		case "-":
			value = value - v.Val
		}
	}

	return value, nil
}
