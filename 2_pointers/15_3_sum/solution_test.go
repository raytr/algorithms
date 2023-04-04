package three_sum

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSolution(t *testing.T) {
	testCases := []struct {
		name string
		nums []int
		want [][]int
	}{
		{"[-1, 0, 1, 2, -1, -4]", []int{-1, 0, 1, 2, -1, -4}, [][]int{{-1, -1, 2}, {-1, 0, 1}}},
		{"[0, 0, 0]", []int{0, 0, 0}, [][]int{{0, 0, 0}}},
		{"[-2, 0, 1, 1, 2]", []int{-2, 0, 1, 1, 2}, [][]int{{-2, 0, 2}, {-2, 1, 1}}},
		{"[]", []int{}, [][]int{}},
		{"[0]", []int{0}, [][]int{}},
		{"[-4, -2, 1, -5, -4, -4, 4, -2, 0, 4, 0, -2, 3, 1, -5, 0]", []int{-4, -2, 1, -5, -4, -4, 4, -2, 0, 4, 0, -2, 3, 1, -5, 0}, [][]int{{-5, 1, 4}, {-4, 0, 4}, {-4, 1, 3}, {-2, -2, 4}, {-2, 1, 1}, {0, 0, 0}}},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			require.Equal(t, tt.want, threeSum(tt.nums))
			require.Equal(t, tt.want, threeSum1(tt.nums))
		})
	}
}
