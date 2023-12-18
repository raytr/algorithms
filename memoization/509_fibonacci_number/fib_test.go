package fibonacci_number

/*
	problem : https://leetcode.com/problems/fibonacci-number/
*/

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFib(t *testing.T) {
	testcases := []struct {
		description string
		n           int
		want        int
	}{
		{"n = 2", 2, 1},
		{"n = 3", 3, 2},
		{"n = 4", 4, 3},
		{"n = 30", 30, 832040},
		{"n = 1000", 1000, 817770325994397771},
	}

	for _, tc := range testcases {
		t.Run(tc.description, func(t *testing.T) {
			require.Equal(t, tc.want, fib(tc.n))
		})
	}
}
