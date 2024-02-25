package integration

import (
	"errors"
	"os/exec"
	"strings"
	"testing"
)

const (
	usageMsg           = "Usage: main <operand1> <operator> <operand2>"
	invalidOperandMsg  = "invalid operand"
	invalidOperatorMsg = "invalid operator"
	divisionByZeroMsg  = "division by zero is not allowed"
)

func TestCalculator(t *testing.T) {
	errExitStatus := errors.New("exit status 1")
	tests := []struct {
		input    []string
		expected string
		err      error
	}{
		{[]string{"1", "+", "1"}, "1 + 1 = 2.000000", nil},
		{[]string{"5", "-", "3"}, "5 - 3 = 2.000000", nil},
		{[]string{"2", "*", "3"}, "2 * 3 = 6.000000", nil},
		{[]string{"6", "/", "2"}, "6 / 2 = 3.000000", nil},
		{[]string{"6", "/", "0"}, "division by zero is not allowed", errExitStatus},
		// number of arguments
		{[]string{}, usageMsg, errExitStatus},
		{[]string{"1"}, usageMsg, errExitStatus},
		{[]string{"1", "+"}, usageMsg, errExitStatus},
		{[]string{"1", "+", "2"}, "1 + 2 = 3.000000", nil},
		{[]string{"1", "+", "2", "3"}, usageMsg, errExitStatus},
		// invalid arguments
		{[]string{"invalid", "+", "2"}, invalidOperandMsg, errExitStatus},
		{[]string{"1", "invalid", "2"}, invalidOperatorMsg, errExitStatus},
		{[]string{"1", "+", "invalid"}, invalidOperandMsg, errExitStatus},
		// negative operands
		{[]string{"-1", "+", "-1"}, "-1 + -1 = -2.000000", nil},
		// decimal places
		{[]string{"1.5", "+", "1"}, "1.5 + 1 = 2.500000", nil},
		{[]string{"1.55", "+", "1"}, "1.55 + 1 = 2.550000", nil},
		{[]string{"1.555", "+", "1"}, "1.555 + 1 = 2.555000", nil},
		{[]string{"1.5555", "+", "1"}, "1.5555 + 1 = 2.555500", nil},
		{[]string{"1.55555", "+", "1"}, "1.55555 + 1 = 2.555550", nil},
		{[]string{"1.555555", "+", "1"}, "1.555555 + 1 = 2.555555", nil},
		{[]string{"1.5555555", "+", "1"}, "1.5555555 + 1 = 2.555556", nil},
		{[]string{"1.55555555", "+", "1"}, "1.55555555 + 1 = 2.555556", nil},
	}

	for _, test := range tests {
		cmd := exec.Command(builtBinaryPath, test.input...)
		result, err := cmd.CombinedOutput()

		if err != nil && test.err != nil && err.Error() != test.err.Error() {
			t.Errorf("Expected error: %v, got: %v", test.err, err)
		}

		if !strings.Contains(string(result), test.expected) {
			t.Errorf("Expected result contains: %v, got: %v", test.expected, string(result))
		}
	}
}
