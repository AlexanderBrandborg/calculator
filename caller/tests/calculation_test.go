package calculation

import (
	"alexander/caller/calculation"
	"alexander/caller/model"
	"log"
	"math"
	"slices"
	"testing"

	"github.com/google/uuid"
)

func operatorEqual(op1 model.Operation, op2 model.Operation) bool {
	return op1.Operator == op2.Operator && op1.Val == op2.Val
}

func defaultEmptyCalculation(initialValue int) model.Calculation {
	newUuid := uuid.New()
	return model.Calculation{Id: newUuid.String(), InitialValue: initialValue, Operations: make([]model.Operation, 0)}
}

// ADDITION
func TestAdditionToEmptyOperations(t *testing.T) {
	newCalculation := defaultEmptyCalculation(5)
	calculation.Add(&newCalculation, 5)

	expectedOperations := []model.Operation{{Operator: "+", Val: 5}}
	if newCalculation.InitialValue != 5 || !slices.EqualFunc(newCalculation.Operations, expectedOperations, operatorEqual) {
		t.Error()
	}
}

func TestAdditionToNonEmptyOperations(t *testing.T) {
	newCalculation := defaultEmptyCalculation(5)
	newCalculation.Operations = append(newCalculation.Operations, model.Operation{Operator: "+", Val: 5})
	calculation.Add(&newCalculation, 5)

	expectedOperations := []model.Operation{{Operator: "+", Val: 5}, {Operator: "+", Val: 5}}
	if newCalculation.InitialValue != 5 || !slices.EqualFunc(newCalculation.Operations, expectedOperations, operatorEqual) {
		t.Error()
	}
}

// SUBTRACTION
func TestSubtractionToEmptyOperations(t *testing.T) {
	newCalculation := defaultEmptyCalculation(5)
	calculation.Subtract(&newCalculation, 5)

	expectedOperations := []model.Operation{{Operator: "-", Val: 5}}
	if newCalculation.InitialValue != 5 || !slices.EqualFunc(newCalculation.Operations, expectedOperations, operatorEqual) {
		t.Error()
	}
}

func TestSubtractionToNonEmptyOperations(t *testing.T) {
	newCalculation := defaultEmptyCalculation(5)
	newCalculation.Operations = append(newCalculation.Operations, model.Operation{Operator: "-", Val: 5})
	calculation.Subtract(&newCalculation, 5)

	expectedOperations := []model.Operation{{Operator: "-", Val: 5}, {Operator: "-", Val: 5}}
	if newCalculation.InitialValue != 5 || !slices.EqualFunc(newCalculation.Operations, expectedOperations, operatorEqual) {
		t.Error()
	}
}

// MULTIPLY
func TestMultiplocationToEmptyOperations(t *testing.T) {
	newCalculation := defaultEmptyCalculation(5)
	calculation.Multiply(&newCalculation, 5)

	expectedOperations := []model.Operation{{Operator: "*", Val: 5}}
	if newCalculation.InitialValue != 5 || !slices.EqualFunc(newCalculation.Operations, expectedOperations, operatorEqual) {
		t.Error()
	}
}

func TestMultiplicationToNonEmptyOperations(t *testing.T) {
	newCalculation := defaultEmptyCalculation(5)
	newCalculation.Operations = append(newCalculation.Operations, model.Operation{Operator: "*", Val: 5})
	calculation.Multiply(&newCalculation, 5)

	expectedOperations := []model.Operation{{Operator: "*", Val: 5}, {Operator: "*", Val: 5}}
	if newCalculation.InitialValue != 5 || !slices.EqualFunc(newCalculation.Operations, expectedOperations, operatorEqual) {
		t.Error()
	}
}

// DIVIDE
func TestDivisionToEmptyOperations(t *testing.T) {
	newCalculation := defaultEmptyCalculation(5)
	calculation.Divide(&newCalculation, 5)

	expectedOperations := []model.Operation{{Operator: "/", Val: 5}}
	if newCalculation.InitialValue != 5 || !slices.EqualFunc(newCalculation.Operations, expectedOperations, operatorEqual) {
		t.Error()
	}
}

func TestDivisionToNonEmptyOperations(t *testing.T) {
	newCalculation := defaultEmptyCalculation(5)
	newCalculation.Operations = append(newCalculation.Operations, model.Operation{Operator: "/", Val: 5})
	calculation.Divide(&newCalculation, 5)

	expectedOperations := []model.Operation{{Operator: "/", Val: 5}, {Operator: "/", Val: 5}}
	if newCalculation.InitialValue != 5 || !slices.EqualFunc(newCalculation.Operations, expectedOperations, operatorEqual) {
		t.Error()
	}
}

// MIXED
func TestAllOpsToEmptyOperations(t *testing.T) {
	newCalculation := defaultEmptyCalculation(5)
	calculation.Add(&newCalculation, 5)
	calculation.Subtract(&newCalculation, 5)
	calculation.Multiply(&newCalculation, 5)
	calculation.Divide(&newCalculation, 5)

	expectedOperations := []model.Operation{{Operator: "+", Val: 5}, {Operator: "-", Val: 5}, {Operator: "*", Val: 5}, {Operator: "/", Val: 5}}
	if newCalculation.InitialValue != 5 || !slices.EqualFunc(newCalculation.Operations, expectedOperations, operatorEqual) {
		t.Error()
	}
}

func TestAllOpsToNonEmptyOperations(t *testing.T) {
	newCalculation := defaultEmptyCalculation(5)
	newCalculation.Operations = append(newCalculation.Operations, model.Operation{Operator: "+", Val: 5}, model.Operation{Operator: "-", Val: 5}, model.Operation{Operator: "*", Val: 5}, model.Operation{Operator: "/", Val: 5})

	calculation.Add(&newCalculation, 5)
	calculation.Subtract(&newCalculation, 5)
	calculation.Multiply(&newCalculation, 5)
	calculation.Divide(&newCalculation, 5)

	expectedOperations := []model.Operation{{Operator: "+", Val: 5}, {Operator: "-", Val: 5}, {Operator: "*", Val: 5}, {Operator: "/", Val: 5}, {Operator: "+", Val: 5}, {Operator: "-", Val: 5}, {Operator: "*", Val: 5}, {Operator: "/", Val: 5}}
	if newCalculation.InitialValue != 5 || !slices.EqualFunc(newCalculation.Operations, expectedOperations, operatorEqual) {
		t.Error()
	}
}

// CALCULATE
func TestEnterEmptyOperationsPositiveInitialValue(t *testing.T) {
	newCalculation := defaultEmptyCalculation(5)
	result, err := calculation.Enter(&newCalculation)
	var expectedResult float64 = 5

	if err != nil || result != expectedResult {
		t.Errorf("Calculation didn't evaluate correctly. Expected %f, but got %f", expectedResult, result)
	}
}

func TestEnterEmptyOperationsNegativeInitialValue(t *testing.T) {
	newCalculation := defaultEmptyCalculation(-5)
	result, err := calculation.Enter(&newCalculation)
	var expectedResult float64 = -5

	if err != nil || result != expectedResult {
		t.Errorf("Calculation didn't evaluate correctly. Expected %f, but got %f", expectedResult, result)
	}
}

func TestEnterEmptyOperationsZeroInitialValue(t *testing.T) {
	newCalculation := defaultEmptyCalculation(0)
	result, err := calculation.Enter(&newCalculation)
	var expectedResult float64 = 0

	if err != nil || result != expectedResult {
		t.Errorf("Calculation didn't evaluate correctly. Expected %f, but got %f", expectedResult, result)
	}
}

// CALCULATE - ADD
func TestEnterAddTwoPositiveValues(t *testing.T) {
	newCalculation := defaultEmptyCalculation(5)
	newCalculation.Operations = []model.Operation{{Operator: "+", Val: 5}}

	result, err := calculation.Enter(&newCalculation)
	var expectedResult float64 = 10

	if err != nil || result != expectedResult {
		t.Errorf("Calculation didn't evaluate correctly. Expected %f, but got %f", expectedResult, result)
	}
}

func TestEnterAddTwoNegativeValues(t *testing.T) {
	newCalculation := defaultEmptyCalculation(-5)
	newCalculation.Operations = []model.Operation{{Operator: "+", Val: -5}}

	result, err := calculation.Enter(&newCalculation)
	var expectedResult float64 = -10

	if err != nil || result != expectedResult {
		t.Errorf("Calculation didn't evaluate correctly. Expected %f, but got %f", expectedResult, result)
	}
}

func TestEnterAddNegativeAndPositiveValues(t *testing.T) {
	newCalculation := defaultEmptyCalculation(-5)
	newCalculation.Operations = []model.Operation{{Operator: "+", Val: 5}}

	result, err := calculation.Enter(&newCalculation)
	var expectedResult float64 = 0

	if err != nil || result != expectedResult {
		t.Errorf("Calculation didn't evaluate correctly. Expected %f, but got %f", expectedResult, result)
	}
}

var calculateTests = []struct {
	name           string
	calculation    model.Calculation
	expectedResult float64
}{
	// Only initial value
	{"TestEmptyCalculationPositiveInitialValue", model.Calculation{InitialValue: 5, Operations: []model.Operation{}}, 5},
	{"TestEmptyCalculationNegativeInitialValue", model.Calculation{InitialValue: -5, Operations: []model.Operation{}}, -5},
	{"TestEmptyCalculationZeroInitialValue", model.Calculation{InitialValue: 0, Operations: []model.Operation{}}, 0},

	// Only additions
	{"TestAdd2PositiveIntegers", model.Calculation{InitialValue: 5, Operations: []model.Operation{{Operator: "+", Val: 3}}}, 8},
	{"TestAdd2NegativeIntegers", model.Calculation{InitialValue: -5, Operations: []model.Operation{{Operator: "+", Val: -3}}}, -8},
	{"TestAdd2Negative&PositiveIntegers", model.Calculation{InitialValue: -5, Operations: []model.Operation{{Operator: "+", Val: 3}}}, -2},
	{"TestAdd2Positive&NegativeIntegers", model.Calculation{InitialValue: 5, Operations: []model.Operation{{Operator: "+", Val: -3}}}, 2},

	// Only subtractions
	{"TestSub2PositiveIntegers", model.Calculation{InitialValue: 5, Operations: []model.Operation{{Operator: "-", Val: 3}}}, 2},
	{"TestSub2NegativeIntegers", model.Calculation{InitialValue: -5, Operations: []model.Operation{{Operator: "-", Val: -3}}}, -2},
	{"TestAdd2Negative&PositiveIntegers", model.Calculation{InitialValue: -5, Operations: []model.Operation{{Operator: "-", Val: 3}}}, -8},
	{"TestSub2Positive&NegativeIntegers", model.Calculation{InitialValue: 5, Operations: []model.Operation{{Operator: "-", Val: -3}}}, 8},

	// Mix addition and subtractions
	{"TestAdd2PositiveIntegersThenSubtract", model.Calculation{InitialValue: 5, Operations: []model.Operation{{Operator: "+", Val: 3}, {Operator: "-", Val: 4}}}, 4},
	{"TestAdd2NegativeIntegersThenSubtract", model.Calculation{InitialValue: -5, Operations: []model.Operation{{Operator: "+", Val: -3}, {Operator: "-", Val: 4}}}, -12},
	{"TestSub2PositiveIntegersThenAdd", model.Calculation{InitialValue: 5, Operations: []model.Operation{{Operator: "-", Val: 3}, {Operator: "+", Val: 4}}}, 6},
	{"TestSub2NegativeIntegersThenAdd", model.Calculation{InitialValue: -5, Operations: []model.Operation{{Operator: "-", Val: -3}, {Operator: "+", Val: 4}}}, 2},

	// Only Multiplication
	{"TestMult2PositiveIntegers", model.Calculation{InitialValue: 5, Operations: []model.Operation{{Operator: "*", Val: 3}}}, 15},
	{"TestMult2NegativeIntegers", model.Calculation{InitialValue: -5, Operations: []model.Operation{{Operator: "*", Val: -3}}}, 15},
	{"TestMult2Negative&PositiveIntegers", model.Calculation{InitialValue: -5, Operations: []model.Operation{{Operator: "*", Val: 3}}}, -15},
	{"TestMult2Positive&NegativeIntegers", model.Calculation{InitialValue: 5, Operations: []model.Operation{{Operator: "*", Val: -3}}}, -15},

	// Only Division
	//{"TestDivByZero", model.Calculation{InitialValue: 5, Operations: []model.Operation{{Operator: "/", Val: 0}}}, 1.666666667}, // TODO: Handle divide by zero
	{"TestDiv2PositiveIntegers", model.Calculation{InitialValue: 5, Operations: []model.Operation{{Operator: "/", Val: 3}}}, 1.666666667},
	{"TestDiv2NegativeIntegers", model.Calculation{InitialValue: -5, Operations: []model.Operation{{Operator: "/", Val: -3}}}, 1.666666667},
	{"TestDiv2Negative&PositiveIntegers", model.Calculation{InitialValue: -5, Operations: []model.Operation{{Operator: "/", Val: 3}}}, -1.666666667},
	{"TestDiv2Positive&NegativeIntegers", model.Calculation{InitialValue: 5, Operations: []model.Operation{{Operator: "/", Val: -3}}}, -1.666666667},

	// Mix presedence level
	{"TestMultHasPresedenceOverAdditionLeft", model.Calculation{InitialValue: 5, Operations: []model.Operation{{Operator: "*", Val: 3}, {Operator: "+", Val: 5}}}, 20},
	{"TestMultHasPresedenceOverAdditionRight", model.Calculation{InitialValue: 5, Operations: []model.Operation{{Operator: "+", Val: 3}, {Operator: "*", Val: 5}}}, 20},
	{"TestMultHasPresedenceOverSubtractionLeft", model.Calculation{InitialValue: 5, Operations: []model.Operation{{Operator: "*", Val: 3}, {Operator: "-", Val: 5}}}, 10},
	{"TestMultHasPresedenceOverSubtractionRight", model.Calculation{InitialValue: 5, Operations: []model.Operation{{Operator: "-", Val: 3}, {Operator: "*", Val: 5}}}, -10},

	{"TestDivisionHasPresedenceOverAdditionLeft", model.Calculation{InitialValue: 5, Operations: []model.Operation{{Operator: "/", Val: 3}, {Operator: "+", Val: 5}}}, 6.666666667},
	{"TestMultHasPresedenceOverAdditionRight", model.Calculation{InitialValue: 5, Operations: []model.Operation{{Operator: "+", Val: 3}, {Operator: "/", Val: 5}}}, 5.6},
	{"TestMultHasPresedenceOverSubtractionLeft", model.Calculation{InitialValue: 5, Operations: []model.Operation{{Operator: "/", Val: 3}, {Operator: "-", Val: 5}}}, -3.33333333},
	{"TestMultHasPresedenceOverSubtractionRight", model.Calculation{InitialValue: 5, Operations: []model.Operation{{Operator: "-", Val: 3}, {Operator: "/", Val: 5}}}, 4.4},

	// NOTE: Add some more complicated expressions
	{"TestComplicatedOneGroupWithpresedence", model.Calculation{InitialValue: 5, Operations: []model.Operation{{Operator: "+", Val: 3}, {Operator: "*", Val: 5}, {Operator: "/", Val: 2}, {Operator: "-", Val: 1}}}, 11.5},
	{"TestComplicatedTwoGroupWithpresedence", model.Calculation{InitialValue: 5, Operations: []model.Operation{{Operator: "*", Val: 3}, {Operator: "+", Val: 5}, {Operator: "/", Val: 2}, {Operator: "-", Val: 1}}}, 16.5},
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
			if tt.name == "TestSub2PositiveIntegers" {
				log.Print("Give me something to break!")
			}
			result, err := calculation.Enter(&tt.calculation)
			if err != nil || !floatEq(result, tt.expectedResult) {
				t.Errorf("Calculation didn't evaluate correctly. Expected %f, but got %f", tt.expectedResult, result)
			}
		})
	}
}
