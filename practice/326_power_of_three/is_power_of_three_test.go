package power_of_three

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestIsPowerOfThree(t *testing.T) {
	testCases := []struct {
		name        string
		n           int
		expectation bool
	}{
		{"n = 27", 27, true},
		{"n = 0", 0, false},
		{"n = -1", -1, false},
		{"n = 1", 1, true},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			require.Equal(t, tt.expectation, isPowerOfThree(tt.n))
		})
	}
}
