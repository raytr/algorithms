package longest_mountain_in_array

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLongestMountain(t *testing.T) {
	testCases := []struct {
		name     string
		arr      []int
		expected int
	}{
		{"[2,1,4,7,3,2,5]", []int{2, 1, 4, 7, 3, 2, 5}, 5},
		{"[2,1,4,7,3,2]", []int{2, 1, 4, 7, 3, 2}, 5},
		{"[1, 2, 3, 2, 1]", []int{1, 2, 3, 2, 1}, 5},
		{"[2,2,2]", []int{2, 2, 2}, 0},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			got := longestMountain(tt.arr)
			assert.Equal(t, tt.expected, got)
		})
	}
}
