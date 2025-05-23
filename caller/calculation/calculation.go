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

func Enter(calculation *model.Calculation) (float64, error) {

	newList := make([]float64, 0)
	var value float64 = float64(calculation.InitialValue)

	// 5 + 3 = 3

	// Different thought. Just create a list of values we just need to add together
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
	// Add the remaining value
	newList = append(newList, value)

	// Then process addition and subtraction
	var total float64 = 0
	for _, v := range newList {
		total = total + v
	}

	return total, nil
}
