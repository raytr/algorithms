package two_sum_ii_input_array_is_sorted

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTwoSum(t *testing.T) {
	testCases := []struct {
		name        string
		numbers     []int
		target      int
		expectation []int
	}{
		{"numbers = [2,7,11,15], target = 9", []int{2, 7, 11, 15}, 9, []int{1, 2}},
		{"numbers = [2,3,4], target = 6", []int{2, 3, 4}, 6, []int{1, 3}},
		{"numbers = [-1,0], target = -1", []int{-1, 0}, -1, []int{1, 2}},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			got := twoSum(tt.numbers, tt.target)
			assert.Equal(t, tt.expectation, got)
		})
	}
}
