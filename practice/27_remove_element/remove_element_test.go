package remove_element

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	testCases := []struct {
		name        string
		nums        []int
		val         int
		expectation int
	}{
		{"nums = [3,2,2,3], val = 3", []int{3, 2, 2, 3}, 3, 2},
		{"nums = [0,1,2,2,3,0,4,2], val = 2", []int{0, 1, 2, 2, 3, 0, 4, 2}, 2, 5}, //expect [0,1,4,0,3]
		{"nums = [1,2,3,3,2,7], val = 2", []int{1, 2, 3, 3, 2, 7}, 2, 4},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expectation, removeElement(tt.nums, tt.val))
		})
	}
}
