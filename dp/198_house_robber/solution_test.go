package house_robber

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSolution(t *testing.T) {
	testcases := []struct {
		description string
		nums        []int
		expectation int
	}{
		{"nums = [1,2,3,1]", []int{1, 2, 3, 1}, 4},
		{"nums = [2,7,9,3,1]", []int{2, 7, 9, 3, 1}, 12},
	}

	for _, tc := range testcases {
		t.Run(tc.description, func(t *testing.T) {
			require.Equal(t, tc.expectation, rob(tc.nums))
		})
	}
}
