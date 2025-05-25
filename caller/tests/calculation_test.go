package calculation

import (
	"alexander/main/calculation"
	"alexander/main/store"
	"math"
	"slices"
	"testing"

	"github.com/google/uuid"
)

func operatorEqual(op1 store.Operation, op2 store.Operation) bool {
	return op1.Operator == op2.Operator && op1.Val == op2.Val
}

func defaultEmptyCalculation(initialValue int) store.Calculation {
	newUuid := uuid.New()
	return store.Calculation{Id: newUuid.String(), InitialValue: initialValue, Operations: make([]store.Operation, 0)}
}

// ADDITION
func TestAdditionToEmptyOperations(t *testing.T) {
	newCalculation := defaultEmptyCalculation(5)
	calculation.Add(&newCalculation, 5)

	expectedOperations := []store.Operation{{Operator: "+", Val: 5}}
	if newCalculation.InitialValue != 5 || !slices.EqualFunc(newCalculation.Operations, expectedOperations, operatorEqual) {
		t.Error()
	}
}

func TestAdditionToNonEmptyOperations(t *testing.T) {
	newCalculation := defaultEmptyCalculation(5)
	newCalculation.Operations = append(newCalculation.Operations, store.Operation{Operator: "+", Val: 5})
	calculation.Add(&newCalculation, 5)

	expectedOperations := []store.Operation{{Operator: "+", Val: 5}, {Operator: "+", Val: 5}}
	if newCalculation.InitialValue != 5 || !slices.EqualFunc(newCalculation.Operations, expectedOperations, operatorEqual) {
		t.Error()
	}
}

// SUBTRACTION
func TestSubtractionToEmptyOperations(t *testing.T) {
	newCalculation := defaultEmptyCalculation(5)
	calculation.Subtract(&newCalculation, 5)

	expectedOperations := []store.Operation{{Operator: "-", Val: 5}}
	if newCalculation.InitialValue != 5 || !slices.EqualFunc(newCalculation.Operations, expectedOperations, operatorEqual) {
		t.Error()
	}
}

func TestSubtractionToNonEmptyOperations(t *testing.T) {
	newCalculation := defaultEmptyCalculation(5)
	newCalculation.Operations = append(newCalculation.Operations, store.Operation{Operator: "-", Val: 5})
	calculation.Subtract(&newCalculation, 5)

	expectedOperations := []store.Operation{{Operator: "-", Val: 5}, {Operator: "-", Val: 5}}
	if newCalculation.InitialValue != 5 || !slices.EqualFunc(newCalculation.Operations, expectedOperations, operatorEqual) {
		t.Error()
	}
}

// MULTIPLICATION
func TestMultiplocationToEmptyOperations(t *testing.T) {
	newCalculation := defaultEmptyCalculation(5)
	calculation.Multiply(&newCalculation, 5)

	expectedOperations := []store.Operation{{Operator: "*", Val: 5}}
	if newCalculation.InitialValue != 5 || !slices.EqualFunc(newCalculation.Operations, expectedOperations, operatorEqual) {
		t.Error()
	}
}

func TestMultiplicationToNonEmptyOperations(t *testing.T) {
	newCalculation := defaultEmptyCalculation(5)
	newCalculation.Operations = append(newCalculation.Operations, store.Operation{Operator: "*", Val: 5})
	calculation.Multiply(&newCalculation, 5)

	expectedOperations := []store.Operation{{Operator: "*", Val: 5}, {Operator: "*", Val: 5}}
	if newCalculation.InitialValue != 5 || !slices.EqualFunc(newCalculation.Operations, expectedOperations, operatorEqual) {
		t.Error()
	}
}

// DIVISION
func TestDivisionToEmptyOperations(t *testing.T) {
	newCalculation := defaultEmptyCalculation(5)
	calculation.Divide(&newCalculation, 5)

	expectedOperations := []store.Operation{{Operator: "/", Val: 5}}
	if newCalculation.InitialValue != 5 || !slices.EqualFunc(newCalculation.Operations, expectedOperations, operatorEqual) {
		t.Error()
	}
}

func TestDivisionToNonEmptyOperations(t *testing.T) {
	newCalculation := defaultEmptyCalculation(5)
	newCalculation.Operations = append(newCalculation.Operations, store.Operation{Operator: "/", Val: 5})
	calculation.Divide(&newCalculation, 5)

	expectedOperations := []store.Operation{{Operator: "/", Val: 5}, {Operator: "/", Val: 5}}
	if newCalculation.InitialValue != 5 || !slices.EqualFunc(newCalculation.Operations, expectedOperations, operatorEqual) {
		t.Error()
	}
}

func TestDivisionByZero(t *testing.T) {
	newCalculation := defaultEmptyCalculation(5)
	if err := calculation.Divide(&newCalculation, 0); err == nil {
		t.Error()
	}
}

// MIXED OPERATIONS
func TestAllOpsToEmptyOperations(t *testing.T) {
	newCalculation := defaultEmptyCalculation(5)
	calculation.Add(&newCalculation, 5)
	calculation.Subtract(&newCalculation, 5)
	calculation.Multiply(&newCalculation, 5)
	calculation.Divide(&newCalculation, 5)

	expectedOperations := []store.Operation{{Operator: "+", Val: 5}, {Operator: "-", Val: 5}, {Operator: "*", Val: 5}, {Operator: "/", Val: 5}}
	if newCalculation.InitialValue != 5 || !slices.EqualFunc(newCalculation.Operations, expectedOperations, operatorEqual) {
		t.Error()
	}
}

func TestAllOpsToNonEmptyOperations(t *testing.T) {
	newCalculation := defaultEmptyCalculation(5)
	newCalculation.Operations = append(newCalculation.Operations, store.Operation{Operator: "+", Val: 5}, store.Operation{Operator: "-", Val: 5}, store.Operation{Operator: "*", Val: 5}, store.Operation{Operator: "/", Val: 5})

	calculation.Add(&newCalculation, 5)
	calculation.Subtract(&newCalculation, 5)
	calculation.Multiply(&newCalculation, 5)
	calculation.Divide(&newCalculation, 5)

	expectedOperations := []store.Operation{{Operator: "+", Val: 5}, {Operator: "-", Val: 5}, {Operator: "*", Val: 5}, {Operator: "/", Val: 5}, {Operator: "+", Val: 5}, {Operator: "-", Val: 5}, {Operator: "*", Val: 5}, {Operator: "/", Val: 5}}
	if newCalculation.InitialValue != 5 || !slices.EqualFunc(newCalculation.Operations, expectedOperations, operatorEqual) {
		t.Error()
	}
}

// EVALUATIONS
var calculateTests = []struct {
	name           string
	calculation    store.Calculation
	expectedResult float64
}{
	// Only initial value
	{"TestEmptyCalculationPositiveInitialValue", store.Calculation{InitialValue: 5, Operations: []store.Operation{}}, 5},
	{"TestEmptyCalculationNegativeInitialValue", store.Calculation{InitialValue: -5, Operations: []store.Operation{}}, -5},
	{"TestEmptyCalculationZeroInitialValue", store.Calculation{InitialValue: 0, Operations: []store.Operation{}}, 0},

	// Only additions
	{"TestAdd2PositiveIntegers", store.Calculation{InitialValue: 5, Operations: []store.Operation{{Operator: "+", Val: 3}}}, 8},
	{"TestAdd2NegativeIntegers", store.Calculation{InitialValue: -5, Operations: []store.Operation{{Operator: "+", Val: -3}}}, -8},
	{"TestAdd2Negative&PositiveIntegers", store.Calculation{InitialValue: -5, Operations: []store.Operation{{Operator: "+", Val: 3}}}, -2},
	{"TestAdd2Positive&NegativeIntegers", store.Calculation{InitialValue: 5, Operations: []store.Operation{{Operator: "+", Val: -3}}}, 2},

	// Only subtractions
	{"TestSub2PositiveIntegers", store.Calculation{InitialValue: 5, Operations: []store.Operation{{Operator: "-", Val: 3}}}, 2},
	{"TestSub2NegativeIntegers", store.Calculation{InitialValue: -5, Operations: []store.Operation{{Operator: "-", Val: -3}}}, -2},
	{"TestAdd2Negative&PositiveIntegers", store.Calculation{InitialValue: -5, Operations: []store.Operation{{Operator: "-", Val: 3}}}, -8},
	{"TestSub2Positive&NegativeIntegers", store.Calculation{InitialValue: 5, Operations: []store.Operation{{Operator: "-", Val: -3}}}, 8},

	// Mix addition and subtractions
	{"TestAdd2PositiveIntegersThenSubtract", store.Calculation{InitialValue: 5, Operations: []store.Operation{{Operator: "+", Val: 3}, {Operator: "-", Val: 4}}}, 4},
	{"TestAdd2NegativeIntegersThenSubtract", store.Calculation{InitialValue: -5, Operations: []store.Operation{{Operator: "+", Val: -3}, {Operator: "-", Val: 4}}}, -12},
	{"TestSub2PositiveIntegersThenAdd", store.Calculation{InitialValue: 5, Operations: []store.Operation{{Operator: "-", Val: 3}, {Operator: "+", Val: 4}}}, 6},
	{"TestSub2NegativeIntegersThenAdd", store.Calculation{InitialValue: -5, Operations: []store.Operation{{Operator: "-", Val: -3}, {Operator: "+", Val: 4}}}, 2},

	// Only Multiplication
	{"TestMult2PositiveIntegers", store.Calculation{InitialValue: 5, Operations: []store.Operation{{Operator: "*", Val: 3}}}, 15},
	{"TestMult2NegativeIntegers", store.Calculation{InitialValue: -5, Operations: []store.Operation{{Operator: "*", Val: -3}}}, 15},
	{"TestMult2Negative&PositiveIntegers", store.Calculation{InitialValue: -5, Operations: []store.Operation{{Operator: "*", Val: 3}}}, -15},
	{"TestMult2Positive&NegativeIntegers", store.Calculation{InitialValue: 5, Operations: []store.Operation{{Operator: "*", Val: -3}}}, -15},

	// Only Division
	{"TestDiv2PositiveIntegers", store.Calculation{InitialValue: 5, Operations: []store.Operation{{Operator: "/", Val: 3}}}, 1.666666667},
	{"TestDiv2NegativeIntegers", store.Calculation{InitialValue: -5, Operations: []store.Operation{{Operator: "/", Val: -3}}}, 1.666666667},
	{"TestDiv2Negative&PositiveIntegers", store.Calculation{InitialValue: -5, Operations: []store.Operation{{Operator: "/", Val: 3}}}, -1.666666667},
	{"TestDiv2Positive&NegativeIntegers", store.Calculation{InitialValue: 5, Operations: []store.Operation{{Operator: "/", Val: -3}}}, -1.666666667},

	// Mix presedence level
	{"TestMultHasPresedenceOverAdditionLeft", store.Calculation{InitialValue: 5, Operations: []store.Operation{{Operator: "*", Val: 3}, {Operator: "+", Val: 5}}}, 20},
	{"TestMultHasPresedenceOverAdditionRight", store.Calculation{InitialValue: 5, Operations: []store.Operation{{Operator: "+", Val: 3}, {Operator: "*", Val: 5}}}, 20},
	{"TestMultHasPresedenceOverSubtractionLeft", store.Calculation{InitialValue: 5, Operations: []store.Operation{{Operator: "*", Val: 3}, {Operator: "-", Val: 5}}}, 10},
	{"TestMultHasPresedenceOverSubtractionRight", store.Calculation{InitialValue: 5, Operations: []store.Operation{{Operator: "-", Val: 3}, {Operator: "*", Val: 5}}}, -10},

	{"TestDivisionHasPresedenceOverAdditionLeft", store.Calculation{InitialValue: 5, Operations: []store.Operation{{Operator: "/", Val: 3}, {Operator: "+", Val: 5}}}, 6.666666667},
	{"TestDivisionHasPresedenceOverAdditionRight", store.Calculation{InitialValue: 5, Operations: []store.Operation{{Operator: "+", Val: 3}, {Operator: "/", Val: 5}}}, 5.6},
	{"TestDivisionHasPresedenceOverSubtractionLeft", store.Calculation{InitialValue: 5, Operations: []store.Operation{{Operator: "/", Val: 3}, {Operator: "-", Val: 5}}}, -3.33333333},
	{"TestDivisionHasPresedenceOverSubtractionRight", store.Calculation{InitialValue: 5, Operations: []store.Operation{{Operator: "-", Val: 3}, {Operator: "/", Val: 5}}}, 4.4},

	{"TestMultDoesNotHavePresedenceOverDivision", store.Calculation{InitialValue: 5, Operations: []store.Operation{{Operator: "/", Val: 3}, {Operator: "*", Val: 5}}}, 8.3333333},
	{"TestDivisionDoesNotHavePresedenceOverMult", store.Calculation{InitialValue: 5, Operations: []store.Operation{{Operator: "*", Val: 3}, {Operator: "/", Val: 5}}}, 3},

	{"TestComplicatedOneGroupWithpresedence", store.Calculation{InitialValue: 5, Operations: []store.Operation{{Operator: "+", Val: 3}, {Operator: "*", Val: 5}, {Operator: "/", Val: 2}, {Operator: "-", Val: 1}}}, 11.5},
	{"TestComplicatedTwoGroupWithpresedence", store.Calculation{InitialValue: 5, Operations: []store.Operation{{Operator: "*", Val: 3}, {Operator: "+", Val: 5}, {Operator: "/", Val: 2}, {Operator: "-", Val: 1}}}, 16.5},
}

// Credit to: https://medium.com/pragmatic-programmers/testing-floating-point-numbers-in-go-9872fe6de17f
func floatEq(arg1 float64, arg2 float64) bool {
	if arg1 == arg2 {
		return true
	}
	var e = 0.0000001 // Epsilon
	d := math.Abs(arg1 - arg2)

	if arg2 == 0 {
		return d < e
	}
	return (d / math.Abs(arg2)) < e

}

func TestEnter(t *testing.T) {
	for _, tt := range calculateTests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := calculation.Enter(&tt.calculation)
			if err != nil || !floatEq(result, tt.expectedResult) {
				t.Errorf("Calculation didn't evaluate correctly. Expected %f, but got %f", tt.expectedResult, result)
			}
		})
	}
}
