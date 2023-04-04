package max_area_of_island

import (
	"testing"

	"github.com/stretchr/testify/require"
)

var (
	grid1 = [][]int{
		{0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
		{0, 1, 1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 1, 0, 0, 1, 1, 0, 0, 1, 0, 1, 0, 0},
		{0, 1, 0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0},
	}
	grid2 = [][]int{{0, 0, 0, 0, 0, 0, 0, 0}}
	grid3 = [][]int{
		{1, 1, 0, 0, 0},
		{1, 1, 0, 0, 0},
		{0, 0, 0, 1, 1},
		{0, 0, 0, 1, 1},
	}
)

func TestNumberOfIsland(t *testing.T) {
	testCases := []struct {
		name string
		grid [][]int
		want int
	}{
		{"grid 1:", grid1, 6},
		{"grid 2:", grid2, 0},
		{"grid 3:", grid3, 2},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			require.Equal(t, tt.want, maxAreaOfIsland(tt.grid))
		})
	}
}
