package spiral_matrix

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSpiralOrder(t *testing.T) {
	tests := []struct {
		name        string
		matrix      [][]int
		expectation []int
	}{
		{"matrix = [[1,2,3], [4,5,6], [7,8,9]]", [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}, []int{1, 2, 3, 6, 9, 8, 7, 4, 5}},
		{"matrix = [[1,2,3,4], [5,6,7,8], [9,10,11,12]]", [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}}, []int{1, 2, 3, 4, 8, 12, 11, 10, 9, 5, 6, 7}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expectation, spiralOrder(tt.matrix))
		})
	}
}
