package container_with_most_water

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMaxArea(t *testing.T) {

	testCases := []struct {
		name        string
		height      []int
		expectation int
	}{
		{"[1,8,6,2,5,4,8,3,7]", []int{1, 8, 6, 2, 5, 4, 8, 3, 7}, 49},
		{"[1,1]", []int{1, 1}, 1},
		{"[0,2]", []int{0, 2}, 0},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			require.Equal(t, tt.expectation, maxArea(tt.height))
		})
	}

}
