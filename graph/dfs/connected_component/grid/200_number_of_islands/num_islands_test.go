package number_of_islands

import (
	"testing"

	"github.com/stretchr/testify/require"
)

var (
	grid1 = [][]byte{
		{'1', 1, 1, 1, 0},
		{1, 1, 0, 1, 0},
		{1, 1, 0, 0, 0},
		{0, 0, 0, 0, 0},
	}

	grid2 = [][]byte{
		{1, 1, 0, 0, 0},
		{1, 1, 0, 0, 0},
		{0, 0, 1, 0, 0},
		{0, 0, 0, 1, 1},
	}

	grid3 = [][]byte{
		{1, 0},
		{0, 1},
	}

	grid4 = [][]byte{
		{0, 0, 0},
		{0, 1, 0},
		{0, 0, 0},
	}
)

func TestSolution(t *testing.T) {
	testcases := []struct {
		description string
		grid        [][]byte
		want        int
	}{
		{"Input: grid = [[1, 1, 1, 1, 0  " +
			"[1, 1, 0, 1, 0  " +
			"[1, 1, 0, 0, 0 " +
			"[0, 0, 0, 0, 0 ]", grid1, 1},

		{"Input: grid =  [   [1, 1, 0, 0, 0  " +
			"[1, 1, 0, 0, 0  " +
			"[0, 0, 1, 0, 0  " +
			"[0, 0, 0, 1, 1 ] ", grid2, 3},

		{"Input: grid = [[1, 0],  " +
			"[0, 1],  ", grid3, 2},

		{"Input: grid = [ [0, 0, 0]," +
			"[0, 1, 0], " +
			"[0, 0, 0], ", grid4, 1},
	}

	for _, testcase := range testcases {
		t.Run(testcase.description, func(t *testing.T) {
			require.Equal(t, testcase.want, numIslands(testcase.grid))
		})
	}
}
