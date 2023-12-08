package gas_station

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCanCompleteCircuit(t *testing.T) {
	testcases := []struct {
		description string
		gas         []int
		cost        []int
		expectation int
	}{
		{"gas = [1,2,3,4,5], cost = [3,4,5,1,2]", []int{1, 2, 3, 4, 5}, []int{3, 4, 5, 1, 2}, 3},
		{"gas = [2,3,4], cost = [3,4,3]", []int{2, 3, 4}, []int{3, 4, 3}, -1},
		{"gas = [5,1,2,3,4], cost = [4,4,1,5,1]", []int{5, 1, 2, 3, 4}, []int{4, 4, 1, 5, 1}, 4},
	}

	for _, tc := range testcases {
		t.Run(tc.description, func(t *testing.T) {
			require.Equal(t, tc.expectation, canCompleteCircuit(tc.gas, tc.cost))
		})
	}
}
