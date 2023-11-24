package evaluate_reverse_polish_notation

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestEvalRPN(t *testing.T) {
	testCases := []struct {
		description string
		tokens      []string
		expectation int
	}{
		{"tokens = [\"2\",\"1\",\"+\",\"3\",\"*\"]", []string{"2", "1", "+", "3", "*"}, 9},
		{"tokens = [\"4\",\"13\",\"5\",\"/\",\"+\"]", []string{"4", "13", "5", "/", "+"}, 6},
		{"tokens = [\"10\",\"6\",\"9\",\"3\",\"+\",\"-11\",\"*\",\"/\",\"*\",\"17\",\"+\",\"5\",\"+\"]", []string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"}, 22},
	}

	for _, testCase := range testCases {
		t.Run(testCase.description, func(t *testing.T) {
			require.Equal(t, testCase.expectation, evalRPN(testCase.tokens))
		})
	}
}
