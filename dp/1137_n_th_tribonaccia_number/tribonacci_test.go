package n_th_tribonaccia_number

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTribonacci(t *testing.T) {
	testcases := []struct {
		description string
		n           int
		want        int
	}{
		{"n = 4", 4, 4},
		{"n = 25", 25, 1389537},
	}

	for _, tc := range testcases {
		t.Run(tc.description, func(t *testing.T) {
			assert.Equal(t, tc.want, tribonacci(tc.n))
		})
	}
}
