package calculator

import (
	"errors"
	"testing"
)

func TestCalculate(t *testing.T) {
	tests := []struct {
		operand1 float64
		operator string
		operand2 float64
		expected float64
		err      error
	}{
		{1, "+", 1, 2, nil},
		{5, "-", 3, 2, nil},
		{2, "*", 3, 6, nil},
		{6, "/", 2, 3, nil},
		{6, "/", 0, 0, errors.New("division by zero is not allowed")},
		{1, "%", 1, 0, errors.New("invalid operator")},

		// negative lhs
		{-1, "+", 1, 0, nil},
		{-5, "-", 3, -8, nil},
		{-2, "*", 3, -6, nil},
		{-6, "/", 2, -3, nil},

		// negative rhs
		{1, "+", -1, 0, nil},
		{5, "-", -3, 8, nil},
		{2, "*", -3, -6, nil},
		{6, "/", -2, -3, nil},

		// negative lhs and rhs
		{-1, "+", -1, -2, nil},
		{-5, "-", -3, -2, nil},
		{-2, "*", -3, 6, nil},
		{-6, "/", -2, 3, nil},

		// decimal places
		{1.5, "+", 1, 2.5, nil},
		{1.55, "+", 1, 2.55, nil},
		{1.555, "+", 1, 2.5549999999999997, nil},  // TODO: consider rounding
		{1.5555, "+", 1, 2.5555000000000003, nil}, // TODO: consider rounding

		// negative decimal places
		{-1.5, "+", 1, -0.5, nil},
		{-1.55, "+", 1, -0.55, nil},
		{-1.555, "+", 1, -0.5549999999999999, nil},  // TODO: consider rounding
		{-1.5555, "+", 1, -0.5555000000000001, nil}, // TODO: consider rounding

	}

	for _, test := range tests {
		result, err := Calculate(test.operand1, test.operand2, test.operator)

		if err != nil && err.Error() != test.err.Error() {
			t.Errorf("Expected error: %v, got: %v", test.err, err)
		}

		if result != test.expected {
			t.Errorf("Expected result: %v, got: %v", test.expected, result)
		}
	}
}

func TestParseOperand(t *testing.T) {
	tests := []struct {
		input    string
		expected float64
		err      error
	}{
		{"1", 1, nil},
		{"3.14", 3.14, nil},
		{"abc", 0, errors.New("invalid operand")},

		// negative
		{"-3.14", -3.14, nil},

		// alphanumeric
		{"1abc", 0, errors.New("invalid operand")},
		{"abc1", 0, errors.New("invalid operand")},

		// whitespace
		{" 1", 0, errors.New("invalid operand")},
	}

	for _, test := range tests {
		result, err := ParseOperand(test.input)

		if err != nil && err.Error() != test.err.Error() {
			t.Errorf("Expected error: %v, got: %v", test.err, err)
		}

		if result != test.expected {
			t.Errorf("Expected result: %v, got: %v", test.expected, result)
		}
	}
}
