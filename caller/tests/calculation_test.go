package calculation

import (
	"alexander/caller/calculation"
	"alexander/caller/model"
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
	expectedResult := 5

	if err != nil || result != expectedResult {
		t.Errorf("Calculation didn't evaluate correctly. Expected %d, but got %d", expectedResult, result)
	}
}

func TestEnterEmptyOperationsNegativeInitialValue(t *testing.T) {
	newCalculation := defaultEmptyCalculation(-5)
	result, err := calculation.Enter(&newCalculation)
	expectedResult := -5

	if err != nil || result != expectedResult {
		t.Errorf("Calculation didn't evaluate correctly. Expected %d, but got %d", expectedResult, result)
	}
}

func TestEnterEmptyOperationsZeroInitialValue(t *testing.T) {
	newCalculation := defaultEmptyCalculation(0)
	result, err := calculation.Enter(&newCalculation)
	expectedResult := 0

	if err != nil || result != expectedResult {
		t.Errorf("Calculation didn't evaluate correctly. Expected %d, but got %d", expectedResult, result)
	}
}

// CALCULATE - ADD
func TestEnterAddTwoPositiveValues(t *testing.T) {
	newCalculation := defaultEmptyCalculation(5)
	newCalculation.Operations = []model.Operation{{Operator: "+", Val: 5}}

	result, err := calculation.Enter(&newCalculation)
	expectedResult := 10

	if err != nil || result != expectedResult {
		t.Errorf("Calculation didn't evaluate correctly. Expected %d, but got %d", expectedResult, result)
	}
}

func TestEnterAddTwoNegativeValues(t *testing.T) {
	newCalculation := defaultEmptyCalculation(-5)
	newCalculation.Operations = []model.Operation{{Operator: "+", Val: -5}}

	result, err := calculation.Enter(&newCalculation)
	expectedResult := -10

	if err != nil || result != expectedResult {
		t.Errorf("Calculation didn't evaluate correctly. Expected %d, but got %d", expectedResult, result)
	}
}

func TestEnterAddNegativeAndPositiveValues(t *testing.T) {
	newCalculation := defaultEmptyCalculation(-5)
	newCalculation.Operations = []model.Operation{{Operator: "+", Val: 5}}

	result, err := calculation.Enter(&newCalculation)
	expectedResult := -0

	if err != nil || result != expectedResult {
		t.Errorf("Calculation didn't evaluate correctly. Expected %d, but got %d", expectedResult, result)
	}
}


