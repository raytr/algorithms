package degree_of_an_array

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindShortestSubArray(t *testing.T) {
	testCases := []struct {
		name        string
		nums        []int
		expectation int
	}{
		{"nums = [1,2,2,3,1]", []int{1, 2, 2, 3, 1}, 2},
		{"nums = [1,2,2,3,1,4,2]", []int{1, 2, 2, 3, 1, 4, 2}, 6},
		{"nums = [1,1,1]", []int{1, 1, 1}, 3},
		{"nums = [0]", []int{0}, 1},
		{"nums = [2,1,1,2,1,3,3,3,1,3,1,3,2]", []int{2, 1, 1, 2, 1, 3, 3, 3, 1, 3, 1, 3, 2}, 7},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expectation, findShortestSubArray1(tt.nums))
			assert.Equal(t, tt.expectation, findShortestSubArray2(tt.nums))
		})
	}
}
